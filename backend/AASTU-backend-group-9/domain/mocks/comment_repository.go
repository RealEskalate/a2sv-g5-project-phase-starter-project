// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentRepository is an autogenerated mock type for the CommentRepository type
type CommentRepository struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: ctx, id, userID, comment
func (_m *CommentRepository) AddComment(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	ret := _m.Called(ctx, id, userID, comment)

	if len(ret) == 0 {
		panic("no return value specified for AddComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID, *domain.Comment) error); ok {
		r0 = rf(ctx, id, userID, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: ctx, post_id, comment_id, userID
func (_m *CommentRepository) DeleteComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID) error {
	ret := _m.Called(ctx, post_id, comment_id, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(ctx, post_id, comment_id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetComments provides a mock function with given fields: ctx, post_id
func (_m *CommentRepository) GetComments(ctx context.Context, post_id primitive.ObjectID) (*domain.Comment, error) {
	ret := _m.Called(ctx, post_id)

	if len(ret) == 0 {
		panic("no return value specified for GetComments")
	}

	var r0 *domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (*domain.Comment, error)); ok {
		return rf(ctx, post_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Comment); ok {
		r0 = rf(ctx, post_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, post_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateComment provides a mock function with given fields: ctx, post_id, comment_id, userID, comment
func (_m *CommentRepository) UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	ret := _m.Called(ctx, post_id, comment_id, userID, comment)

	if len(ret) == 0 {
		panic("no return value specified for UpdateComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID, primitive.ObjectID, *domain.Comment) error); ok {
		r0 = rf(ctx, post_id, comment_id, userID, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommentRepository creates a new instance of CommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentRepository {
	mock := &CommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}