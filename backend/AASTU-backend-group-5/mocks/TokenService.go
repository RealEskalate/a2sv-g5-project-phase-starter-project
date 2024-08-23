// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// GenerateAccessToken provides a mock function with given fields: user
func (_m *TokenService) GenerateAccessToken(user domain.User) (string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateRefreshToken provides a mock function with given fields: user
func (_m *TokenService) GenerateRefreshToken(user domain.User) (string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateAccessToken provides a mock function with given fields: token
func (_m *TokenService) ValidateAccessToken(token string) (*domain.User, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateAccessToken")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateRefreshToken provides a mock function with given fields: token
func (_m *TokenService) ValidateRefreshToken(token string) (*domain.User, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateRefreshToken")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenService creates a new instance of TokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenService {
	mock := &TokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}