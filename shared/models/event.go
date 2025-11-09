package models

import "time"

type EventDB struct {
	ID       int64
	StartsAt time.Time
	Name     string
	Url      string
	Photo    string
}

type EventCategory struct {
	ID   int64
	Name string
}

type EventResponse struct {
	ID       int64
	StartsAt string
	Name     string
	Url      string
	Photo    string
}

type ManyEventsResponse struct {
	Events []EventResponse `json:"events"`
}

type FilteredEventsRequest struct {
	City       string   `json:"city"`
	Categories []string `json:"categories"`
	Limit      int      `json:"limit"`
	Skip       int      `json:"skip"`
}
