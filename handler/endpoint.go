package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeTestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "hello word!!!!!", nil
	}
}
