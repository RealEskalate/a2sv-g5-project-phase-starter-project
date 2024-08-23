// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CursorInterface is an autogenerated mock type for the CursorInterface type
type CursorInterface struct {
	mock.Mock
}

// Close provides a mock function with given fields: _a0
func (_m *CursorInterface) Close(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Decode provides a mock function with given fields: _a0
func (_m *CursorInterface) Decode(_a0 interface{}) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields: _a0
func (_m *CursorInterface) Next(_a0 context.Context) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Next")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewCursorInterface creates a new instance of CursorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCursorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CursorInterface {
	mock := &CursorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}