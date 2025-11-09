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
