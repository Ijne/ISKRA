package handlers

import (
	"iskra/internal/tools"
	"log"
	"net/http"
)

func StartScreenHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := tools.RenderTemplate(w, "start.html", struct{}{}); err != nil {
			log.Panicln(err)
		} else {
			log.Println("Successefully rendered")
		}
	default:
		// Дописать
	}
}
