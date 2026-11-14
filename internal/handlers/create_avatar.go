package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/Albert-Ti/goph-profile/internal/service"
	"github.com/Albert-Ti/goph-profile/internal/utils"
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
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Ограничиваем общий размер запроса.
		r.Body = http.MaxBytesReader(w, r.Body, 100<<20)

		// Создаём потоковый multipart reader.
		mr, err := r.MultipartReader()
		if err != nil {
			http.Error(w, "Invalid multipart", http.StatusBadRequest)
			return
		}

		var (
			reader   io.Reader
			filename string
			mimetype string
		)

		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}

			if err != nil {
				http.Error(w, "Multipart read error", http.StatusBadRequest)
				return
			}

			filename = part.FileName()
			mimetype = part.Header.Get("Content-Type")
			reader = part

			break
		}

		if reader == nil {
			http.Error(w, "File not found", http.StatusBadRequest)
			return
		}

		counter := &utils.CountReader{Reader: reader}

		userID := r.Header.Get("X-User-Id")
		if err := svc.SaveAvatar(r.Context(), userID, filename, mimetype, counter); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
