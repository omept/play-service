package service

import (
	"context"

	"github.com/ong-gtp/play-service/entities"
	"github.com/ong-gtp/play-service/repository"
)

// Service defines the interface for play-service
type Service interface {
	GetHealth(ctx context.Context) (string, error)
	Play(ctx context.Context, player, opponent int8) (entities.PlayResponse, error)
}

// service is the struct that will implement Service interfacce
type service struct{}

// NewService creates a new instance of service
func NewService() *service {
	return &service{}
}

var repo repository.Repository = repository.NewRepository()

// GetHealth returns health service
func (s *service) GetHealth(ctx context.Context) (string, error) {
	return "ok", nil
}

// Play evaluates game results
func (s *service) Play(ctx context.Context, player, opponent int8) (entities.PlayResponse, error) {
	return repo.EvaluateGame(player, opponent)
}
