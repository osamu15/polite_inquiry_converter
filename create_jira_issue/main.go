package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	summary := "お問い合わせ内容要約"
	description := request.Body
	err := CreateJiraIssue(summary, description)
	if err != nil {
		return Response{Body: "Unable to Create Jira Issue", StatusCode: 500}, err
	}

	return Response{Body: "Jira Issue Created", StatusCode: 200}, nil
}

func CreateJiraIssue(summary, description string) error {
	jiraURL := os.Getenv("JIRA_URL")
	apiToken := os.Getenv("JIRA_API_TOKEN")

	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"project": map[string]string{
				"key": "KOFL",
			},
			"summary":     summary,
			"description": description,
			"issuetype": map[string]string{
				"name": "Task",
			},
		},
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling Jira request: %w", err)
	}

	req, err := http.NewRequest("POST", jiraURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("creating Jira request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("executing Jira request failed: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
