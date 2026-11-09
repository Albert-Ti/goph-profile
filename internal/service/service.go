package service

import (
	"github.com/Albert-Ti/goph-profile/internal/config"
	"github.com/Albert-Ti/goph-profile/internal/repository"
)

type Service struct {
	repository repository.Repository
	opts       *config.Options
}

func NewService(repo repository.Repository, opts *config.Options) *Service {
	return &Service{repo, opts}
}
