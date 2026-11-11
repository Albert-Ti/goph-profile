package repository

import (
	"context"
	"io"
)

type Repository interface {
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
