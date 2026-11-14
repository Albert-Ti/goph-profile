package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Albert-Ti/goph-profile/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrRowAffected = errors.New("rows affected 0")

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(connString string) (*Postgres, error) {
	poolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	poolCfg.MinConns = 4                       // держим минимум 4 живых соединения постоянно
	poolCfg.MaxConns = 20                      // верхний предел под нагрузку
	poolCfg.MaxConnIdleTime = 5 * time.Minute  // не закрывать соединения слишком агрессивно
	poolCfg.MaxConnLifetime = 30 * time.Minute // периодическая ротация соединений (защита от "протухания")

	pool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		pool: pool,
	}, nil
}

func (pg *Postgres) CreateAvatar(ctx context.Context, userID, filename, mimetype, s3Key string, size int64) (string, error) {
	sql := `
	INSERT INTO avatars (user_id, file_name, mime_type, size_bytes, s3_key) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	var id string
	err := pg.pool.QueryRow(ctx, sql, userID, filename, mimetype, size, s3Key).Scan(&id)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (pg *Postgres) UpdateUploadedStatus(ctx context.Context, userID, avatarID string, status types.UploadStatus, size int64) error {
	sql := `
	UPDATE avatars 
	SET upload_status = $1, size_bytes = $2
	WHERE id = $3 AND user_id = $4
	`
	_, err := pg.pool.Exec(ctx, sql, status, size, avatarID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) UpdateProcessingStatus(ctx context.Context, userID, avatarID string, status types.ProcessingStatus) error {
	sql := `
	UPDATE avatars 
	SET processing_status = $1
	WHERE id = $3 AND user_id = $4
	`
	_, err := pg.pool.Exec(ctx, sql, status, avatarID, userID)
	if err != nil {
		return err
	}

	return nil
}
