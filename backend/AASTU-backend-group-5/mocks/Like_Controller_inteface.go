// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Like_Controller_inteface is an autogenerated mock type for the Like_Controller_inteface type
type Like_Controller_inteface struct {
	mock.Mock
}

// CreateLike provides a mock function with given fields: ctx
func (_m *Like_Controller_inteface) CreateLike(ctx *gin.Context) {
	_m.Called(ctx)
}

// GetLikes provides a mock function with given fields: ctx
func (_m *Like_Controller_inteface) GetLikes(ctx *gin.Context) {
	_m.Called(ctx)
}

// RemoveLike provides a mock function with given fields: ctx
func (_m *Like_Controller_inteface) RemoveLike(ctx *gin.Context) {
	_m.Called(ctx)
}

// ToggleLike provides a mock function with given fields: ctx
func (_m *Like_Controller_inteface) ToggleLike(ctx *gin.Context) {
	_m.Called(ctx)
}

// NewLike_Controller_inteface creates a new instance of Like_Controller_inteface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLike_Controller_inteface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Like_Controller_inteface {
	mock := &Like_Controller_inteface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}