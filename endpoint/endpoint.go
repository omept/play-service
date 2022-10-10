package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ong-gtp/play-service/entities"
	"github.com/ong-gtp/play-service/service"
)

// MakeHealthEndpoint creates the go-kit enpoint for GetHealth
func MakeGetHealthEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cs, err := svc.GetHealth(ctx)
		if err != nil {
			return "", err
		}
		return cs, nil
	}
}

// MakePlayEndpoint creates the go-kit enpoint for Play
func MakePlayEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entities.PlayRequest)

		cs, err := svc.Play(ctx, req.Player, req.Computer)
		if err != nil {
			return entities.PlayResponse{}, err
		}
		return cs, nil
	}
}
