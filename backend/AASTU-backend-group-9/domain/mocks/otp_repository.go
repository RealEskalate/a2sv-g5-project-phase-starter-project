// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OTPRepository is an autogenerated mock type for the OTPRepository type
type OTPRepository struct {
	mock.Mock
}

// DeleteOTP provides a mock function with given fields: c, email
func (_m *OTPRepository) DeleteOTP(c context.Context, email string) error {
	ret := _m.Called(c, email)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOTPByEmail provides a mock function with given fields: ctx, email
func (_m *OTPRepository) GetOTPByEmail(ctx context.Context, email string) (*domain.OTP, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetOTPByEmail")
	}

	var r0 *domain.OTP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.OTP, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.OTP); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OTP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOTP provides a mock function with given fields: c, otp
func (_m *OTPRepository) SaveOTP(c context.Context, otp *domain.OTP) error {
	ret := _m.Called(c, otp)

	if len(ret) == 0 {
		panic("no return value specified for SaveOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.OTP) error); ok {
		r0 = rf(c, otp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOTPRepository creates a new instance of OTPRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOTPRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OTPRepository {
	mock := &OTPRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}