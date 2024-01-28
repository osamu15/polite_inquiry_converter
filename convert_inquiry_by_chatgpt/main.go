package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
)

type Request struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statuscode"`
}

type OpenAIRequest struct {
	Model    string     `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	prompt := "## 内容を変えずに文字を丁寧な表現に変更してください。"
	inquiryText := request.Body
	convertedText, err := convertTextByChatGPT(prompt, inquiryText)
	if err != nil {
		return Response{Body: "Unable to Convert Text by ChatGPT", StatusCode: 500}, err
	}

	return Response{Body: convertedText, StatusCode: 200}, nil
}

func convertTextByChatGPT(prompt, inquiryText string) (string, error) {
	requestData := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: prompt},
			{Role: "user", Content: inquiryText},
		},
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %w", err)
	}
	req, err := http.NewRequest("POST", os.Getenv("OPENAI_URL"), bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("creating request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("executing request failed: %w", err)
	}
	defer resp.Body.Close()

	var result OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decoding response failed: %w", err)
	}

	if len(result.Choices) == 0 || result.Choices[0].Message.Content == "" {
		return "", errors.New("no valid response found in the choices")
	}

	return result.Choices[0].Message.Content, nil
}

func main() {
	lambda.Start(HandleRequest)
}
