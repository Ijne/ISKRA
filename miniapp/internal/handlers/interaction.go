package handlers

import (
	"encoding/json"
	"iskra/shared/config"
	"iskra/shared/models"
	"iskra/shared/storage/memgraph"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
)

func InteractionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
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
				log.Println(err)
			}

			if req.Interaction_type == "like" {
				postgresContainer, err := postgres.NewStorage(cfg)
				if err != nil {
					log.Println(err)
				}
				if err := postgresContainer.LikeRepo.SetLike(req.User_id, req.Target_user_id); err != nil {
					log.Println(err)
				}
				if postgresContainer.LikeRepo.IsLike(req.Target_user_id, req.User_id) {
					match := models.MatchDB{
						MothID:  req.User_id,
						LightID: req.Target_user_id,
					}
					if err := postgresContainer.MatchesRepo.Create(match); err != nil {
						log.Println(err)
					}
					// ПРОИЗОШЕЛ МЭТЧ - НАДО ЗАПУСТИТЬ НУЖНУЮ ЛОГИКУ ПОСЛЕ ЭТОГО
				}
			}

			memgraphContainer, err := memgraph.NewStorage(cfg)
			if err != nil {
				log.Println(err)
			}
			if err := memgraphContainer.SocialWebRepo.SetSwipe(req.User_id, req.Target_user_id, req.Interaction_type); err != nil {
				log.Println(err)
			}

		default:
			// Дописать
		}
	}
}
