package repos

import "iskra/shared/models"

type MatchesRepo interface {
	Exists(mothID int64, lightID int64) bool
	Create(match models.MatchDB) error
	Delete(mothID int64, lightID int64) error

	FlamesCreate(match models.MatchDB) error
	FlamesExists(mothID int64, lightID int64) bool
}
