package handlers

import (
	"fmt"
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/miniapp/internal/tools/timepad"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func GetEventsHandler(s *postgres.Storage, t *timepad.TimepadPoller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// return all saved events
			res, err := s.FlamesRepo.GetEvents()
			if err != nil {
				log.Printf("ERROR FROM[GetEventsHandler] GetEvents err: %s", err)
				render.JSON(w, r, struct {
					response.Response
					models.ManyEventsResponse
				}{
					response.Ok(),
					models.ManyEventsResponse{Events: nil},
				})
				return
			}

			fmt.Println(res[0], "PIIIIIIIIIIIIIIIIZDOSSSSSSSSSSSSSSSS")

			// convert structs
			resp := make([]models.EventResponse, len(res))
			for i := 0; i < len(res); i++ {
				event := res[i]
				log.Println("BOLSHAAAAAAAAAAAAYA PISKAAAAAAAAAA")
				curr := models.EventResponse{
					ID:       event.ID,
					StartsAt: event.StartsAt.Format("2006-01-02T15:04:05-0700"),
					Name:     event.Name,
					Url:      event.Url,
					Photo:    event.Photo,
				}
				resp[i] = curr
				log.Println("POSOSALIIIIIIIIIIIIIIIIIIIIIIII", resp)
			}

			render.JSON(w, r, struct {
				response.Response
				models.ManyEventsResponse
			}{
				response.Ok(),
				models.ManyEventsResponse{Events: resp},
			})
			log.Println("POSOSALIIIIIIIIIIIIIIIIIIIIIIII", resp)

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
				log.Printf("ERROR FROM[GetEventsHandler] decode json err: %s", err)
				render.JSON(w, r, response.Error("Wrong json"))
				return
			}

			userID, _ := middleware.GetUserIDFromContext(r.Context())
			log.Printf("userId: %d\n", userID)
			user, err := s.UserRepo.GetUser(userID)
			log.Printf("%v\n", user)

			events := make([]timepad.Event, 0)

			if err != nil {
				log.Printf("ERROR FROM[GetEventsHandler] GetUser err: %s", err)
				events, err = t.GetEvents()
				if err != nil {
					log.Printf("ERROR FROM[GetEventsHandler] GetEvents err: %s", err)
					render.JSON(w, r, response.Error("Server error"))
					return
				}
			} else {
				cats := strings.Split(user.EventPreferences, ",")
				// log.Printf("%v\n", cats)

				filter := timepad.EventsFilter{
					City:       user.City,
					Categories: cats,
					Limit:      req.Limit,
					Skip:       req.Skip,
				}

				events, err = t.GetFilteredEvents(filter)
				log.Printf("%vJOPAAAAAAAAAAAAAAAAAAAAAAAAAA\n", events)
				if err != nil {
					events, err = t.GetEvents()
					if err != nil {
						render.JSON(w, r, response.Error("Server error"))
						return
					}
				}
			}

			fmt.Println(events, "PISOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOSOCHKA")

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

			fmt.Printf("events, resulting slice: %v\n", res)

			render.JSON(w, r, struct {
				response.Response
				models.ManyEventsResponse
			}{
				response.Ok(),
				models.ManyEventsResponse{Events: res},
			})
		}
	}
}
