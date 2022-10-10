package helper

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/ong-gtp/play-service/entities"
)

// DecodeEmptyRequest converts the parameters received via the request body into the struct expected by the endpoint
func DecodeEmptyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interface{}
	return request, nil
}

// DecodeEmptyRequest converts the parameters received via the request body into the struct expected by the endpoint
func DecodePlayRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entities.PlayRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrValidation
	}

	if request.Player < 1 || request.Player > 5 {
		return nil, ErrValidation
	}

	if request.Computer < 0 || request.Computer > 5 {
		return nil, ErrValidation
	}

	return request, nil
}

// EncodeHttpResponse converts the struct returned by the endpoint to a json response
func EncodeHttpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// EncodeErrorResponse converts errors to json response
func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// codeFrom returns the http status code from service errors
func codeFrom(err error) int {
	switch err {
	case ErrInvalidOption:
		return http.StatusBadGateway
	case ErrValidation:
		return http.StatusBadRequest
	case ErrRandomUrlServiceNotPassed:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func Log(l, k, v interface{}) {
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	logger.Log("level", l, k, v)
}
