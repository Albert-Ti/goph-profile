package handlers

import (
	"net/http"

	"github.com/Albert-Ti/goph-profile/internal/service"
)

/*
Валидация формата (JPEG, PNG, WebP) — опционально
Ограничение размера (до 10MB)
Создание миниатюр (100x100, 300x300)
Асинхронная обработка через брокер

POST /api/v1/avatars
Headers:
  X-User-ID: string (required)
  Content-Type: multipart/form-data
Body:
  file: binary (required, max 10MB)
Response 201:
  {
    "id": "uuid",
    "user_id": "string",
    "url": "string",
    "status": "processing",
    "created_at": "2024-01-01T00:00:00Z"
  }
Response 400:
  {
    "error": "Invalid file format",
    "details": "Supported formats: jpeg, png, webp"
  }
Response 413:
  {
    "error": "File too large",
    "max_size": 10485760
  }
*/

func CreateAvatar(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
