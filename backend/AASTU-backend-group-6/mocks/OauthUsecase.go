// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OauthUsecase is an autogenerated mock type for the OauthUsecase type
type OauthUsecase struct {
	mock.Mock
}

// OauthCallback provides a mock function with given fields: c, query
func (_m *OauthUsecase) OauthCallback(c context.Context, query string) interface{} {
	ret := _m.Called(c, query)

	if len(ret) == 0 {
		panic("no return value specified for OauthCallback")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string) interface{}); ok {
		r0 = rf(c, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// OauthService provides a mock function with given fields:
func (_m *OauthUsecase) OauthService() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OauthService")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// NewOauthUsecase creates a new instance of OauthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOauthUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OauthUsecase {
	mock := &OauthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}