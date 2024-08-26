// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "meleket/domain"

	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogUsecaseInterface is an autogenerated mock type for the BlogUsecaseInterface type
type BlogUsecaseInterface struct {
	mock.Mock
}

// CreateBlogPost provides a mock function with given fields: blog
func (_m *BlogUsecaseInterface) CreateBlogPost(blog *domain.BlogPost) (interface{}, error) {
	ret := _m.Called(blog)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlogPost")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.BlogPost) (interface{}, error)); ok {
		return rf(blog)
	}
	if rf, ok := ret.Get(0).(func(*domain.BlogPost) interface{}); ok {
		r0 = rf(blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.BlogPost) error); ok {
		r1 = rf(blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBlogPost provides a mock function with given fields: id
func (_m *BlogUsecaseInterface) DeleteBlogPost(id primitive.ObjectID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlogPost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBlogPosts provides a mock function with given fields:
func (_m *BlogUsecaseInterface) GetAllBlogPosts() ([]domain.BlogPost, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllBlogPosts")
	}

	var r0 []domain.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.BlogPost, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.BlogPost); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogByID provides a mock function with given fields: id
func (_m *BlogUsecaseInterface) GetBlogByID(id primitive.ObjectID) (*domain.BlogPost, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetBlogByID")
	}

	var r0 *domain.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) (*domain.BlogPost, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *domain.BlogPost); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlogPost provides a mock function with given fields: id, blog
func (_m *BlogUsecaseInterface) UpdateBlogPost(id primitive.ObjectID, blog *domain.BlogPost) (*domain.BlogPost, error) {
	ret := _m.Called(id, blog)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlogPost")
	}

	var r0 *domain.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *domain.BlogPost) (*domain.BlogPost, error)); ok {
		return rf(id, blog)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *domain.BlogPost) *domain.BlogPost); ok {
		r0 = rf(id, blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID, *domain.BlogPost) error); ok {
		r1 = rf(id, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlogUsecaseInterface creates a new instance of BlogUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogUsecaseInterface {
	mock := &BlogUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}