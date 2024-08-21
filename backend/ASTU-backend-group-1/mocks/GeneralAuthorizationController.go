// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// GeneralAuthorizationController is an autogenerated mock type for the GeneralAuthorizationController type
type GeneralAuthorizationController struct {
	mock.Mock
}

// AdminMiddlewareGin provides a mock function with given fields:
func (_m *GeneralAuthorizationController) AdminMiddlewareGin() gin.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AdminMiddlewareGin")
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

// AuthMiddlewareGIn provides a mock function with given fields:
func (_m *GeneralAuthorizationController) AuthMiddlewareGIn() gin.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AuthMiddlewareGIn")
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

// UserMiddlewareGin provides a mock function with given fields:
func (_m *GeneralAuthorizationController) UserMiddlewareGin() gin.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UserMiddlewareGin")
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

// NewGeneralAuthorizationController creates a new instance of GeneralAuthorizationController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGeneralAuthorizationController(t interface {
	mock.TestingT
	Cleanup(func())
}) *GeneralAuthorizationController {
	mock := &GeneralAuthorizationController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
