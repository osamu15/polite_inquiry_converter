package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Issue Issue `json:"issue"`
}

type Issue struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	Description string `json:"description"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

func HandleRequest(ctx context.Context, event Payload) (Response, error) {
	description := event.Issue.Fields.Description
	if description == "" {
		return Response{Body: "Empty Description in JSON", StatusCode: 404}, nil
	}

	return Response{Body: description, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
