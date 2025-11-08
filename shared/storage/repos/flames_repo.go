package repos

import "iskra/shared/models"

type FlamesRepo interface {
	GetLim(limit int) ([]models.FlameDB, error)
	GetBeEvent()
	Create(flame models.FlameDB) error
	GetUsers(limit int) ([]models.UserDB, error)
}
