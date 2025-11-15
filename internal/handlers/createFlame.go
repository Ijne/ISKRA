package handlers

import (
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/miniapp/internal/tools/timepad"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

func CreateFlameHandler(s *postgres.Storage, t *timepad.TimepadPoller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			log.Println("ERROR FROM[CreateFlameHandler] GetUserIDFromContext not ok")
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		var req models.FlameCreate
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Printf("ERROR FROM[CreateFlameHandler] decode json err: %s", err)
			render.JSON(w, r, response.Error("Wrong json"))
			return
		}

		// save event if there is no in db
		saved := s.FlamesRepo.EventSaved(req.EventID)
		if !saved {
			log.Println("Save event")
			event, err := t.GetEventByID(req.EventID)
			if err != nil {
				log.Printf("ERROR FROM[CreateFlameHandler] GetEventByID err: %s", err)
				render.JSON(w, r, response.Error("Wrong event id"))
				return
			}

			date, err := time.Parse("2006-01-02T15:04:05-0700", event.StartsAt)
			if err != nil {
				log.Printf("ERROR FROM[CreateFlameHandler] wrong date err: %s", err)
				render.JSON(w, r, response.Error("Server error"))
				return
			}

			s.FlamesRepo.CreateEvent(models.EventDB{
				ID:       event.ID,
				StartsAt: date,
				Name:     event.Name,
				Url:      event.URL,
				Photo:    event.PosterImage.DefaultURL,
			})
		}

		err = s.FlamesRepo.Create(models.FlameDB{
			EventID:     req.EventID,
			UserID:      userID,
			Description: req.Description,
		})
		if err != nil {
			log.Printf("ERROR FROM[CreateFlameHandler] Create Flame err: %s", err)
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		render.JSON(w, r, response.Ok())
	}
}
