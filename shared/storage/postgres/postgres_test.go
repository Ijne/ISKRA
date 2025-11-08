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

	s, err := NewStorage(cfg)
	assert.Nil(t, err)

	t.Run("Basic test", func(t *testing.T) {
		err := s.UserRepo.CreateUser(models.UserCreate{
			ID:       1,
			Username: "@test",
			Name:     "Test",
			Surname:  "Test",
			Age:      19,
			Gender:   0,
		})
		assert.Nil(t, err)

		user, err := s.UserRepo.GetUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "@test", user.Username)
	})

}
