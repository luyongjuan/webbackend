package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

type emptyRequest struct { }


// MakeHandler returns a handler for the booking service.
func MakeHandler(conIns Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	spingtesthandler := kithttp.NewServer(
		makeTestEndpoint(conIns),
		decodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/tt/test", spingtesthandler).Methods("POST")

	return r
}


func decodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return emptyRequest{}, nil
}


var errBadRoute = errors.New("bad route")

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
	// here is error code in this service

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
