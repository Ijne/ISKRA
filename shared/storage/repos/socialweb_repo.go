package repos

import "iskra/shared/models"

type SocialWebRepo interface {
	CreateUser(user models.UserCreate) error
	UpdateUser(user models.UserDB) error
	GetRecommendations(id int64) ([]models.UserResponse, error)
	SetSwipe(id1, id2 int64, interaction_type string) error
}
