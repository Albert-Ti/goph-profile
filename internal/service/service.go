package service

import (
	"context"
	"strconv"
	"time"

	"github.com/Albert-Ti/goph-profile/internal/config"
	"github.com/Albert-Ti/goph-profile/internal/repository"
	"github.com/Albert-Ti/goph-profile/internal/types"
	"github.com/Albert-Ti/goph-profile/internal/utils"
)

type Service struct {
	repository repository.Repository
	storage    repository.ObjStorage
	opts       *config.Options
}

func NewService(repo repository.Repository, storage repository.ObjStorage, opts *config.Options) *Service {
	return &Service{repo, storage, opts}
}

func (s *Service) SaveAvatar(ctx context.Context, userID, filename, mimetype string, r *utils.CountReader) error {
	s3Key := "users/" + userID + "/avatars" + "/avatar_" +
		strconv.Itoa(int(time.Now().UnixMilli())) + "/original"

	id, err := s.repository.CreateAvatar(ctx, userID, filename, mimetype, s3Key, r.Size)
	if err != nil {
		return err
	}

	if err := s.storage.Put(ctx, s3Key, r); err != nil {
		if err := s.repository.UpdateUploadedStatus(
			ctx, userID, id, types.UploadStatusFailed, r.Size); err != nil {
			return err
		}
		return err
	}

	if err := s.repository.UpdateUploadedStatus(
		ctx, userID, id, types.UploadStatusUploaded, r.Size); err != nil {
		return err
	}

	return nil
}
