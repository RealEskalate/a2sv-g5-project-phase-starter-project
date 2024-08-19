package test

import (
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
)

type BlogUseCaseTestSuite struct {
	suite.Suite
	mockBlogRepo *mocks.Blog_Rerpository_interface
	blogUseCase  *usecase.BlogUseCase
}

func (suite *BlogUseCaseTestSuite) SetupTest() {
	suite.mockBlogRepo = mocks.NewBlog_Rerpository_interface(suite.T())
	suite.blogUseCase = &usecase.BlogUseCase{
		BlogRepo: suite.mockBlogRepo,
	}
}

func (suite *BlogUseCaseTestSuite) TestCreateBlog() {
	blog := domain.Blog{Title: "Test Blog", Content: "Test Content"}

	suite.mockBlogRepo.On("CreateBlogDocunent", blog).Return(blog, nil)

	createdBlog, err := suite.blogUseCase.CreateBlog(blog)
	suite.NoError(err)
	suite.Equal(blog.Title, createdBlog.Title)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetOneBlog() {
	id := "some-id"
	expectedBlogs := []domain.Blog{{Title: "Test Blog"}}

	suite.mockBlogRepo.On("GetOneBlogDocunent", id).Return(expectedBlogs, nil)

	blogs, err := suite.blogUseCase.GetOneBlog(id)
	suite.NoError(err)
	suite.Len(blogs, 1)
	suite.Equal(expectedBlogs[0].Title, blogs[0].Title)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestUpdateBlog() {
	id := "some-id"
	updatedBlog := domain.Blog{Title: "Updated Title"}

	suite.mockBlogRepo.On("UpdateBlogDocunent", id, updatedBlog).Return(updatedBlog, nil)

	result, err := suite.blogUseCase.UpdateBlog(id, updatedBlog)
	suite.NoError(err)
	suite.Equal(updatedBlog.Title, result.Title)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlog() {
	id := "some-id"

	suite.mockBlogRepo.On("DeleteBlogDocunent", id).Return(nil)

	err := suite.blogUseCase.DeleteBlog(id)
	suite.NoError(err)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func TestBlogUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseTestSuite))
}
