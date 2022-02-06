package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Url string `json:"url"`
}

func HandleRequest(ctx context.Context, e Event) (bool, error) {
	a, _ := checker("http://xml.kishou.go.jp/revise.html")
	return a, nil
}

func main() {
	lambda.Start(HandleRequest)
}
