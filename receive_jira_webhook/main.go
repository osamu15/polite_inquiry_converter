package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

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

func HandleRequest(ctx context.Context, event Payload) (Response, error) {
	log.Println(event)
	description := event.Issue.Description
	if description == "" {
		return Response{Body: "Empty Description in JSON", StatusCode: 400}, nil
	}

	return Response{Body: description, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
