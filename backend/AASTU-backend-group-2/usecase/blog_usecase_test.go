package usecase_test

import (
	"blog_g2/domain"
	"blog_g2/mocks"
	"blog_g2/usecase"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// BlogUseCasetestSuite struct to hold any shared resources or setup for the tests
type BlogUseCasetestSuite struct {
	suite.Suite
	mockBlogRepository *mocks.BlogRepository
	BlogUsecase        domain.BlogUsecase
}

// SetupTest runs before each test case
func (s *BlogUseCasetestSuite) SetupTest() {
	s.mockBlogRepository = new(mocks.BlogRepository)
	s.BlogUsecase = usecase.NewBlogUsecase(s.mockBlogRepository, time.Second*2)
}

// TearDownTest runs after each test case
func (s *BlogUseCasetestSuite) TearDownTest() {
	// Clean up resources if needed
}

// TestRunSuite runs the test suite
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCasetestSuite))
	suite.Run(t, new(UserUseCasetestSuite))
}

// TestCreateBlog tests the CreateBlog method
func (s *BlogUseCasetestSuite) TestCreateBlog() {
	blog := domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog",
		Tags:    []string{"test", "blog"},
	}

	s.mockBlogRepository.On("CreateBlog", &blog).Return(nil)

	err := s.BlogUsecase.CreateBlog(context.Background(), &blog)
	s.NoError(err)
}

// TestRetrieveBlog tests the RetrieveBlog method
func (s *BlogUseCasetestSuite) TestRetrieveBlog() {
	s.mockBlogRepository.On("RetrieveBlog", 1, "popularity", "asc").Return([]domain.Blog{}, 0, nil)

	blogs, count, err := s.BlogUsecase.RetrieveBlog(context.Background(), 1, "popularity", "asc")
	s.NoError(err)
	s.Equal([]domain.Blog{}, blogs)
	s.Equal(0, count)
}

// TestUpdateBlog tests the UpdateBlog method
func (s *BlogUseCasetestSuite) TestUpdateBlog() {
	blog := domain.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog",
		Tags:    []string{"test", "blog"},
	}

	s.mockBlogRepository.On("UpdateBlog", blog, "123", true, "my-uid").Return(nil)

	err := s.BlogUsecase.UpdateBlog(context.Background(), blog, "123", true, "my-uid")
	s.NoError(err)
}

// TestDeleteBlog tests the DeleteBlog method
func (s *BlogUseCasetestSuite) TestDeleteBlog() {
	s.mockBlogRepository.On("DeleteBlog", "123", true, "my-uid").Return(nil)

	err := s.BlogUsecase.DeleteBlog(context.Background(), "123", true, "my-uid")
	s.NoError(err)
}

// TestSearchBlog tests the SearchBlog method
func (s *BlogUseCasetestSuite) TestSearchBlog() {
	s.mockBlogRepository.On("SearchBlog", "Test", "Author").Return([]domain.Blog{}, nil)

	blogs, err := s.BlogUsecase.SearchBlog(context.Background(), "Test", "Author")
	s.NoError(err)
	s.Equal([]domain.Blog{}, blogs)
}

// TestFilterBlog tests the FilterBlog method
func (s *BlogUseCasetestSuite) TestFilterBlog() {
	tags := []string{"test", "blog"}
	date := time.Now()

	s.mockBlogRepository.On("FilterBlog", tags, date).Return([]domain.Blog{}, nil)

	blogs, err := s.BlogUsecase.FilterBlog(context.Background(), tags, date)
	s.NoError(err)
	s.Equal([]domain.Blog{}, blogs)
}
