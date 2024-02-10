package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/osamu15/polite_inquiry_converter/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	prompt       = "## 内容を変えずに文章を丁寧な表現に変更してください."
	inquiryText  = "ポイントが反映されていないんだけど。ちゃんと調査しろよカス"
	expectedText = "ポイントが反映されていないんですが。ちゃんと調査してください。"
)

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
