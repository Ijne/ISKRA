package handlers

import (
	"iskra/bot"
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func LikeUserHandler(s *postgres.Storage, b *bot.Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.MatchRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, response.Error("Wrong json"))
			return
		}

		log.Printf("light_id: %d", req.LightID)

		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		// произошёл ли match
		haveMatch := s.MatchesRepo.Exists(req.LightID, userID)
		log.Printf("Before match: %v\n", haveMatch)
		if haveMatch {
			// с помощью бота рассылаем ники
			if b != nil {
				user1, err := s.UserRepo.GetUser(userID)
				if err != nil {
					render.JSON(w, r, response.Error("Server error"))
					return
				}
				user2, err := s.UserRepo.GetUser(req.LightID)
				if err != nil {
					render.JSON(w, r, response.Error("Server error"))
					return
				}

				b.SendNick(user1.ID, user2.Username)
				b.SendNick(user2.ID, user1.Username)
			}

			s.MatchesRepo.Delete(req.LightID, userID)
			s.MatchesRepo.Delete(userID, req.LightID)

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
			log.Printf("Error while createing match: %v\n", err)
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		render.JSON(w, r, response.Ok())
	}
}
