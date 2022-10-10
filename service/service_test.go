package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayService(t *testing.T) {
	t.Setenv("RANDOM_CHOICE_URL", "https://codechallenge.boohma.com/random")
	service := NewService()

	t.Run("health", func(t *testing.T) {
		c, _ := service.GetHealth(context.Background())
		assert.Equal(t, "ok", c)
	})

	t.Run("play game", func(t *testing.T) {
		_, err := service.Play(context.Background(), 1, 3)
		assert.Equal(t, nil, err)
	})
}
