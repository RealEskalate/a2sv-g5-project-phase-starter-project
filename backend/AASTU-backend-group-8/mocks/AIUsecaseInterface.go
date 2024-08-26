// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AIUsecaseInterface is an autogenerated mock type for the AIUsecaseInterface type
type AIUsecaseInterface struct {
	mock.Mock
}

// GenerateBlogContent provides a mock function with given fields: title, tags
func (_m *AIUsecaseInterface) GenerateBlogContent(title string, tags []string) (string, error) {
	ret := _m.Called(title, tags)

	if len(ret) == 0 {
		panic("no return value specified for GenerateBlogContent")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (string, error)); ok {
		return rf(title, tags)
	}
	if rf, ok := ret.Get(0).(func(string, []string) string); ok {
		r0 = rf(title, tags)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(title, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAIUsecaseInterface creates a new instance of AIUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAIUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AIUsecaseInterface {
	mock := &AIUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}