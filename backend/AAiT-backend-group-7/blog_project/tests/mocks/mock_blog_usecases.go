package mocks

import (
	"blog_project/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockBlogUsecases is a mock type for the BlogUsecases interface
type MockBlogUsecases struct {
	mock.Mock
}

func (m *MockBlogUsecases) GetAllBlogs(ctx context.Context, sortOrder string, page, limit int) ([]domain.Blog, error) {
	args := m.Called(ctx, sortOrder, page, limit)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) GetBlogByID(ctx context.Context, id int) (domain.Blog, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) CreateBlog(ctx context.Context, blog domain.Blog) (domain.Blog, error) {
	args := m.Called(ctx, blog)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) UpdateBlog(ctx context.Context, id int, updatedBlog domain.Blog) (domain.Blog, error) {
	args := m.Called(ctx, id, updatedBlog)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) DeleteBlog(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBlogUsecases) Search(ctx context.Context, author string, tags []string, title string) ([]domain.Blog, error) {
	args := m.Called(ctx, author, tags, title)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) LikeBlog(ctx context.Context, blogID int, authorID int) (domain.Blog, error) {
	args := m.Called(ctx, blogID, authorID)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) DislikeBlog(ctx context.Context, blogID int, authorID int) (domain.Blog, error) {
	args := m.Called(ctx, blogID, authorID)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) AddComent(ctx context.Context, blogID int, authorID int, content string) (domain.Blog, error) {
	args := m.Called(ctx, blogID, authorID, content)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogUsecases) AiRecommendation(ctx context.Context, content string) (string, error) {
	args := m.Called(ctx, content)
	return args.String(0), args.Error(1)
}
