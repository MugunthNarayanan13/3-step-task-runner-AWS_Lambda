package main

import (
    "context"
    "errors"
    "github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
    Payload string `json:"payload"`
}

type Output struct {
    Payload string `json:"payload"`
}

func handler(ctx context.Context, in Input) (Output, error) {
    if in.Payload == "" {
        return Output{}, errors.New("empty payload")
    }
    return Output{Payload: in.Payload}, nil
}

func main() {
    lambda.Start(handler)
}