package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"blog_project/domain"
)

// MockBlogRepository is a mock type for the BlogRepository interface
type MockBlogRepository struct {
	mock.Mock
}

// MockCursor is a mock type for the mongo.Cursor interface
type MockCursor struct {
	mock.Mock
}

// MockSingleResult is a mock type for the mongo.SingleResult interface
type MockSingleResult struct {
	mock.Mock
}

// Define the methods that will be called in the tests

func (m *MockBlogRepository) GetAllBlogs(ctx context.Context) ([]domain.Blog, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) GetBlogsByPage(ctx context.Context, offset, limit int) ([]domain.Blog, error) {
	args := m.Called(ctx, offset, limit)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) GetBlogByID(ctx context.Context, id int) (domain.Blog, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) CreateBlog(ctx context.Context, blog domain.Blog) (domain.Blog, error) {
	args := m.Called(ctx, blog)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) UpdateBlog(ctx context.Context, id int, blog domain.Blog) (domain.Blog, error) {
	args := m.Called(ctx, id, blog)
	return args.Get(0).(domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) DeleteBlog(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBlogRepository) SearchByTitle(ctx context.Context, title string) ([]domain.Blog, error) {
	args := m.Called(ctx, title)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) SearchByTags(ctx context.Context, tags []string) ([]domain.Blog, error) {
	args := m.Called(ctx, tags)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

func (m *MockBlogRepository) SearchByAuthor(ctx context.Context, author string) ([]domain.Blog, error) {
	args := m.Called(ctx, author)
	return args.Get(0).([]domain.Blog), args.Error(1)
}

// Mock methods for the cursor
func (m *MockCursor) Next(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

func (m *MockCursor) Decode(val interface{}) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Mock methods for the single result
func (m *MockSingleResult) Decode(val interface{}) error {
	args := m.Called(val)
	return args.Error(0)
}
