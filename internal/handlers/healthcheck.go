package handlers

import (
	"net/http"

	"github.com/Albert-Ti/goph-profile/internal/service"
)

/*
Поддержка query-параметров (размер, формат) — опционально
Кеширование заголовков — опционально
Content-Type в зависимости от формата
GET /api/v1/avatars/{avatar_id}?size=300x300&format=webp
Parameters:
  size: string (optional, values: "100x100", "300x300", "original")
  format: string (optional, values: "jpeg", "png", "webp")
Response 200:
  Binary image data
  Headers:
    Content-Type: image/jpeg|png|webp
    Cache-Control: max-age=86400
    ETag: "hash"
Response 404:
  {
    "error": "Avatar not found"
  }
*/

func HealthCheck(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
