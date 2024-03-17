package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type StepFunctionInput struct {
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	Payload Payload `json:"Payload"`
}

type Payload struct {
	Issue Issue `json:"issue"`
}

type Issue struct {
	Description string `json:"description"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

func HandleRequest(ctx context.Context, event StepFunctionInput) (Response, error) {
	description := event.Parameters.Payload.Issue.Description
	if description == "" {
		return Response{Body: "Empty Description in JSON", StatusCode: 400}, nil
	}

	return Response{Body: description, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
