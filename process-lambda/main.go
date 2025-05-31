package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

type StepInput struct {
    Payload string `json:"payload"`
}

type StepOutput struct {
    Result string `json:"result"`
}

func handler(ctx context.Context, in StepInput) (StepOutput, error) {
    // pretend work
    out := fmt.Sprintf("processed: %s", in.Payload)
    return StepOutput{Result: out}, nil
}

func main() {
    lambda.Start(handler)
}
