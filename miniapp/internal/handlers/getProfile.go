package handlers

import (
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"net/http"

	"github.com/go-chi/render"
)

func GetProfileHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		user, err := s.UserRepo.GetUser(userID)
		if err != nil {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		render.JSON(w, r, struct {
			response.Response
			models.UserDB
		}{
			response.Ok(),
			user,
		})
	}
}
