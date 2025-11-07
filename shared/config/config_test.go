package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("No errors, default env", func(t *testing.T) {
		cfg, err := New("./../../config/local.yaml")
		assert.Nil(t, err)
		assert.Equal(t, "localhost", cfg.Bot.Host)
		assert.Equal(t, "localhost", cfg.MiniApp.Host)
	})
}
