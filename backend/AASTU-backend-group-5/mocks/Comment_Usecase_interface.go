// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// Comment_Usecase_interface is an autogenerated mock type for the Comment_Usecase_interface type
type Comment_Usecase_interface struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: post_id, user_id
func (_m *Comment_Usecase_interface) CreateComment(post_id string, user_id string) error {
	ret := _m.Called(post_id, user_id)

	if len(ret) == 0 {
		panic("no return value specified for CreateComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(post_id, user_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: comment_id
func (_m *Comment_Usecase_interface) DeleteComment(comment_id string) error {
	ret := _m.Called(comment_id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(comment_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetComments provides a mock function with given fields: post_id
func (_m *Comment_Usecase_interface) GetComments(post_id string) ([]domain.Comment, error) {
	ret := _m.Called(post_id)

	if len(ret) == 0 {
		panic("no return value specified for GetComments")
	}

	var r0 []domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Comment, error)); ok {
		return rf(post_id)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Comment); ok {
		r0 = rf(post_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(post_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateComment provides a mock function with given fields: comment_id
func (_m *Comment_Usecase_interface) UpdateComment(comment_id string) error {
	ret := _m.Called(comment_id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(comment_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewComment_Usecase_interface creates a new instance of Comment_Usecase_interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComment_Usecase_interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Comment_Usecase_interface {
	mock := &Comment_Usecase_interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
