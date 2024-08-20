// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// SignupRepository is an autogenerated mock type for the SignupRepository type
type SignupRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, user
func (_m *SignupRepository) Create(c context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (domain.User, error)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(c, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByEmail provides a mock function with given fields: c, email
func (_m *SignupRepository) FindUserByEmail(c context.Context, email string) (domain.User, error) {
	ret := _m.Called(c, email)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByEmail")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(c, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByResetToken provides a mock function with given fields: c, token
func (_m *SignupRepository) FindUserByResetToken(c context.Context, token string) (domain.User, error) {
	ret := _m.Called(c, token)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByResetToken")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(c, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, token)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetOTP provides a mock function with given fields: c, email, otp
func (_m *SignupRepository) SetOTP(c context.Context, email string, otp string) error {
	ret := _m.Called(c, email, otp)

	if len(ret) == 0 {
		panic("no return value specified for SetOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, email, otp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetResetToken provides a mock function with given fields: c, email, token, expiration
func (_m *SignupRepository) SetResetToken(c context.Context, email domain.ForgotPasswordRequest, token string, expiration time.Time) (domain.User, error) {
	ret := _m.Called(c, email, token, expiration)

	if len(ret) == 0 {
		panic("no return value specified for SetResetToken")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ForgotPasswordRequest, string, time.Time) (domain.User, error)); ok {
		return rf(c, email, token, expiration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ForgotPasswordRequest, string, time.Time) domain.User); ok {
		r0 = rf(c, email, token, expiration)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ForgotPasswordRequest, string, time.Time) error); ok {
		r1 = rf(c, email, token, expiration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: c, user
func (_m *SignupRepository) UpdateUser(c context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (domain.User, error)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(c, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyUser provides a mock function with given fields: c, user
func (_m *SignupRepository) VerifyUser(c context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for VerifyUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (domain.User, error)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(c, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSignupRepository creates a new instance of SignupRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSignupRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SignupRepository {
	mock := &SignupRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
