package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Url string `json:"url"`
}

type Response struct {
	Result int `json:"result"`
}

func HandleRequest(ctx context.Context, e Event) (Response, error) {
	result, _ := checker(e.Url)
	return Response{Result: result}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
