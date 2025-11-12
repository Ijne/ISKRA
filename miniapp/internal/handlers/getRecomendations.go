package handlers

import (
	"encoding/json"
	"iskra/shared/config"
	"iskra/shared/storage/memgraph"
	"log"
	"net/http"
	"strconv"
)

func GetRecomendationsHandler(cfg *config.Config, g *memgraph.Storage) http.HandlerFunc {
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

			log.Println(id)

			// memgraphContainer, err := memgraph.NewStorage(cfg)
			// if err != nil {
			// 	log.Println(err)
			// }
			recommendedUsers, err := g.SocialWebRepo.GetRecommendations(id)
			if err != nil {
				log.Println(err)
			}

			if err := json.NewEncoder(w).Encode(recommendedUsers); err != nil {
				log.Println(err)
			}

		default:
			// Дописать
		}
	}
}
