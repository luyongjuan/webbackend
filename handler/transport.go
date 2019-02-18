package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var ErrInvalidArgument = errors.New("invalid argument")


// MakeHandler returns a handler for the booking service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}


	helloHandler := kithttp.NewServer(
		makeHelloWordEndpoint(bs),
		decodeListLocationsRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/booking/v1/locations", helloHandler).Methods("GET")

	return r
}

var errBadRoute = errors.New("bad route")


func decodeListLocationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "hello", nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
