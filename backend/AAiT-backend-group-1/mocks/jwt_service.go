// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	jwt "github.com/dgrijalva/jwt-go"

	mock "github.com/stretchr/testify/mock"
)

// JwtService is an autogenerated mock type for the JwtService type
type JwtService struct {
	mock.Mock
}

// GenerateAccessTokenWithPayload provides a mock function with given fields: user
func (_m *JwtService) GenerateAccessTokenWithPayload(user domain.User) (string, domain.Error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateAccessTokenWithPayload")
	}

	var r0 string
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(domain.User) (string, domain.Error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.User) domain.Error); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// GenerateRefreshTokenWithPayload provides a mock function with given fields: user
func (_m *JwtService) GenerateRefreshTokenWithPayload(user domain.User) (string, domain.Error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateRefreshTokenWithPayload")
	}

	var r0 string
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(domain.User) (string, domain.Error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.User) domain.Error); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// GenerateResetToken provides a mock function with given fields: email
func (_m *JwtService) GenerateResetToken(email string) (string, domain.Error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GenerateResetToken")
	}

	var r0 string
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(string) (string, domain.Error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) domain.Error); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// GenerateVerificationToken provides a mock function with given fields: user
func (_m *JwtService) GenerateVerificationToken(user domain.User) (string, domain.Error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateVerificationToken")
	}

	var r0 string
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(domain.User) (string, domain.Error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.User) domain.Error); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// RevokedToken provides a mock function with given fields: token
func (_m *JwtService) RevokedToken(token string) domain.Error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for RevokedToken")
	}

	var r0 domain.Error
	if rf, ok := ret.Get(0).(func(string) domain.Error); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Error)
		}
	}

	return r0
}

// ValidateAccessToken provides a mock function with given fields: token
func (_m *JwtService) ValidateAccessToken(token string) (*jwt.Token, domain.Error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateAccessToken")
	}

	var r0 *jwt.Token
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, domain.Error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) domain.Error); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// ValidateRefreshToken provides a mock function with given fields: token
func (_m *JwtService) ValidateRefreshToken(token string) (*jwt.Token, domain.Error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateRefreshToken")
	}

	var r0 *jwt.Token
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, domain.Error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) domain.Error); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// ValidateResetToken provides a mock function with given fields: token
func (_m *JwtService) ValidateResetToken(token string) (*jwt.Token, domain.Error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateResetToken")
	}

	var r0 *jwt.Token
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, domain.Error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) domain.Error); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// ValidateVerificationToken provides a mock function with given fields: token
func (_m *JwtService) ValidateVerificationToken(token string) (*jwt.Token, domain.Error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateVerificationToken")
	}

	var r0 *jwt.Token
	var r1 domain.Error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, domain.Error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) domain.Error); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.Error)
		}
	}

	return r0, r1
}

// NewJwtService creates a new instance of JwtService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtService(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtService {
	mock := &JwtService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
