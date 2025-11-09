package handlers

import (
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"net/http"

	"github.com/go-chi/render"
)

func LikeUserHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.MatchRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, response.Error("Wrong json"))
			return
		}

		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			render.JSON(w, r, response.Error("Wrong authorization"))
			return
		}

		// произошёл ли match
		haveMatch := s.MatchesRepo.Exists(req.LightID, userID)
		if haveMatch {
			// TODO: с помощью бота рассылаем ники

			render.JSON(w, r, response.Ok())
			return
		}

		// лишний раз не сохраняем данные в бд
		exists := s.MatchesRepo.Exists(userID, req.LightID)
		if exists {
			render.JSON(w, r, response.Ok())
			return
		}

		err = s.MatchesRepo.Create(models.MatchDB{MothID: userID, LightID: req.LightID})
		if err != nil {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		render.JSON(w, r, response.Ok())
	}
}
