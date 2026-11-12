package handlers

import (
	"net/http"

	"github.com/Albert-Ti/goph-profile/internal/service"
)

/*
Response 200:
  {
    "id": "uuid",
    "user_id": "string",
    "file_name": "avatar.jpg",
    "mime_type": "image/jpeg",
    "size": 1024000,
    "dimensions": {
      "width": 1920,
      "height": 1080
    },
    "thumbnails": [
      {"size": "100x100", "url": "..."},
      {"size": "300x300", "url": "..."}
    ],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
*/

func GetAvatarMetadata(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
