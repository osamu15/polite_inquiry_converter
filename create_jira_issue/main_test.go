package main

import (
	"context"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/osamu15/polite_inquiry_converter/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandleRequest(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockJiraCreator := mock_main.NewMockJiraIssueCreator(ctrl)
		mockJiraCreator.EXPECT().CreateJiraIssue("JPWN", "お問い合わせ内容の要約", "お問い合わせ内容", "Task", os.Getenv("JIRA_URL"), os.Getenv("JIRA_EMAIL"), os.Getenv("JIRA_API_TOKEN")).Return(nil)

		request := Request{
			Body:       "お問い合わせ内容",
			StatusCode: 200,
		}

		jira := JiraIssueCreator(mockJiraCreator)
		response, err := HandleRequest(context.Background(), request, jira)

		assert.Nil(t, err)
		assert.Equal(t, Response{Body: "Jira Issue Created", StatusCode: 200}, response)
	})

	t.Run("Create Failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockJiraCreator := mock_main.NewMockJiraIssueCreator(ctrl)
		mockJiraCreator.EXPECT().CreateJiraIssue("JPWN", "お問い合わせ内容の要約", "お問い合わせ内容", "Task", os.Getenv("JIRA_URL"), os.Getenv("JIRA_EMAIL"), os.Getenv("JIRA_API_TOKEN")).Return(assert.AnError)

		request := Request{
			Body:       "お問い合わせ内容",
			StatusCode: 200,
		}

		jira := JiraIssueCreator(mockJiraCreator)
		response, err := HandleRequest(context.Background(), request, jira)

		assert.NotNil(t, err)
		assert.Equal(t, Response{Body: "Unable to Create Jira Issue", StatusCode: 500}, response)
	})
}
