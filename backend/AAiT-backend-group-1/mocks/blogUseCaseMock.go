package mocks

import (
	"time"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/stretchr/testify/mock"
)

type MockBlogUseCase struct {
	mock.Mock
}

// Like implements domain.BlogUseCase.
func (m *MockBlogUseCase) Like(blogId string, userID string) domain.Error {
	panic("unimplemented")
}

func (m *MockBlogUseCase) CreateBlog(blog *domain.Blog, authorID string) domain.Error {
	args := m.Called(blog, authorID)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) GetBlog(blogID string, userID string) (*domain.Blog, domain.Error) {
	args := m.Called(blogID, userID)
	blog := args.Get(0).(*domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blog, customErr
		}
	}
	return blog, nil
}

func (m *MockBlogUseCase) GetBlogs(page_number string) ([]domain.Blog, domain.Error) {
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

func (m *MockBlogUseCase) UpdateBlog(blogID string, blog *domain.Blog, userID string) domain.Error {
	args := m.Called(blogID, blog, userID)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) DeleteBlog(blogID string , userID string) domain.Error {
	args := m.Called(blogID)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) SearchBlogsByTitle(title string, page_number string) ([]domain.Blog, domain.Error) {
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

func (m *MockBlogUseCase) SearchBlogsByAuthor(author string, page_number string) ([]domain.Blog, domain.Error) {
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

func (m *MockBlogUseCase) FilterBlogs(tags []string, dateAfter time.Time, popular bool) ([]domain.Blog, domain.Error) {
	args := m.Called(tags, dateAfter, popular)
	blogs := args.Get(0).([]domain.Blog)
	err := args.Get(1)
	if err != nil {
		if customErr, ok := err.(domain.Error); ok {
			return blogs, customErr
		}
	}
	return blogs, nil
}

func (m *MockBlogUseCase) LikeBlog(userID, blogID string) domain.Error {
	args := m.Called(userID, blogID)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) DisLike(blogID, userID string) domain.Error {
	args := m.Called(blogID, userID)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) AddComment(blogID string, comment *domain.Comment) domain.Error {
	args := m.Called(blogID, comment)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) DeleteComment(blogID, commentID string) domain.Error {
	args := m.Called(blogID, commentID)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}

func (m *MockBlogUseCase) EditComment(blogID string, commentID string, comment *domain.Comment) domain.Error {
	args := m.Called(blogID, commentID, comment)
	if err, ok := args.Get(0).(domain.Error); ok {
		return err
	}
	return nil
}
