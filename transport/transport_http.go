package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/ong-gtp/play-service/endpoint"
	"github.com/ong-gtp/play-service/helper"
	"github.com/ong-gtp/play-service/service"
)

// NewHttpServer defines the http handler endpoints for play service
func NewHttpServer(svc service.Service, logger log.Logger) *mux.Router {
	// options provided by the go-kit to facilitate error control
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(helper.EncodeErrorResponse),
	}

	// define health endpoint handler
	healthHandler := httptransport.NewServer(
		endpoint.MakeGetHealthEndpoint(svc), //use the endpoint
		helper.DecodeEmptyRequest,           //converts the parameters received via the request body into the struct expected by the endpoint
		helper.EncodeHttpResponse,           //converts the struct returned by the endpoint to a json response
		options...,
	)

	// define play endpoint handler
	playHandler := httptransport.NewServer(
		endpoint.MakePlayEndpoint(svc), //use the endpoint
		helper.DecodePlayRequest,       //converts the parameters received via the request body into the struct expected by the endpoint
		helper.EncodeHttpResponse,      //converts the struct returned by the endpoint to a json response
		options...,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/playsv/health").Handler(healthHandler)
	r.Methods("POST").Path("/play").Handler(playHandler)
	return r
}
