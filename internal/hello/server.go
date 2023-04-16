package hello

import (
	"context"
	"fmt"
	"github.com/digitalhouse-tech/go-lib-kit/response"
	"github.com/digitalhouse-tech/go-lib-util/lambda"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/awslambda"
	"log"
)

func NewHelloLambda(endpoints Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoints.Get, decodeGetHandler, lambda.EncodeResponse,
		lambda.HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(lambda.HandlerFinalizer(nil)))
}

func decodeGetHandler(ctx context.Context, payload []byte) (interface{}, error) {
	fmt.Println(ctx)
	fmt.Println(string(payload))
	return nil, nil
}

type (
	Endpoints struct {
		Get endpoint.Endpoint
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Get: makeGetEndpoint(),
	}
}

func makeGetEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("hello my friend!")
		return response.OK("success", nil, nil, nil), nil
	}
}
