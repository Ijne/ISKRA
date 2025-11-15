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
		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
			if err != nil {
				log.Printf("ERROR FROM[GetRecomendationsHandler] parseint err: %s", err)
				return
			}

			recommendedUsers, err := g.SocialWebRepo.GetRecommendations(id)
			if err != nil {
				log.Printf("ERROR FROM[GetRecomendationsHandler] memgrph err: %s", err)
				return
			}

			if err := json.NewEncoder(w).Encode(recommendedUsers); err != nil {
				log.Printf("ERROR FROM[GetRecomendationsHandler] json encode err: %s", err)
				return
			}

			log.Printf("SUCCESS FROM[GetRecomendationsHandler] User[id%d] got his recommendations", id)
		default:
			log.Printf("ERROR FROM[GetRecomendationsHandler] Not allowed http method: %s", r.Method)
		}
	}
}
