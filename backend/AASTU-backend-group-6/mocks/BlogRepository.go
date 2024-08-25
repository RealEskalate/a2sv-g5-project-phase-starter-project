// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/Domain"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// BlogRepository is an autogenerated mock type for the BlogRepository type
type BlogRepository struct {
	mock.Mock
}

// CommentOnBlog provides a mock function with given fields: user_id, comment
func (_m *BlogRepository) CommentOnBlog(user_id string, comment domain.Comment) error {
	ret := _m.Called(user_id, comment)

	if len(ret) == 0 {
		panic("no return value specified for CommentOnBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Comment) error); ok {
		r0 = rf(user_id, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBlog provides a mock function with given fields: user_id, blog
func (_m *BlogRepository) CreateBlog(user_id string, blog domain.Blog) (domain.Blog, error) {
	ret := _m.Called(user_id, blog)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlog")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.Blog) (domain.Blog, error)); ok {
		return rf(user_id, blog)
	}
	if rf, ok := ret.Get(0).(func(string, domain.Blog) domain.Blog); ok {
		r0 = rf(user_id, blog)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(string, domain.Blog) error); ok {
		r1 = rf(user_id, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBlogByID provides a mock function with given fields: user_id, blog_id
func (_m *BlogRepository) DeleteBlogByID(user_id string, blog_id string) domain.ErrorResponse {
	ret := _m.Called(user_id, blog_id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlogByID")
	}

	var r0 domain.ErrorResponse
	if rf, ok := ret.Get(0).(func(string, string) domain.ErrorResponse); ok {
		r0 = rf(user_id, blog_id)
	} else {
		r0 = ret.Get(0).(domain.ErrorResponse)
	}

	return r0
}

// FilterBlogsByTag provides a mock function with given fields: tags, pageNo, pageSize, startDate, endDate, popularity
func (_m *BlogRepository) FilterBlogsByTag(tags []string, pageNo int64, pageSize int64, startDate time.Time, endDate time.Time, popularity string) ([]domain.Blog, domain.Pagination, error) {
	ret := _m.Called(tags, pageNo, pageSize, startDate, endDate, popularity)

	if len(ret) == 0 {
		panic("no return value specified for FilterBlogsByTag")
	}

	var r0 []domain.Blog
	var r1 domain.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func([]string, int64, int64, time.Time, time.Time, string) ([]domain.Blog, domain.Pagination, error)); ok {
		return rf(tags, pageNo, pageSize, startDate, endDate, popularity)
	}
	if rf, ok := ret.Get(0).(func([]string, int64, int64, time.Time, time.Time, string) []domain.Blog); ok {
		r0 = rf(tags, pageNo, pageSize, startDate, endDate, popularity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, int64, int64, time.Time, time.Time, string) domain.Pagination); ok {
		r1 = rf(tags, pageNo, pageSize, startDate, endDate, popularity)
	} else {
		r1 = ret.Get(1).(domain.Pagination)
	}

	if rf, ok := ret.Get(2).(func([]string, int64, int64, time.Time, time.Time, string) error); ok {
		r2 = rf(tags, pageNo, pageSize, startDate, endDate, popularity)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBlogByID provides a mock function with given fields: blog_id, isCalled
func (_m *BlogRepository) GetBlogByID(blog_id string, isCalled bool) (domain.Blog, error) {
	ret := _m.Called(blog_id, isCalled)

	if len(ret) == 0 {
		panic("no return value specified for GetBlogByID")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (domain.Blog, error)); ok {
		return rf(blog_id, isCalled)
	}
	if rf, ok := ret.Get(0).(func(string, bool) domain.Blog); ok {
		r0 = rf(blog_id, isCalled)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(blog_id, isCalled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogs provides a mock function with given fields: pageNo, pageSize, popularity
func (_m *BlogRepository) GetBlogs(pageNo int64, pageSize int64, popularity string) ([]domain.Blog, domain.Pagination, error) {
	ret := _m.Called(pageNo, pageSize, popularity)

	if len(ret) == 0 {
		panic("no return value specified for GetBlogs")
	}

	var r0 []domain.Blog
	var r1 domain.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, int64, string) ([]domain.Blog, domain.Pagination, error)); ok {
		return rf(pageNo, pageSize, popularity)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, string) []domain.Blog); ok {
		r0 = rf(pageNo, pageSize, popularity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, string) domain.Pagination); ok {
		r1 = rf(pageNo, pageSize, popularity)
	} else {
		r1 = ret.Get(1).(domain.Pagination)
	}

	if rf, ok := ret.Get(2).(func(int64, int64, string) error); ok {
		r2 = rf(pageNo, pageSize, popularity)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetMyBlogByID provides a mock function with given fields: user_id, blog_id
func (_m *BlogRepository) GetMyBlogByID(user_id string, blog_id string) (domain.Blog, error) {
	ret := _m.Called(user_id, blog_id)

	if len(ret) == 0 {
		panic("no return value specified for GetMyBlogByID")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (domain.Blog, error)); ok {
		return rf(user_id, blog_id)
	}
	if rf, ok := ret.Get(0).(func(string, string) domain.Blog); ok {
		r0 = rf(user_id, blog_id)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(user_id, blog_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyBlogs provides a mock function with given fields: user_id, pageNo, pageSize, popularity
func (_m *BlogRepository) GetMyBlogs(user_id string, pageNo int64, pageSize int64, popularity string) ([]domain.Blog, domain.Pagination, error) {
	ret := _m.Called(user_id, pageNo, pageSize, popularity)

	if len(ret) == 0 {
		panic("no return value specified for GetMyBlogs")
	}

	var r0 []domain.Blog
	var r1 domain.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(string, int64, int64, string) ([]domain.Blog, domain.Pagination, error)); ok {
		return rf(user_id, pageNo, pageSize, popularity)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int64, string) []domain.Blog); ok {
		r0 = rf(user_id, pageNo, pageSize, popularity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64, int64, string) domain.Pagination); ok {
		r1 = rf(user_id, pageNo, pageSize, popularity)
	} else {
		r1 = ret.Get(1).(domain.Pagination)
	}

	if rf, ok := ret.Get(2).(func(string, int64, int64, string) error); ok {
		r2 = rf(user_id, pageNo, pageSize, popularity)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserRoleByID provides a mock function with given fields: id
func (_m *BlogRepository) GetUserRoleByID(id string) (string, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserRoleByID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReactOnBlog provides a mock function with given fields: user_id, reactionType, blog_id
func (_m *BlogRepository) ReactOnBlog(user_id string, reactionType bool, blog_id string) domain.ErrorResponse {
	ret := _m.Called(user_id, reactionType, blog_id)

	if len(ret) == 0 {
		panic("no return value specified for ReactOnBlog")
	}

	var r0 domain.ErrorResponse
	if rf, ok := ret.Get(0).(func(string, bool, string) domain.ErrorResponse); ok {
		r0 = rf(user_id, reactionType, blog_id)
	} else {
		r0 = ret.Get(0).(domain.ErrorResponse)
	}

	return r0
}

// SearchBlogByTitleAndAuthor provides a mock function with given fields: title, author, pageNo, pageSize, popularity
func (_m *BlogRepository) SearchBlogByTitleAndAuthor(title string, author string, pageNo int64, pageSize int64, popularity string) ([]domain.Blog, domain.Pagination, error) {
	ret := _m.Called(title, author, pageNo, pageSize, popularity)

	if len(ret) == 0 {
		panic("no return value specified for SearchBlogByTitleAndAuthor")
	}

	var r0 []domain.Blog
	var r1 domain.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, int64, int64, string) ([]domain.Blog, domain.Pagination, error)); ok {
		return rf(title, author, pageNo, pageSize, popularity)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, int64, string) []domain.Blog); ok {
		r0 = rf(title, author, pageNo, pageSize, popularity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, int64, string) domain.Pagination); ok {
		r1 = rf(title, author, pageNo, pageSize, popularity)
	} else {
		r1 = ret.Get(1).(domain.Pagination)
	}

	if rf, ok := ret.Get(2).(func(string, string, int64, int64, string) error); ok {
		r2 = rf(title, author, pageNo, pageSize, popularity)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateBlogByID provides a mock function with given fields: user_id, blog_id, blog
func (_m *BlogRepository) UpdateBlogByID(user_id string, blog_id string, blog domain.Blog) (domain.Blog, error) {
	ret := _m.Called(user_id, blog_id, blog)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlogByID")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, domain.Blog) (domain.Blog, error)); ok {
		return rf(user_id, blog_id, blog)
	}
	if rf, ok := ret.Get(0).(func(string, string, domain.Blog) domain.Blog); ok {
		r0 = rf(user_id, blog_id, blog)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(string, string, domain.Blog) error); ok {
		r1 = rf(user_id, blog_id, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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