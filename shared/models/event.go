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
	City       string   `json:"city,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Skip       int      `json:"skip,omitempty"`
}
