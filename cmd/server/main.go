package main

import (
	"log/slog"
	"net/http"

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

	slog.Info("Running server", "host", opts.RunAddr)

	if err := http.ListenAndServe(opts.RunAddr, r); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
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
