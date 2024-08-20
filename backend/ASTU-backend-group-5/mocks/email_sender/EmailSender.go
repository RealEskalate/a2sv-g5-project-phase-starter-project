// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EmailSender is an autogenerated mock type for the EmailSender type
type EmailSender struct {
	mock.Mock
}

// SendPasswordResetEmail provides a mock function with given fields: userEmail, token
func (_m *EmailSender) SendPasswordResetEmail(userEmail string, token string) error {
	ret := _m.Called(userEmail, token)

	if len(ret) == 0 {
		panic("no return value specified for SendPasswordResetEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userEmail, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendVerificationEmail provides a mock function with given fields: userEmail, token
func (_m *EmailSender) SendVerificationEmail(userEmail string, token string) error {
	ret := _m.Called(userEmail, token)

	if len(ret) == 0 {
		panic("no return value specified for SendVerificationEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userEmail, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEmailSender creates a new instance of EmailSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailSender(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailSender {
	mock := &EmailSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
