// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// AuthInterface is an autogenerated mock type for the AuthInterface type
type AuthInterface struct {
	mock.Mock
}

// AuthenticationMiddleware provides a mock function with given fields:
func (_m *AuthInterface) AuthenticationMiddleware() gin.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AuthenticationMiddleware")
	}

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// NewAuthInterface creates a new instance of AuthInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthInterface {
	mock := &AuthInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
