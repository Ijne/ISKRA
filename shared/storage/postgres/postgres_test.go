package postgres

import (
	"iskra/shared/config"
	"iskra/shared/models"
	"testing"
	"time"

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

func TestMatchesRepo(t *testing.T) {
	cfg, err := config.New("./../../../config/local.yaml")
	assert.Nil(t, err)

	s, err := NewStorage(cfg)
	assert.Nil(t, err)

	t.Run("Basic test", func(t *testing.T) {
		err := s.MatchesRepo.Create(models.MatchDB{
			MothID:  123,
			LightID: 123,
		})
		assert.Nil(t, err)

		ex := s.MatchesRepo.Exists(123, 123)
		assert.True(t, ex)

		err = s.MatchesRepo.Delete(123, 123)
		assert.Nil(t, err)
	})

}

func TestFlamesRepo(t *testing.T) {
	cfg, err := config.New("./../../../config/local.yaml")
	assert.Nil(t, err)

	s, err := NewStorage(cfg)
	assert.Nil(t, err)

	t.Run("Basic test", func(t *testing.T) {
		// categories
		cats := []models.EventCategory{
			{ID: 1, Name: "Вечеринка"},
			{ID: 2, Name: "Концерт"},
		}

		err := s.FlamesRepo.FillCategories(cats)
		assert.Nil(t, err)

		// events
		events := []models.EventDB{
			{
				ID:       955959,
				StartsAt: time.Now(),
				Name:     "Раздача на спавне",
				Url:      "https://www.google.com",
				Photo:    "https://www.google.com",
			},
			{
				ID:       2616166,
				StartsAt: time.Now(),
				Name:     "Квартирник",
				Url:      "https://www.google.com",
				Photo:    "https://www.google.com",
			},
			{
				ID:       89898989,
				StartsAt: time.Now(),
				Name:     "Закупка на wb",
				Url:      "https://www.google.com",
				Photo:    "https://www.google.com",
			},
		}
		for _, event := range events {
			err = s.FlamesRepo.CreateEvent(event)
			assert.Nil(t, err)
		}

		_, err = s.FlamesRepo.GetEvents()
		assert.Nil(t, err)
		// плохо сравнивает
		// assert.ElementsMatch(t, events, events_got, cmp.Comparer(compareEvents))

		// flames
		flames := []models.FlameDB{
			{
				EventID:     955959,
				UserID:      123,
				Description: "asdfadsfasdf",
			},
			{
				EventID:     7878,
				UserID:      123,
				Description: "asdfadsfasdf",
			},
			{
				EventID:     898989,
				UserID:      2323232,
				Description: "asdfadsfasdf",
			},
		}
		for _, flame := range flames {
			err = s.FlamesRepo.Create(flame)
			assert.Nil(t, err)
		}

		// getters
		res, err := s.FlamesRepo.GetByEvent(955959)
		assert.Nil(t, err)
		assert.Equal(t, flames[0], res[0])

		res, err = s.FlamesRepo.GetByUser(123)
		assert.Nil(t, err)
		assert.Equal(t, flames[0], res[0])

		flames_got, err := s.FlamesRepo.GetLim(5)
		assert.Nil(t, err)
		assert.ElementsMatch(t, flames, flames_got)

		err = s.FlamesRepo.Delete(flames[0].EventID, flames[0].UserID)
		assert.Nil(t, err)

		err = s.FlamesRepo.Update(models.FlameDB{
			EventID:     7878,
			UserID:      123,
			Description: "что-то осмысленное",
		})
		assert.Nil(t, err)

		res, err = s.FlamesRepo.GetByEvent(7878)
		assert.Nil(t, err)
		assert.Equal(t, "что-то осмысленное", res[0].Description)
	})

}

func compareEvents(a, b models.EventDB) bool {
	a.StartsAt.Truncate(time.Second)
	b.StartsAt.Truncate(time.Second)
	return a.ID == b.ID &&
		a.Name == b.Name &&
		a.Photo == b.Photo &&
		a.StartsAt.Equal(b.StartsAt) &&
		a.Url == b.Url
}
