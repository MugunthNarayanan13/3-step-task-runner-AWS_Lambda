package main

import (
    "context"
    "github.com/aws/aws-lambda-go/lambda"
)

type NotifyInput struct {
    Result string `json:"result"`
}

func handler(ctx context.Context, in NotifyInput) (map[string]string, error) {
    return map[string]string{"status": "notified"}, nil
}

func main() {
    lambda.Start(handler)
}