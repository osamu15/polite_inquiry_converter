package main

import (
	"context"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	type args struct {
		context context.Context
		event   Payload
	}
	tests := []struct {
		name             string
		args             args
		wantStatusCode   int
		wantErr          bool
		wantResponseBody string
	}{
		{
			name: "Valid Request",
			args: args{
				context: context.Background(),
				event:   Payload{Issue: Issue{Fields: Fields{Description: "Test description"}}},
			},
			wantStatusCode:   200,
			wantErr:          false,
			wantResponseBody: "Test description",
		},
		{
			name: "Empty Description in JSON",
			args: args{
				context: context.Background(),
				event:   Payload{Issue: Issue{Fields: Fields{Description: ""}}},
			},
			wantStatusCode:   400,
			wantErr:          false,
			wantResponseBody: "Empty Description in JSON",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(tt.args.context, tt.args.event)
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
