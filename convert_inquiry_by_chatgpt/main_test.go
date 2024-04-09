package main

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/osamu15/polite_inquiry_converter/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	prompt       = "内容を変えずに文章を丁寧な表現に変更かつ要約してください。"
	inquiryText  = "ポイントが反映されていないんだけど。ちゃんと調査しろよカス"
	expectedText = "ポイントが反映されていません。ちゃんと調査してください。"
)

func TestHandleRequest(t *testing.T) {
	t.Run("Convert Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockConvertInquiryByChatGPT := mock_main.NewMockConvertInquiryByChatGPT(ctrl)
		mockConvertInquiryByChatGPT.EXPECT().ConvertTextByChatGPT(prompt, inquiryText).Return(expectedText, nil)

		request := Request{Body: inquiryText}
		expectedResponse := Response{Body: expectedText, StatusCode: 200}

		response, err := HandleRequest(context.Background(), request, mockConvertInquiryByChatGPT)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Unable to Convert Text by ChatGPT", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockConvertInquiryByChatGPT := mock_main.NewMockConvertInquiryByChatGPT(ctrl)
		expectedErr := errors.New("unable to convert text by ChatGPT")
		mockConvertInquiryByChatGPT.EXPECT().ConvertTextByChatGPT(prompt, inquiryText).Return("", expectedErr)

		request := Request{Body: inquiryText}
		expectedResponse := Response{Body: "Unable to Convert Text by ChatGPT", StatusCode: 404}

		response, err := HandleRequest(context.Background(), request, mockConvertInquiryByChatGPT)

		assert.Error(t, err)
		assert.Equal(t, expectedResponse, response)
		assert.Equal(t, expectedErr, err)
	})
}

func TestConvertTextByChatGPT(t *testing.T) {
	t.Run("Convert Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockConvertInquiryByChatGPT := mock_main.NewMockConvertInquiryByChatGPT(ctrl)
		mockConvertInquiryByChatGPT.EXPECT().ConvertTextByChatGPT(prompt, inquiryText).Return(expectedText, nil)

		convertedText, err := mockConvertInquiryByChatGPT.ConvertTextByChatGPT(prompt, inquiryText)

		assert.NoError(t, err)
		assert.Equal(t, expectedText, convertedText)
	})

	t.Run("No Valid Response Found In the Choices", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockConvertInquiryByChatGPT := mock_main.NewMockConvertInquiryByChatGPT(ctrl)
		expectedErr := errors.New("no valid response found in the choices")
		mockConvertInquiryByChatGPT.EXPECT().ConvertTextByChatGPT(prompt, inquiryText).Return("", expectedErr)

		convertedText, err := mockConvertInquiryByChatGPT.ConvertTextByChatGPT(prompt, inquiryText)

		assert.Error(t, err)
		assert.Equal(t, "", convertedText)
		assert.Equal(t, expectedErr, err)
	})
}
