package repository

import (
	"context"
	"io"

	"github.com/Albert-Ti/goph-profile/internal/types"
)

type Repository interface {
	CreateAvatar(ctx context.Context, userID, filename, mimetype, s3Key string, size int64) (string, error)
	UpdateUploadedStatus(ctx context.Context, userID, avatarID string, status types.UploadStatus, size int64) error
	UpdateProcessingStatus(ctx context.Context, userID, avatarID string, status types.ProcessingStatus) error
}

func NewRepository(connString string) (Repository, error) {
	return NewPostgres(connString)
}

type ObjStorage interface {
	Put(ctx context.Context, key string, r io.Reader) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}

func NewObjStorage(endpoint, accessKey, secretKey, bucketName string) (ObjStorage, error) {
	return NewS3Client(context.Background(), endpoint, accessKey, secretKey, bucketName)
}
