package handlers

import (
	"encoding/json"
	"iskra/miniapp/internal/tools/image"
	"iskra/shared/config"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
	"strconv"
)

func ProfileScreenHandler(cfg *config.Config, s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			if err != nil {
				log.Printf("ERROR FROM[ProfileScreenHandler] parseint err: %s", err)
				return
			}

			user, err := s.UserRepo.GetUser(id)
			if err != nil {
				log.Printf("ERROR FROM[ProfileScreenHandler] GetUser err: %s", err)
				return
			}

			user.Photo = image.GetUserPhoto(user.Photo)

			if err := json.NewEncoder(w).Encode(user); err != nil {
				log.Printf("ERROR FROM[ProfileScreenHandler] json encode err: %s", err)
				return
			}

			log.Printf("SUCCESS FROM[ProfileScreenHandler] Data for ProfileScreen for user [id%d] sent", user.ID)
		default:
			log.Printf("ERROR FROM[ProfileScreenHandler] Not allowed method: %s", r.Method)
		}
	}
}
