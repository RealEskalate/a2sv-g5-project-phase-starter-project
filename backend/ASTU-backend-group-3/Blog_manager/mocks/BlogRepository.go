// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	Domain "ASTU-backend-group-3/Blog_manager/Domain"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// BlogRepository is an autogenerated mock type for the BlogRepository type
type BlogRepository struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: blogID, comment
func (_m *BlogRepository) AddComment(blogID string, comment Domain.Comment) error {
	ret := _m.Called(blogID, comment)

	if len(ret) == 0 {
		panic("no return value specified for AddComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, Domain.Comment) error); ok {
		r0 = rf(blogID, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlogByID provides a mock function with given fields: id
func (_m *BlogRepository) DeleteBlogByID(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlogByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlogs provides a mock function with given fields: tags, startDate, endDate, sortBy
func (_m *BlogRepository) FilterBlogs(tags []string, startDate time.Time, endDate time.Time, sortBy string) ([]Domain.Blog, error) {
	ret := _m.Called(tags, startDate, endDate, sortBy)

	if len(ret) == 0 {
		panic("no return value specified for FilterBlogs")
	}

	var r0 []Domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func([]string, time.Time, time.Time, string) ([]Domain.Blog, error)); ok {
		return rf(tags, startDate, endDate, sortBy)
	}
	if rf, ok := ret.Get(0).(func([]string, time.Time, time.Time, string) []Domain.Blog); ok {
		r0 = rf(tags, startDate, endDate, sortBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, time.Time, time.Time, string) error); ok {
		r1 = rf(tags, startDate, endDate, sortBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *BlogRepository) FindByID(id string) (*Domain.Blog, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *Domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*Domain.Blog, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *Domain.Blog); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementViewCount provides a mock function with given fields: blogID
func (_m *BlogRepository) IncrementViewCount(blogID string) error {
	ret := _m.Called(blogID)

	if len(ret) == 0 {
		panic("no return value specified for IncrementViewCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(blogID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveBlogs provides a mock function with given fields: page, pageSize, sortBy
func (_m *BlogRepository) RetrieveBlogs(page int, pageSize int, sortBy string) ([]Domain.Blog, int64, error) {
	ret := _m.Called(page, pageSize, sortBy)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBlogs")
	}

	var r0 []Domain.Blog
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]Domain.Blog, int64, error)); ok {
		return rf(page, pageSize, sortBy)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []Domain.Blog); ok {
		r0 = rf(page, pageSize, sortBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, pageSize, sortBy)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, pageSize, sortBy)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Save provides a mock function with given fields: blog
func (_m *BlogRepository) Save(blog *Domain.Blog) (*Domain.Blog, error) {
	ret := _m.Called(blog)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *Domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(*Domain.Blog) (*Domain.Blog, error)); ok {
		return rf(blog)
	}
	if rf, ok := ret.Get(0).(func(*Domain.Blog) *Domain.Blog); ok {
		r0 = rf(blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(*Domain.Blog) error); ok {
		r1 = rf(blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBlogs provides a mock function with given fields: title, author, tags
func (_m *BlogRepository) SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error) {
	ret := _m.Called(title, author, tags)

	if len(ret) == 0 {
		panic("no return value specified for SearchBlogs")
	}

	var r0 []Domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, []string) ([]Domain.Blog, error)); ok {
		return rf(title, author, tags)
	}
	if rf, ok := ret.Get(0).(func(string, string, []string) []Domain.Blog); ok {
		r0 = rf(title, author, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, []string) error); ok {
		r1 = rf(title, author, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToggleDislike provides a mock function with given fields: blogID, username
func (_m *BlogRepository) ToggleDislike(blogID string, username string) error {
	ret := _m.Called(blogID, username)

	if len(ret) == 0 {
		panic("no return value specified for ToggleDislike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(blogID, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ToggleLike provides a mock function with given fields: blogID, username
func (_m *BlogRepository) ToggleLike(blogID string, username string) error {
	ret := _m.Called(blogID, username)

	if len(ret) == 0 {
		panic("no return value specified for ToggleLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(blogID, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlogRepository creates a new instance of BlogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogRepository {
	mock := &BlogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
