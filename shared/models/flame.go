package models

import "database/sql"

type FlameDB struct {
	EventID     int64  `json:"event_id"`
	UserID      int64  `json:"user_id"`
	Description string `json:"description"`
}

type FlameWithUserDB struct {
	EventID     int64  `json:"event_id"`
	UserID      int64  `json:"user_id"`
	Description string `json:"description"`

	Username         string         `json:"username" db:"username"`
	Name             string         `json:"name" db:"name"`
	Surname          string         `json:"surname,omitempty" db:"surname"`
	Age              int            `json:"age" db:"age"`
	Gender           int            `json:"gender" db:"gender"`
	PreferredGender  int            `json:"preferred_gender,omitempty" db:"preferred_gender"`
	CareerType       sql.NullString `json:"career_type,omitempty" db:"career_ty"`
	PersonalityType  sql.NullString `json:"personality_type,omitempty" db:"personality_type"`
	RelationshipGoal sql.NullString `json:"relationship_goal,omitempty" db:"relationship_goal"`
	ImportantValues  sql.NullString `json:"important_values,omitempty" db:"important_values"`
	City             sql.NullString `json:"city,omitempty" db:"city"`
	CareerPlace      sql.NullString `json:"career_place,omitempty" db:"career_place"`
	Music            sql.NullString `json:"music,omitempty" db:"music"`
	Films            sql.NullString `json:"films,omitempty" db:"films"`
	Hobbies          sql.NullString `json:"hobbies,omitempty" db:"hobbies"`
	EventPreferences sql.NullString `json:"event_preferences,omitempty" db:"event_preferences"`
}

type FlameWithUserResponse struct {
	EventID     int64  `json:"event_id"`
	UserID      int64  `json:"user_id"`
	Description string `json:"description"`

	Username         string `json:"username" db:"username"`
	Name             string `json:"name" db:"name"`
	Surname          string `json:"surname,omitempty" db:"surname"`
	Age              int    `json:"age" db:"age"`
	Gender           int    `json:"gender" db:"gender"`
	PreferredGender  int    `json:"preferred_gender,omitempty" db:"preferred_gender"`
	CareerType       string `json:"career_type,omitempty" db:"career_ty"`
	PersonalityType  string `json:"personality_type,omitempty" db:"personality_type"`
	RelationshipGoal string `json:"relationship_goal,omitempty" db:"relationship_goal"`
	ImportantValues  string `json:"important_values,omitempty" db:"important_values"`
	City             string `json:"city,omitempty" db:"city"`
	CareerPlace      string `json:"career_place,omitempty" db:"career_place"`
	Music            string `json:"music,omitempty" db:"music"`
	Films            string `json:"films,omitempty" db:"films"`
	Hobbies          string `json:"hobbies,omitempty" db:"hobbies"`
	EventPreferences string `json:"event_preferences,omitempty" db:"event_preferencestring"`
}

type FlamesRequest struct {
	EventID int64 `json:"event_id"`
}

type ManyFlamesResponse struct {
	Flames []FlameWithUserResponse `json:"flames"`
}

type FlameCreate struct {
	EventID     int64  `json:"event_id"`
	Description string `json:"description"`
}

type FlameUpdate struct {
	EventID     int64  `json:"event_id"`
	Description string `json:"description"`
}

type FlameDelete struct {
	EventID int64 `json:"event_id"`
}

func FlameWithUserDBToResponse(a FlameWithUserDB) FlameWithUserResponse {
	return FlameWithUserResponse{
		EventID:     a.EventID,
		Description: a.Description,

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
