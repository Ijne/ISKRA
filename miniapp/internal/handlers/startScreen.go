package handlers

import (
	"iskra/miniapp/internal/tools"
	"iskra/shared/config"
	"log"
	"net/http"
)

func StartScreenHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if err := tools.RenderTemplate(w, cfg, "start.html", struct{}{}); err != nil {
				log.Panicln(err)
			} else {
				log.Println("Successefully rendered")
			}
		default:
			// Дописать
		}
	}
}
