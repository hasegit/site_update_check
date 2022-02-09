package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Url string `json:"url"`
}

func HandleRequest(ctx context.Context, e Event) (bool, error) {
	result, _ := checker(e.Url)
	return result, nil
}

func main() {
	lambda.Start(HandleRequest)
}
