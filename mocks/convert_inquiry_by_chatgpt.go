// Code generated by MockGen. DO NOT EDIT.
// Source: convert_inquiry_by_chatgpt/main.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConvertInquiryByChatGPT is a mock of ConvertInquiryByChatGPT interface.
type MockConvertInquiryByChatGPT struct {
	ctrl     *gomock.Controller
	recorder *MockConvertInquiryByChatGPTMockRecorder
}

// MockConvertInquiryByChatGPTMockRecorder is the mock recorder for MockConvertInquiryByChatGPT.
type MockConvertInquiryByChatGPTMockRecorder struct {
	mock *MockConvertInquiryByChatGPT
}

// NewMockConvertInquiryByChatGPT creates a new mock instance.
func NewMockConvertInquiryByChatGPT(ctrl *gomock.Controller) *MockConvertInquiryByChatGPT {
	mock := &MockConvertInquiryByChatGPT{ctrl: ctrl}
	mock.recorder = &MockConvertInquiryByChatGPTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConvertInquiryByChatGPT) EXPECT() *MockConvertInquiryByChatGPTMockRecorder {
	return m.recorder
}

// ConvertTextByChatGPT mocks base method.
func (m *MockConvertInquiryByChatGPT) ConvertTextByChatGPT(prompt, inquiryText string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertTextByChatGPT", prompt, inquiryText)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertTextByChatGPT indicates an expected call of ConvertTextByChatGPT.
func (mr *MockConvertInquiryByChatGPTMockRecorder) ConvertTextByChatGPT(prompt, inquiryText interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertTextByChatGPT", reflect.TypeOf((*MockConvertInquiryByChatGPT)(nil).ConvertTextByChatGPT), prompt, inquiryText)
}
