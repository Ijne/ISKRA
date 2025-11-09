package handlers

import (
	"iskra/miniapp/internal/tools/response"
	"iskra/miniapp/internal/tools/timepad"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"net/http"

	"github.com/go-chi/render"
)

func GetEventsHandler(s *postgres.Storage, t *timepad.TimepadPoller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// without filters
			// return all saved events
			res, err := s.FlamesRepo.GetEvents()
			if err != nil {
				render.JSON(w, r, struct {
					response.Response
					models.ManyEventsResponse
				}{
					response.Error("Server error"),
					models.ManyEventsResponse{Events: nil},
				})
				return
			}

			// convert structs
			resp := make([]models.EventResponse, len(res))
			for i, event := range res {
				curr := models.EventResponse{
					ID:       event.ID,
					StartsAt: event.StartsAt.Format("2006-01-02T15:04:05-0700"),
					Name:     event.Name,
					Url:      event.Url,
					Photo:    event.Photo,
				}
				resp[i] = curr
			}

			render.JSON(w, r, struct {
				response.Response
				models.ManyEventsResponse
			}{
				response.Ok(),
				models.ManyEventsResponse{Events: resp},
			})

			// events, err := t.GetEvents()
			// if err != nil {
			// 	log.Fatalf("get_events_handler: %v\n", err)
			// 	return
			// }
			// render.JSON(w, r, events)

		case http.MethodPost:
			// with or without filters
			// return found events
			var req models.FilteredEventsRequest
			err := render.DecodeJSON(r.Body, &req)
			if err != nil {
				render.JSON(w, r, response.Error("Wrong json"))
				return
			}

			events, err := t.GetFilteredEvents(req)
			if err != nil {
				render.JSON(w, r, response.Error("Server error"))
				return
			}

			// convert structs
			res := make([]models.EventResponse, len(events))
			for i, event := range events {
				curr := models.EventResponse{
					ID:       event.ID,
					StartsAt: event.StartsAt,
					Name:     event.Name,
					Url:      event.URL,
					Photo:    event.PosterImage.DefaultURL,
				}
				res[i] = curr
			}

			render.JSON(w, r, res)
		}
	}
}
