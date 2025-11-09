package handlers

import (
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"net/http"

	"github.com/go-chi/render"
)

func CreateFlameHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		var req models.FlameCreate
		err := render.DecodeJSON(r.Body, req)
		if err != nil {
			render.JSON(w, r, response.Error("Wrong json"))
			return
		}

		err = s.FlamesRepo.Create(models.FlameDB{
			EventID:     req.EventID,
			UserID:      userID,
			Description: req.Description,
		})
		if err != nil {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		render.JSON(w, r, response.Ok())
	}
}
