// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// GoogleCallBack provides a mock function with given fields: _a0, _a1
func (_m *AuthUsecase) GoogleCallBack(_a0 string, _a1 string) (*domain.User, string, string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GoogleCallBack")
	}

	var r0 *domain.User
	var r1 string
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(string, string) (*domain.User, string, string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) *domain.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) string); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(string, string) error); ok {
		r3 = rf(_a0, _a1)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GoogleLogin provides a mock function with given fields:
func (_m *AuthUsecase) GoogleLogin() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GoogleLogin")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: _a0, _a1
func (_m *AuthUsecase) LoginUser(_a0 string, _a1 string) (domain.User, string, string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 domain.User
	var r1 string
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(string, string) (domain.User, string, string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) domain.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) string); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(string, string) error); ok {
		r3 = rf(_a0, _a1)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// RefreshTokens provides a mock function with given fields: _a0
func (_m *AuthUsecase) RefreshTokens(_a0 string) (string, string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RefreshTokens")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (string, string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterUser provides a mock function with given fields: _a0
func (_m *AuthUsecase) RegisterUser(_a0 domain.RegisterUser) (domain.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.RegisterUser) (domain.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(domain.RegisterUser) domain.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(domain.RegisterUser) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthUsecase creates a new instance of AuthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthUsecase {
	mock := &AuthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}