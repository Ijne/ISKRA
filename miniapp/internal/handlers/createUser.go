package handlers

import (
	"encoding/json"
	"iskra/miniapp/internal/tools/image"
	"iskra/shared/config"
	"iskra/shared/models"
	"iskra/shared/storage/memgraph"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
)

func CreateUserHandler(cfg *config.Config, s *postgres.Storage, g *memgraph.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			var user models.UserCreate
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] json err: %s\n", err)
				return
			}

			path, err := image.SaveUserAvatar(user.ID, user.Photo)
			if err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] SaveUserAvatart err: %s\n", err)
			}
			user.Photo = path

			if err := s.UserRepo.CreateUser(user); err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] psgrs err: %s\n", err)
				return
			}

			if err := g.SocialWebRepo.CreateUser(user); err != nil {
				log.Printf("ERROR FROM[CrateUserHandler] memgrph err: %s\n", err)
				return
			}

		default:
			log.Printf("ERROR FROM[CrateUserHandler] Not allowed http method: %s", r.Method)
		}
	}
}
