package main

import (
    "context"
    "encoding/json"
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

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var jiraReq JiraWebhookRequest
    err := json.Unmarshal([]byte(request.Body), &jiraReq)
    if err != nil {
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
    }

    description := jiraReq.Issue.Fields.Description

    responseBody := map[string]string{
        "description": description,
    }

    jsonBody, err := json.Marshal(responseBody)
    if err != nil {
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
    }

    return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: 200}, nil
}


func main() {
    lambda.Start(HandleRequest)
}
