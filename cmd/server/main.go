package main

import (
	"log/slog"
	"net/http"

	"github.com/Albert-Ti/goph-profile/internal/config"
	"github.com/Albert-Ti/goph-profile/internal/repository"
	"github.com/Albert-Ti/goph-profile/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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
