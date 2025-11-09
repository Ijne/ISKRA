package repos

import "iskra/shared/models"

type FlamesRepo interface {
	GetLim(limit int) ([]models.FlameDB, error)
	GetByEvent(eventID int64) ([]models.FlameDB, error)
	GetByEventJoinUsers(eventID int64) ([]models.FlameWithUserDB, error)
	GetByUser(userID int64) ([]models.FlameDB, error)
	Create(flame models.FlameDB) error
	Update(flame models.FlameDB) error
	Delete(eventID int64, userID int64) error

	CreateEvent(event models.EventDB) error
	GetEvents() ([]models.EventDB, error)
	FillCategories(categories []models.EventCategory) error
}
