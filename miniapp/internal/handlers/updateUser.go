package handlers

import (
	"encoding/json"
	"iskra/shared/config"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
)

func UpdateUserHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			var user models.UserDB
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Println(err)
			}

			container, err := postgres.NewStorage(cfg)
			if err != nil {
				log.Println(err)
			}
			if err := container.UserRepo.UpdateUser(user); err != nil {
				log.Println(err)
			}
		default:
			// Дописать
		}
	}
}
