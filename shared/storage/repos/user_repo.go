package repos

import "iskra/shared/models"

type UserRepo interface {
	GetUser(chatID int64) (models.UserDB, error)
	CreateUser(user models.UserCreate) error
	UpdateUser(user models.UserDB) error
}
