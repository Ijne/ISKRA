package handlers

import (
	"fmt"
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/response"
	"iskra/shared/models"
	"iskra/shared/storage/postgres"
	"net/http"

	"github.com/go-chi/render"
)

func GetProfileHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			fmt.Printf("no userID\n")
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		user, err := s.UserRepo.GetUser(userID)
		if err != nil {
			fmt.Printf("postgres: %v\n", err)
			render.JSON(w, r, response.Error("Server error"))
			return
		}

		render.JSON(w, r, struct {
			response.Response
			models.UserResponse
		}{
			response.Ok(),
			UserDBToUser(user),
		})
	}
}

func UserDBToUser(a models.UserDB) models.UserResponse {
	return models.UserResponse{
		ID:               a.ID,
		Username:         a.Username,
		Name:             a.Name,
		Surname:          a.Surname,
		Age:              a.Age,
		Gender:           a.Gender,
		PreferredGender:  a.PreferredGender,
		CareerType:       a.CareerType.String,
		PersonalityType:  a.PersonalityType.String,
		RelationshipGoal: a.RelationshipGoal.String,
		ImportantValues:  a.ImportantValues.String,
		City:             a.City.String,
		CareerPlace:      a.CareerPlace.String,
		Music:            a.Music.String,
		Films:            a.Films.String,
		Hobbies:          a.Hobbies.String,
		EventPreferences: a.EventPreferences.String,
	}
}
