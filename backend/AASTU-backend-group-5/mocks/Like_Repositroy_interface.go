// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// Like_Repositroy_interface is an autogenerated mock type for the Like_Repositroy_interface type
type Like_Repositroy_interface struct {
	mock.Mock
}

// CreateLike provides a mock function with given fields: user_id, post_id
func (_m *Like_Repositroy_interface) CreateLike(user_id string, post_id string) error {
	ret := _m.Called(user_id, post_id)

	if len(ret) == 0 {
		panic("no return value specified for CreateLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(user_id, post_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLikes provides a mock function with given fields: post_id
func (_m *Like_Repositroy_interface) GetLikes(post_id string) ([]domain.Like, error) {
	ret := _m.Called(post_id)

	if len(ret) == 0 {
		panic("no return value specified for GetLikes")
	}

	var r0 []domain.Like
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Like, error)); ok {
		return rf(post_id)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Like); ok {
		r0 = rf(post_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Like)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(post_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveLike provides a mock function with given fields: user_id, post_id
func (_m *Like_Repositroy_interface) RemoveLike(user_id string, post_id string) error {
	ret := _m.Called(user_id, post_id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(user_id, post_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ToggleLike provides a mock function with given fields: user_id, post_id
func (_m *Like_Repositroy_interface) ToggleLike(user_id string, post_id string) error {
	ret := _m.Called(user_id, post_id)

	if len(ret) == 0 {
		panic("no return value specified for ToggleLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(user_id, post_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLike_Repositroy_interface creates a new instance of Like_Repositroy_interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLike_Repositroy_interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Like_Repositroy_interface {
	mock := &Like_Repositroy_interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}