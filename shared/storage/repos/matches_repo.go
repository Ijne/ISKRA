package repos

import "iskra/shared/models"

type MatchesRepo interface {
	Exists(mothID int64, flameID int64) (bool, error)
	Create(match models.MatchDB) error
	Delete(mothID int64) error
}
