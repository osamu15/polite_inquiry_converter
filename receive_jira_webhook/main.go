package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type JiraWebhookRequest struct {
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

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var jiraReq JiraWebhookRequest
	log.Println(request.Body)
	err := json.Unmarshal([]byte(request.Body), &jiraReq)
	if err != nil {
		return Response{Body: "Unable to parse JSON from request", StatusCode: 400}, err
	}

	description := jiraReq.Issue.Fields.Description
	if description == "" {
		return Response{Body: "Empty Description in JSON", StatusCode: 400}, nil
	}

	return Response{Body: description, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
