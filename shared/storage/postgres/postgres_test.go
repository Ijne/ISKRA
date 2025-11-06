package postgres

import (
	"iskra/shared/config"
	"iskra/shared/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T) {
	cfg, err := config.New("./../../../config/local.yaml")
	assert.Nil(t, err)

	s, err := New(cfg)
	assert.Nil(t, err)

	t.Run("Basic test", func(t *testing.T) {
		err := s.UserRepo.CreateUser(models.UserCreate{
			ChatID:      123,
			Username:    "abc",
			Nick:        "abc",
			Description: "cba",
			Icon:        "https:///",
		})
		assert.Nil(t, err)

		user, err := s.UserRepo.GetUser(123)
		assert.Nil(t, err)
		assert.Equal(t, "abc", user.Username)
	})

}
