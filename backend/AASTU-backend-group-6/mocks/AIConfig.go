// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	genai "github.com/google/generative-ai-go/genai"

	mock "github.com/stretchr/testify/mock"
)

// AIConfig is an autogenerated mock type for the AIConfig type
type AIConfig struct {
	mock.Mock
}

// Ask provides a mock function with given fields: ctx, question
func (_m *AIConfig) Ask(ctx context.Context, question string) (*genai.GenerateContentResponse, error) {
	ret := _m.Called(ctx, question)

	if len(ret) == 0 {
		panic("no return value specified for Ask")
	}

	var r0 *genai.GenerateContentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*genai.GenerateContentResponse, error)); ok {
		return rf(ctx, question)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *genai.GenerateContentResponse); ok {
		r0 = rf(ctx, question)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*genai.GenerateContentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, question)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAIConfig creates a new instance of AIConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAIConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *AIConfig {
	mock := &AIConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}