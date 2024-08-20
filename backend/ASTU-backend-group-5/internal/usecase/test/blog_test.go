package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"blogApp/internal/domain"
	"blogApp/internal/usecase/blog"
	"blogApp/mocks/repository" 
	"github.com/stretchr/testify/mock"
)

// BlogUseCaseTestSuite defines the test suite for the BlogUseCase
type BlogUseCaseTestSuite struct {
	suite.Suite
	useCase    blog.BlogUseCase
	repoMock   *mocks.BlogRepository
	ctx        context.Context
}

// SetupTest initializes the test suite
func (suite *BlogUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(mocks.BlogRepository)
	suite.useCase = blog.NewBlogUseCase(suite.repoMock)
	suite.ctx = context.TODO()
}

// TestCreateBlog_BlogNil checks if CreateBlog returns an error when the blog is nil
func (suite *BlogUseCaseTestSuite) TestCreateBlog_BlogNil() {
	err := suite.useCase.CreateBlog(suite.ctx, nil)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "blog cannot be nil", err.Error())
}

// TestCreateBlog_RepoError checks if CreateBlog returns an error when the repository returns an error
func (suite *BlogUseCaseTestSuite) TestCreateBlog_RepoError() {
	blog := &domain.Blog{
		ID:        primitive.NewObjectID(),
		Author:    primitive.NewObjectID(),
		Title:     "Test Blog",
		Content:   []interface{}{"Sample content"},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		Tags:      []domain.BlogTag{},
	}

	expectedErr := errors.New("some repository error")

	suite.repoMock.On("CreateBlog", suite.ctx, blog).Return(expectedErr)

	err := suite.useCase.CreateBlog(suite.ctx, blog)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "failed to create blog")
	suite.repoMock.AssertExpectations(suite.T())
}

// TestCreateBlog_Success checks if CreateBlog succeeds without any errors
func (suite *BlogUseCaseTestSuite) TestCreateBlog_Success() {
	blog := &domain.Blog{
		ID:        primitive.NewObjectID(),
		Author:    primitive.NewObjectID(),
		Title:     "Test Blog",
		Content:   []interface{}{"Sample content"},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		Tags:      []domain.BlogTag{},
	}

	suite.repoMock.On("CreateBlog", suite.ctx, blog).Return(nil)

	err := suite.useCase.CreateBlog(suite.ctx, blog)

	assert.NoError(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlog_Success() {
	// Arrange
	blogID := "12345"
	suite.repoMock.On("DeleteBlog", mock.Anything, blogID).Return(nil)

	// Act
	err := suite.useCase.DeleteBlog(context.Background(), blogID)

	// Assert
	suite.NoError(err)
	suite.repoMock.AssertCalled(suite.T(), "DeleteBlog", mock.Anything, blogID)
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlog_Failure() {
	// Arrange
	blogID := "12345"
	expectedError := fmt.Errorf("some repository error")
	suite.repoMock.On("DeleteBlog", mock.Anything, blogID).Return(expectedError)

	// Act
	err := suite.useCase.DeleteBlog(context.Background(), blogID)

	// Assert
	suite.Error(err)
	suite.EqualError(err, "failed to delete blog: some repository error")
	suite.repoMock.AssertCalled(suite.T(), "DeleteBlog", mock.Anything, blogID)
}

// TestBlogUseCaseTestSuite runs the test suite
func TestBlogUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseTestSuite))
}
