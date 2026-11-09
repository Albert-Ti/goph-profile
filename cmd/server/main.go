package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Albert-Ti/goph-profile/internal/config"
	"github.com/Albert-Ti/goph-profile/internal/repository"
	"github.com/Albert-Ti/goph-profile/internal/service"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // ← для работы с PostgreSQL
	_ "github.com/golang-migrate/migrate/v4/source/file"       // ← для чтения файлов
)

func main() {
	opts, errCfg := config.Build()
	if errCfg != nil {
		panic(errCfg)
	}

	repo, err := repository.NewRepository(opts.DBConnStr)
	if err != nil {
		panic(err)
	}

	if err := runMigrations(opts.DBConnStr); err != nil {
		panic(err)
	}

	_ = service.NewService(repo, opts)

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	server := &http.Server{Addr: opts.RunAddr, Handler: r}

	idleConnsClosed := make(chan struct{})

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	slog.Info("Running server", "host", opts.RunAddr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("HTTP server shutdown", "error", err)
		return
	}
	<-idleConnsClosed

	slog.Info("HTTP server shutdown gracefully")

}

func runMigrations(dsn string) error {
	if dsn == "" {
		return nil
	}
	m, errMigrate := migrate.New("file://migrations", dsn)
	if errMigrate != nil {
		return errMigrate
	}
	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			slog.Error("failed to close migrate instance", "source_error", srcErr, "db_error", dbErr)
		}
	}()

	errMigrate = m.Up()
	switch errMigrate {
	case nil:
		slog.Info("migrations applied successfully")
	case migrate.ErrNoChange:
		slog.Info("database schema is up-to-date")
	default:
		return errMigrate
	}
	return nil
}
