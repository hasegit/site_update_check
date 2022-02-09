package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Url string `json:"url"`
}

func HandleRequest(ctx context.Context, e Event) (bool, error) {
	fmt.Printf("%#v", ctx)
	fmt.Printf("%#v", e)
	result, _ := checker(e.Url)
	return result, nil
}

func main() {
	lambda.Start(HandleRequest)
}
