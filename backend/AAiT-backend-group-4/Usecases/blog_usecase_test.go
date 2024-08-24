package usecases_test

import (
	domain "aait-backend-group4/Domain"
	usecases "aait-backend-group4/Usecases"
	"aait-backend-group4/mocks"
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecaseTestSuite struct {
	suite.Suite
	mockBlogRepo    *mocks.BlogRepository
	mockUserRepo    *mocks.UserRepository
	mockRedisClient *redis.Client
	blogUsecase     domain.BlogUsecase
}

func (suite *BlogUsecaseTestSuite) SetupTest() {
	suite.mockBlogRepo = new(mocks.BlogRepository)
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.mockRedisClient = redis.NewClient(&redis.Options{})
	suite.blogUsecase = usecases.NewBlogUsecase(suite.mockBlogRepo, suite.mockUserRepo, 2*time.Second, suite.mockRedisClient)
}

func (suite *BlogUsecaseTestSuite) TestCreateBlog_Success() {
	ctx := context.Background()
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: "Test Content",
	}

	suite.mockBlogRepo.On("CreateBlog", mock.Anything, blog).Return(nil)

	err := suite.blogUsecase.CreateBlog(ctx, blog)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestFetchByBlogID_Success() {
	ctx := context.Background()
	blogID := "blogID123"
	expectedBlog := domain.Blog{
		ID:    primitive.NewObjectID(),
		Title: "Test Blog",
	}

	suite.mockBlogRepo.On("FetchByBlogID", mock.Anything, blogID).Return(expectedBlog, nil)

	blog, err := suite.blogUsecase.FetchByBlogID(ctx, blogID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedBlog, blog)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestUpdateBlog_Success() {
	ctx := context.Background()
	blogID := primitive.NewObjectID()
	updatingID := "userID123"
	blogTitle := "Updated Title"
	blogUpdate := domain.BlogUpdate{
		Title: &blogTitle,
	}

	suite.mockBlogRepo.On("BlogExists", mock.Anything, blogID).Return(true, nil)
	suite.mockBlogRepo.On("UserIsAuthor", mock.Anything, blogID, updatingID).Return(true, nil)
	suite.mockBlogRepo.On("UpdateBlog", mock.Anything, blogID, blogUpdate).Return(nil)

	err := suite.blogUsecase.UpdateBlog(ctx, blogID, blogUpdate, updatingID)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestDeleteBlog_Success() {
	ctx := context.Background()
	blogID := primitive.NewObjectID()
	deletingID := "userID123"

	suite.mockBlogRepo.On("BlogExists", mock.Anything, blogID).Return(true, nil)
	suite.mockBlogRepo.On("UserIsAuthor", mock.Anything, blogID, deletingID).Return(true, nil)
	suite.mockBlogRepo.On("DeleteBlog", mock.Anything, blogID).Return(nil)

	err := suite.blogUsecase.DeleteBlog(ctx, blogID, deletingID)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestAddComment_Success() {
	ctx := context.Background()
	blogID := "blogID123"
	comment := domain.Comment{
		Content: "Test Comment",
	}

	suite.mockBlogRepo.On("UpdateFeedback", mock.Anything, blogID, mock.AnythingOfType("func(*domain.Feedback) error")).Return(nil)

	err := suite.blogUsecase.AddComment(ctx, blogID, comment)

	assert.NoError(suite.T(), err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func TestBlogUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseTestSuite))
}
