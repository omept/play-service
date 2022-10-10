package endpoint

import (
	"context"
	"testing"

	"github.com/ong-gtp/play-service/entities"
	"github.com/ong-gtp/play-service/service"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetHealthEndpoint(t *testing.T) {
	s := service.NewService()
	ep := MakeGetHealthEndpoint(s)
	t.Run("health", func(t *testing.T) {
		var request interface{}
		r, err := ep(context.Background(), request)
		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
		assert.Equal(t, "ok", r)
	})
}

func TestPlayEndpoint(t *testing.T) {
	t.Setenv("RANDOM_CHOICE_URL", "https://codechallenge.boohma.com/random")
	s := service.NewService()
	ep := MakePlayEndpoint(s)
	t.Run("make play endoint is not breaking", func(t *testing.T) {
		request := entities.PlayRequest{Player: 2}
		_, err := ep(context.Background(), request)
		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
	})
}
