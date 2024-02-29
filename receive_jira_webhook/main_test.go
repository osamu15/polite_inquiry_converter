package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	type args struct {
		context context.Context
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name            string
		args            args
		wantStatusCode  int
		wantErr         bool
		wantResponseBody string
	}{
		{
			name: "Valid Request",
			args: args{
				context: context.Background(),
				request: events.APIGatewayProxyRequest{
					Body: `{"issue": {"fields": {"description": "Test description"}}}`,
				},
			},
			wantStatusCode:  200,
			wantErr:         false,
			wantResponseBody: "Test description",
		},
        {
            name: "Invalid JSON",
			args: args{
				context: context.Background(),
				request: events.APIGatewayProxyRequest{
					Body: `{invalid_json}`,
				},
			},
			wantStatusCode: 400,
			wantErr: true,
			wantResponseBody: "Unable to parse JSON from request",
        },
		{
			name: "Empty Description in JSON",
			args: args{
				context: context.Background(),
				request: events.APIGatewayProxyRequest{
					Body: `{"issue": {"fields": { "description": "" }}}`,
				},
			},
			wantStatusCode: 400,
			wantErr: false,
			wantResponseBody: "Empty Description in JSON",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(tt.args.context, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.StatusCode != tt.wantStatusCode {
				t.Errorf("HandleRequest() StatusCode = %v, want %v", got.StatusCode, tt.wantStatusCode)
			}
			if got.Body != tt.wantResponseBody {
				t.Errorf("HandleRequest() Body = %v, want %v", got.Body, tt.wantResponseBody)
			}
		})
	}
}
