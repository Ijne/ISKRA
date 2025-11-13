package handlers

import (
	"encoding/json"
	"iskra/bot"
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/config"
	"iskra/shared/models"
	"iskra/shared/storage/memgraph"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func InteractionHandler(cfg *config.Config, s *postgres.Storage, g *memgraph.Storage, b *bot.Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			var req struct {
				User_id          int64  `json:"user_id"`
				Target_user_id   int64  `json:"target_user_id"`
				Interaction_type string `json:"interaction_type"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				log.Printf("ERROR FROM[InteractionHandler] json decode err: %s", err)
				return
			}

			if req.Interaction_type == "like" {
				if err := s.LikeRepo.SetLike(req.User_id, req.Target_user_id); err != nil {
					log.Printf("ERROR FROM[InteractionHandler] SetLike err: %s", err)
					return
				}
				if s.LikeRepo.IsLike(req.Target_user_id, req.User_id) {
					match := models.MatchDB{
						MothID:  req.User_id,
						LightID: req.Target_user_id,
					}
					if err := s.MatchesRepo.Create(match); err != nil {
						log.Printf("ERROR FROM[InteractionHandler] CreateMatch err: %s", err)
						return
					}
					// ПРОИЗОШЕЛ МЭТЧ - НАДО ЗАПУСТИТЬ НУЖНУЮ ЛОГИКУ ПОСЛЕ ЭТОГО
					haveMatch := s.MatchesRepo.Exists(req.Target_user_id, req.User_id)
					log.Printf("Before match: %v\n", haveMatch)
					if haveMatch {
						// с помощью бота рассылаем ники
						if b != nil {
							user1, err := s.UserRepo.GetUser(req.User_id)
							if err != nil {
								log.Println("error while getting user1")
								render.JSON(w, r, response.Error("Server error"))
								return
							}
							user2, err := s.UserRepo.GetUser(req.Target_user_id)
							if err != nil {
								log.Println("error while getting user2")
								render.JSON(w, r, response.Error("Server error"))
								return
							}

							b.SendNick(user1.ID, user2.Username)
							b.SendNick(user2.ID, user1.Username)
						}

						s.MatchesRepo.Delete(req.Target_user_id, req.User_id)
						s.MatchesRepo.Delete(req.User_id, req.Target_user_id)

						render.JSON(w, r, response.Ok())
						return
					}
				}
			}

			if err := g.SocialWebRepo.SetSwipe(req.User_id, req.Target_user_id, req.Interaction_type); err != nil {
				log.Printf("ERROR FROM[InteractionHandler] SetSwipe err: %s", err)
				return
			}

			log.Printf("SUCCESS FROM[InteractionHandler] Interaction of type[%s] of user[id%d] set", req.Interaction_type, req.User_id)
		default:
			log.Printf("ERROR FROM[InteractionHandler] Not allowed method: %s", r.Method)
		}
	}
}
