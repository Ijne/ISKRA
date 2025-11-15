package models

type MatchDB struct {
	MothID  int64
	LightID int64
}

type MatchRequest struct {
	LightID int64 `json:"light_id"`
}
