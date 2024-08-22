package mocks

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/stretchr/testify/mock"
)

type MockBlogRepository struct {
	mock.Mock
}

func (m *MockBlogRepository) Create(blog *domain.Blog) (*domain.Blog, domain.Error) {
	args := m.Called(blog)
	blogResult := args.Get(0).(*domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogResult, customErr
		}
	}
	return blogResult, nil
}

func (m *MockBlogRepository) FindById(id string) (*domain.Blog, domain.Error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, &domain.CustomError{
			Code:    404,
			Message: "Blog not found",
		}
	}
	blogResult := args.Get(0).(*domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogResult, customErr
		}
	}
	return blogResult, nil
}

func (m *MockBlogRepository) FindAll(page_number string) ([]domain.Blog, domain.Error) {
	args := m.Called(page_number)
	blogs := args.Get(0).([]domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogs, customErr
		}
	}
	return blogs, nil
}

func (m *MockBlogRepository) Update(blogID string, blog *domain.Blog) (*domain.Blog, domain.Error) {
	args := m.Called(blogID, blog)
	blogResult := args.Get(0).(*domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogResult, customErr
		}
	}
	return blogResult, nil
}

func (m *MockBlogRepository) Delete(id string) domain.Error {
	args := m.Called(id)
	err := args.Get(0)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return customErr
		}
	}
	return nil
}

func (m *MockBlogRepository) SearchByTitle(title string, page_number string) ([]domain.Blog, domain.Error) {
	args := m.Called(title, page_number)
	blogs := args.Get(0).([]domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogs, customErr
		}
	}
	return blogs, nil
}

func (m *MockBlogRepository) SearchByAuthor(author string, page_number string) ([]domain.Blog, domain.Error) {
	args := m.Called(author, page_number)
	blogs := args.Get(0).([]domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogs, customErr
		}
	}
	return blogs, nil
}

func (m *MockBlogRepository) Filter(filters map[string]interface{}) ([]domain.Blog, domain.Error) {
	args := m.Called(filters)
	blogs := args.Get(0).([]domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogs, customErr
		}
	}
	return blogs, nil
}

func (m *MockBlogRepository) AddComment(blogID string, comment *domain.Comment) domain.Error {
	args := m.Called(blogID, comment)
	err := args.Get(0)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return customErr
		}
	}
	return nil
}

func (m *MockBlogRepository) DeleteComment(blogID, commentID string) domain.Error {
	args := m.Called(blogID, commentID)
	err := args.Get(0)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return customErr
		}
	}
	return nil
}

func (m *MockBlogRepository) EditComment(blogID, commentID string, comment *domain.Comment) domain.Error {
	args := m.Called(blogID, commentID, comment)
	err := args.Get(0)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return customErr
		}
	}
	return nil
}

func (m *MockBlogRepository) Like(blogID, userID string) domain.Error {
	args := m.Called(blogID, userID)
	err := args.Get(0)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return customErr
		}
	}
	return nil
}

func (m *MockBlogRepository) DisLike(blogID, userID string) domain.Error {
	args := m.Called(blogID, userID)
	err := args.Get(0)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return customErr
		}
	}
	return nil
}
