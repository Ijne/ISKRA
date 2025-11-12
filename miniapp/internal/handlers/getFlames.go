package handlers

import (
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"net/http"

	"github.com/go-chi/render"
)

func GetFlamesHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.FlamesRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, response.Error("Wrong json"))
			return
		}

		res, err := s.FlamesRepo.GetByEventJoinUsers(req.EventID)
		if err != nil {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		// resp := make([]models.FlameWithUserDB, len(res))
		// for i, flame := range res {
		// 	resp[i] = models.FlameWithUserDBToResponse(flame)
		// }

		render.JSON(w, r, struct {
			response.Response
			models.ManyFlamesResponse
		}{
			response.Ok(),
			models.ManyFlamesResponse{Flames: res},
		})
	}
}
