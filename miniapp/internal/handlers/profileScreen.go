package handlers

import (
	"encoding/json"
	"iskra/shared/config"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
	"strconv"
)

func ProfileScreenHandler(cfg *config.Config, s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			if err != nil {
				log.Println(err)
			}

			// container, err := postgres.NewStorage(cfg)
			// if err != nil {
			// 	log.Println(err)
			// }
			user, err := s.UserRepo.GetUser(id)
			if err != nil {
				log.Println(err)
			}

			if err := json.NewEncoder(w).Encode(user); err != nil {
				log.Println(err)
			}

		default:
			// Дописать
		}
	}
}
