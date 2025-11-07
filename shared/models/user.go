package models

// что будет возвращаться из postgres
type UserDB struct {
	ID               int    `json:"id" db:"id"`
	Username         string `json:"username" db:"username"`
	Name             string `json:"name" db:"name"`
	Surname          string `json:"surname,omitempty" db:"surname"`
	Age              int    `json:"age" db:"age"`
	Gender           int    `json:"gender" db:"gender"`
	PreferredGender  int    `json:"preferred_gender,omitempty" db:"preferred_gender"`
	CareerType       string `json:"career_type,omitempty" db:"career_type"`
	PersonalityType  string `json:"personality_type,omitempty" db:"personality_type"`
	RelationshipGoal string `json:"relationship_goal,omitempty" db:"relationship_goal"`
	ImportantValues  string `json:"important_values,omitempty" db:"important_values"`
	City             string `json:"city,omitempty" db:"city"`
	CareerPlace      string `json:"career_place,omitempty" db:"career_place"`
}

type UserCreate struct {
	ID               int    `json:"id" db:"id"`
	Username         string `json:"username" db:"username"`
	Name             string `json:"name" db:"name"`
	Surname          string `json:"surname,omitempty" db:"surname"`
	Age              int    `json:"age" db:"age"`
	Gender           int    `json:"gender" db:"gender"`
	PreferredGender  int    `json:"preferred_gender,omitempty" db:"preferred_gender"`
	CareerType       string `json:"career_type,omitempty" db:"career_type"`
	PersonalityType  string `json:"personality_type,omitempty" db:"personality_type"`
	RelationshipGoal string `json:"relationship_goal,omitempty" db:"relationship_goal"`
	ImportantValues  string `json:"important_values,omitempty" db:"important_values"`
	City             string `json:"city,omitempty" db:"city"`
	CareerPlace      string `json:"career_place,omitempty" db:"career_place"`
}

type UserResponse struct {
}

type UserRequest struct {
}
