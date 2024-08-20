package test

import (
	"errors"
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
)

type BlogUseCaseTestSuite struct {
	suite.Suite
	mockBlogRepo *mocks.Blog_Repository_interface
	blogUseCase  *usecase.BlogUseCase
}

func (suite *BlogUseCaseTestSuite) SetupTest() {
	suite.mockBlogRepo = mocks.NewBlog_Repository_interface(suite.T())
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
	expectedBlog := domain.Blog{Title: "Test Blog"}

	suite.mockBlogRepo.On("GetOneBlogDocunent", id).Return(expectedBlog, nil)

	blog, err := suite.blogUseCase.GetOneBlog(id)
	suite.NoError(err)
	suite.Equal(expectedBlog.Title, blog.Title)

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
	userId := "some_user_id"

	suite.mockBlogRepo.On("DeleteBlogDocument", id, userId).Return(nil)

	err := suite.blogUseCase.DeleteBlog(id, userId)
	suite.NoError(err)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogWithInvalidID() {
	id := "invalid-id"
	user_id := "some_user_id"
	expectedErr := errors.New("blog not found")

	suite.mockBlogRepo.On("DeleteBlogDocument", id, user_id).Return(expectedErr)

	err := suite.blogUseCase.DeleteBlog(id, user_id)
	suite.Error(err)
	suite.Equal(expectedErr, err)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogWithUnauthorizedUser() {
	id := "some-id"
	user_id := "unauthorized_user_id"
	expectedErr := errors.New("unauthorized to delete this blog")

	suite.mockBlogRepo.On("DeleteBlogDocument", id, user_id).Return(expectedErr)

	err := suite.blogUseCase.DeleteBlog(id, user_id)
	suite.Error(err)
	suite.Equal(expectedErr, err)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogWithEmptyID() {
	id := ""
	user_id := "some_user_id"
	expectedErr := errors.New("invalid blog id")

	suite.mockBlogRepo.On("DeleteBlogDocument", id, user_id).Return(expectedErr)

	err := suite.blogUseCase.DeleteBlog(id, user_id)
	suite.Error(err)
	suite.Equal(expectedErr, err)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogWithEmptyUserID() {
	id := "some-id"
	user_id := ""
	expectedErr := errors.New("invalid user id")

	suite.mockBlogRepo.On("DeleteBlogDocument", id, user_id).Return(expectedErr)

	err := suite.blogUseCase.DeleteBlog(id, user_id)
	suite.Error(err)
	suite.Equal(expectedErr, err)

	suite.mockBlogRepo.AssertExpectations(suite.T())
}
func TestBlogUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseTestSuite))
}
