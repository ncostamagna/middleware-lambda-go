package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-kit/kit/transport/awslambda"
	"github.com/ncostamagna/middleware-lambda-go/internal/hello"
)

var h *awslambda.Handler

func init() {
	h = hello.NewHelloLambda(hello.MakeEndpoints())
}

func main() {
	lambda.StartHandler(h)
}
