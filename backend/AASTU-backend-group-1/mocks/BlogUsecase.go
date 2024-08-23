// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "blogs/domain"

	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	time "time"
)

// BlogUsecase is an autogenerated mock type for the BlogUsecase type
type BlogUsecase struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: comment
func (_m *BlogUsecase) AddComment(comment *domain.Comment) error {
	ret := _m.Called(comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Comment) error); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddLike provides a mock function with given fields: like
func (_m *BlogUsecase) AddLike(like *domain.Like) error {
	ret := _m.Called(like)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Like) error); ok {
		r0 = rf(like)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddView provides a mock function with given fields: view, claim
func (_m *BlogUsecase) AddView(view []primitive.ObjectID, claim domain.LoginClaims) error {
	ret := _m.Called(view, claim)

	var r0 error
	if rf, ok := ret.Get(0).(func([]primitive.ObjectID, domain.LoginClaims) error); ok {
		r0 = rf(view, claim)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlogByID provides a mock function with given fields: id, claim
func (_m *BlogUsecase) DeleteBlogByID(id string, claim *domain.LoginClaims) error {
	ret := _m.Called(id, claim)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.LoginClaims) error); ok {
		r0 = rf(id, claim)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: commentID, claim
func (_m *BlogUsecase) DeleteComment(commentID string, claim *domain.LoginClaims) error {
	ret := _m.Called(commentID, claim)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.LoginClaims) error); ok {
		r0 = rf(commentID, claim)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlog provides a mock function with given fields: tags, dateFrom, dateTo
func (_m *BlogUsecase) FilterBlog(tags []string, dateFrom time.Time, dateTo time.Time) ([]*domain.Blog, error) {
	ret := _m.Called(tags, dateFrom, dateTo)

	var r0 []*domain.Blog
	if rf, ok := ret.Get(0).(func([]string, time.Time, time.Time) []*domain.Blog); ok {
		r0 = rf(tags, dateFrom, dateTo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, time.Time, time.Time) error); ok {
		r1 = rf(tags, dateFrom, dateTo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAiContent provides a mock function with given fields: prompt
func (_m *BlogUsecase) GenerateAiContent(prompt string) (string, error) {
	ret := _m.Called(prompt)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prompt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogByID provides a mock function with given fields: id
func (_m *BlogUsecase) GetBlogByID(id string) (*domain.Blog, error) {
	ret := _m.Called(id)

	var r0 *domain.Blog
	if rf, ok := ret.Get(0).(func(string) *domain.Blog); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogComments provides a mock function with given fields: blogID
func (_m *BlogUsecase) GetBlogComments(blogID string) ([]*domain.Comment, error) {
	ret := _m.Called(blogID)

	var r0 []*domain.Comment
	if rf, ok := ret.Get(0).(func(string) []*domain.Comment); ok {
		r0 = rf(blogID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(blogID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogLikes provides a mock function with given fields: blogID
func (_m *BlogUsecase) GetBlogLikes(blogID string) ([]*domain.Like, error) {
	ret := _m.Called(blogID)

	var r0 []*domain.Like
	if rf, ok := ret.Get(0).(func(string) []*domain.Like); ok {
		r0 = rf(blogID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Like)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(blogID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogs provides a mock function with given fields: sortBy, page, limit, reverse
func (_m *BlogUsecase) GetBlogs(sortBy string, page int, limit int, reverse bool) ([]*domain.Blog, int, error) {
	ret := _m.Called(sortBy, page, limit, reverse)

	var r0 []*domain.Blog
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(string, int, int, bool) ([]*domain.Blog, int, error)); ok {
		return rf(sortBy, page, limit, reverse)
	}
	if rf, ok := ret.Get(0).(func(string, int, int, bool) []*domain.Blog); ok {
		r0 = rf(sortBy, page, limit, reverse)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int, bool) error); ok {
	if rf, ok := ret.Get(1).(func(string, int, int, bool) int); ok {
		r1 = rf(sortBy, page, limit, reverse)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(string, int, int, bool) error); ok {
		r2 = rf(sortBy, page, limit, reverse)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InsertBlog provides a mock function with given fields: blog
func (_m *BlogUsecase) InsertBlog(blog *domain.Blog) (*domain.Blog, error) {
	ret := _m.Called(blog)

	var r0 *domain.Blog
	if rf, ok := ret.Get(0).(func(*domain.Blog) *domain.Blog); ok {
		r0 = rf(blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Blog) error); ok {
		r1 = rf(blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveLike provides a mock function with given fields: id, claim
func (_m *BlogUsecase) RemoveLike(id string, claim *domain.LoginClaims) error {
	ret := _m.Called(id, claim)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.LoginClaims) error); ok {
		r0 = rf(id, claim)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchBlog provides a mock function with given fields: title, author, tags
func (_m *BlogUsecase) SearchBlog(title string, author string, tags []string) ([]*domain.Blog, error) {
	ret := _m.Called(title, author, tags)

	var r0 []*domain.Blog
	if rf, ok := ret.Get(0).(func(string, string, []string) []*domain.Blog); ok {
		r0 = rf(title, author, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []string) error); ok {
		r1 = rf(title, author, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlogByID provides a mock function with given fields: id, blog, claim
func (_m *BlogUsecase) UpdateBlogByID(id string, blog *domain.Blog, claim *domain.LoginClaims) (*domain.Blog, error) {
	ret := _m.Called(id, blog, claim)

	var r0 *domain.Blog
	if rf, ok := ret.Get(0).(func(string, *domain.Blog, *domain.LoginClaims) *domain.Blog); ok {
		r0 = rf(id, blog, claim)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *domain.Blog, *domain.LoginClaims) error); ok {
		r1 = rf(id, blog, claim)
	} else {
		r1 = ret.Error(1)
	}


	var r0 *domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *domain.Blog, *domain.LoginClaims) (*domain.Blog, error)); ok {
		return rf(id, blog, claim)
	}
	if rf, ok := ret.Get(0).(func(string, *domain.Blog, *domain.LoginClaims) *domain.Blog); ok {
		r0 = rf(id, blog, claim)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *domain.Blog, *domain.LoginClaims) error); ok {
		r1 = rf(id, blog, claim)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBlogUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlogUsecase creates a new instance of BlogUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlogUsecase(t mockConstructorTestingTNewBlogUsecase) *BlogUsecase {
	mock := &BlogUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
