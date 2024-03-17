package main

import (
	"context"
	"encoding/json"
	"log"

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

func HandleRequest(ctx context.Context, event json.RawMessage) (Response, error) {
	var inputData interface{}
	err := json.Unmarshal(event, &inputData)
	if err != nil {
		log.Println("Error unmarshalling event:", err)
		return Response{Body: "Error parsing input JSON", StatusCode: 400}, nil
	}

	log.Println(event)
	description := inputData.(map[string]interface{})["parameters"].(map[string]interface{})["Payload"].(map[string]interface{})["issue"].(map[string]interface{})["description"].(string)
	if description == "" {
		return Response{Body: "Empty Description in JSON", StatusCode: 400}, nil
	}

	return Response{Body: description, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
