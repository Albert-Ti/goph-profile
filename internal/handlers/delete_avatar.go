package handlers

import (
	"net/http"

	"github.com/Albert-Ti/goph-profile/internal/service"
)

/*
DELETE /api/v1/avatars/{avatar_id}
Headers:
  X-User-ID: string (required)
Response 204: No Content
Response 403:
  {
    "error": "Forbidden",
    "details": "You can only delete your own avatars"
  }
*/

func DeleteAvatar(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
