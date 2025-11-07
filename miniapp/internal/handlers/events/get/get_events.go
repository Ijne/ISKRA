package get

import (
	"iskra/miniapp/internal/tools/timepad"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func GetEventsHandler(t *timepad.TimepadPoller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			events, err := t.GetEvents()
			if err != nil {
				log.Fatalf("get_events_handler: %v\n", err)
				return
			}
			render.JSON(w, r, events)
		default:
			// Дописать
		}
	}
}
