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

type JiraIssue struct {
	Fields Field `json:"fields"`
}

type Field struct {
	Project     Project   `json:"project"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Issuetype   Issuetype `json:"issuetype"`
}

type Project struct {
	Key string `json:"key"`
}

type Issuetype struct {
	Name string `json:"name"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	jiraIssue := JiraIssue{
		Fields: Field{
			Project:     Project{Key: "JPWN"},
			Summary:     "お問い合わせ内容の要約",
			Description: request.Body,
			Issuetype:   Issuetype{Name: "Task"},
		},
	}
	jiraURL := os.Getenv("JIRA_URL")
	jiraEmail := os.Getenv("JIRA_EMAIL")
	apiToken := os.Getenv("JIRA_API_TOKEN")
	err := CreateJiraIssue(jiraIssue, jiraURL, jiraEmail, apiToken)
	if err != nil {
		return Response{Body: "Unable to Create Jira Issue", StatusCode: 500}, err
	}

	return Response{Body: "Jira Issue Created", StatusCode: 200}, nil
}

func CreateJiraIssue(jiraIssue JiraIssue, jiraURL, jiraEmail, apiToken string) error {
	requestBody, err := json.Marshal(jiraIssue)
	if err != nil {
		return fmt.Errorf("error marshalling Jira request: %w", err)
	}

	req, err := http.NewRequest("POST", jiraURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("creating Jira request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(jiraEmail, apiToken)

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
