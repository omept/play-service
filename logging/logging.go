package logging

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/ong-gtp/play-service/entities"
	"github.com/ong-gtp/play-service/service"
)

// NewLoggingMiddleware returns the logging middleware for the service
func NewLoggingMiddleware(logger log.Logger, next service.Service) logmw {
	return logmw{logger, next}
}

// logmw defines the interface for the logging middleware implemented for service.Service
type logmw struct {
	logger log.Logger
	service.Service
}

// GetHealth defines the logging middleware for service.GetHealth
func (mw logmw) GetHealth(ctx context.Context) (response string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getHealth",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	response, err = mw.Service.GetHealth(ctx)
	return
}

// Play defines the logging middleware for service.Play
func (mw logmw) Play(ctx context.Context, player, opponent int8) (r entities.PlayResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "play",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	r, err = mw.Service.Play(ctx, player, opponent)
	return
}
