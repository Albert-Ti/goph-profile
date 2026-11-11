package service

import (
	"github.com/Albert-Ti/goph-profile/internal/config"
	"github.com/Albert-Ti/goph-profile/internal/repository"
)

type Service struct {
	repository repository.Repository
	storage    repository.ObjStorage
	opts       *config.Options
}

func NewService(repo repository.Repository, storage repository.ObjStorage, opts *config.Options) *Service {
	return &Service{repo, storage, opts}
}
