// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "meleket/domain"

	mock "github.com/stretchr/testify/mock"
)

// OTPUsecaseInterface is an autogenerated mock type for the OTPUsecaseInterface type
type OTPUsecaseInterface struct {
	mock.Mock
}

// ForgotPassword provides a mock function with given fields: email
func (_m *OTPUsecaseInterface) ForgotPassword(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for ForgotPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateAndSendOTP provides a mock function with given fields: user
func (_m *OTPUsecaseInterface) GenerateAndSendOTP(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateAndSendOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyOTP provides a mock function with given fields: email, otp
func (_m *OTPUsecaseInterface) VerifyOTP(email string, otp string) (*domain.OTP, error) {
	ret := _m.Called(email, otp)

	if len(ret) == 0 {
		panic("no return value specified for VerifyOTP")
	}

	var r0 *domain.OTP
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*domain.OTP, error)); ok {
		return rf(email, otp)
	}
	if rf, ok := ret.Get(0).(func(string, string) *domain.OTP); ok {
		r0 = rf(email, otp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTP)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOTPUsecaseInterface creates a new instance of OTPUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOTPUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OTPUsecaseInterface {
	mock := &OTPUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}