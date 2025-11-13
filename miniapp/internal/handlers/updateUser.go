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

func UpdateUserHandler(cfg *config.Config, s *postgres.Storage, g *memgraph.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			var user models.UserDB
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Printf("ERROR FROM[UpdateUserHandler] json decode err: %s", err)
				return
			}

			if err := s.UserRepo.UpdateUser(user); err != nil {
				log.Printf("ERROR FROM[UpdateUserHandler] postgres Updateuser err: %s", err)
				return
			}

			if err := g.SocialWebRepo.UpdateUser(user); err != nil {
				log.Printf("ERROR FROM[UpdateUserHandler] memgrpah UpdateUser err: %s", err)
				return
			}

			log.Printf("SUCCESS FROM[UpdateUserHandler] User[id%d] updated", user.ID)
		default:
			log.Printf("ERROR FROM[UpdateUserHandler] Not allowed method: %s", r.Method)
		}
	}
}
