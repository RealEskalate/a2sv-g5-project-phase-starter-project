package usecase_test

import (
	"context"
	"errors"
	"time"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blogApp/internal/domain"
	"blogApp/internal/usecase/blog"
	"blogApp/mocks/repository"
)

type BlogUsecaseTestSuite struct {
	suite.Suite
	mockRepo   *mocks.BlogRepository
	blogUsecase blog.BlogUseCase
}

// SetupTest sets up the test environment
func (suite *BlogUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.BlogRepository)
	suite.blogUsecase = blog.NewBlogUseCase(suite.mockRepo)
}

func (suite *BlogUsecaseTestSuite) TestDeleteComment_Success() {
	ctx := context.Background()
	commentID := primitive.NewObjectID()
	userID := primitive.NewObjectID()

	testComment := &domain.Comment{
		ID:     commentID,
		UserID: userID,
	}

	suite.mockRepo.On("GetCommentById", mock.Anything, commentID.Hex()).Return(testComment, nil)
	suite.mockRepo.On("DeleteComment", mock.Anything, testComment.ID).Return(nil)

	err := suite.blogUsecase.DeleteComment(ctx, commentID.Hex(), userID.Hex(), "user")
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestDeleteComment_Unauthorized() {
	ctx := context.Background()
	commentID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	otherUserID := primitive.NewObjectID()

	testComment := &domain.Comment{
		ID:     commentID,
		UserID: otherUserID,
	}

	suite.mockRepo.On("GetCommentById", mock.Anything, commentID.Hex()).Return(testComment, nil)

	err := suite.blogUsecase.DeleteComment(ctx, commentID.Hex(), userID.Hex(), "user")
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "you are not authorized to delete this comment")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestDeleteComment_RepoError() {
	ctx := context.Background()
	commentID := primitive.NewObjectID()
	userID := primitive.NewObjectID()

	testComment := &domain.Comment{
		ID:     commentID,
		UserID: userID,
	}

	suite.mockRepo.On("GetCommentById", mock.Anything, commentID.Hex()).Return(testComment, nil)
	suite.mockRepo.On("DeleteComment", mock.Anything, testComment.ID).Return(errors.New("failed to delete comment"))

	err := suite.blogUsecase.DeleteComment(ctx, commentID.Hex(), userID.Hex(), "user")
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to delete comment: failed to delete comment")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestRemoveLike_Success() {
	ctx := context.Background()
	likeID := primitive.NewObjectID()
	userID := primitive.NewObjectID()

	testLike := &domain.Like{
		ID:     likeID,
		UserID: userID,
	}

	suite.mockRepo.On("GetLikeById", mock.Anything, likeID.Hex()).Return(testLike, nil)
	suite.mockRepo.On("RemoveLike", mock.Anything, testLike.ID).Return(nil)

	err := suite.blogUsecase.RemoveLike(ctx, likeID.Hex(), userID.Hex(), "user")
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestRemoveLike_Unauthorized() {
	ctx := context.Background()
	likeID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	otherUserID := primitive.NewObjectID()

	testLike := &domain.Like{
		ID:     likeID,
		UserID: otherUserID,
	}

	suite.mockRepo.On("GetLikeById", mock.Anything, likeID.Hex()).Return(testLike, nil)

	err := suite.blogUsecase.RemoveLike(ctx, likeID.Hex(), userID.Hex(), "user")
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "you are not authorized to delete this like")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestRemoveLike_RepoError() {
	ctx := context.Background()
	likeID := primitive.NewObjectID()
	userID := primitive.NewObjectID()

	testLike := &domain.Like{
		ID:     likeID,
		UserID: userID,
	}

	suite.mockRepo.On("GetLikeById", mock.Anything, likeID.Hex()).Return(testLike, nil)
	suite.mockRepo.On("RemoveLike", mock.Anything, testLike.ID).Return(errors.New("failed to delete like"))

	err := suite.blogUsecase.RemoveLike(ctx, likeID.Hex(), userID.Hex(), "user")
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to delete like: failed to delete like")
	suite.mockRepo.AssertExpectations(suite.T())
}


func (suite *BlogUsecaseTestSuite) TestCreateBlog_Success() {
	ctx := context.Background()
	authorID := primitive.NewObjectID().Hex()
	blogID := primitive.NewObjectID()
	testBlog := &domain.Blog{
		ID:        blogID,
		Title:     "Test Blog",
		Content:    []interface{}{"This is the first line of content", 123, true},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	suite.mockRepo.On("CreateBlog", ctx, mock.Anything).Return(nil)

	err := suite.blogUsecase.CreateBlog(ctx, testBlog, authorID)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestCreateBlog_NilBlog() {
	ctx := context.Background()
	authorID := primitive.NewObjectID().Hex()

	err := suite.blogUsecase.CreateBlog(ctx, nil, authorID)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "blog cannot be nil")
}

func (suite *BlogUsecaseTestSuite) TestAddComment_Success() {
	ctx := context.Background()
	commentID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	blogID := primitive.NewObjectID()

	testComment := &domain.Comment{
		ID:     commentID,
		BlogID: blogID,
		UserID: userID,
	}

	suite.mockRepo.On("AddComment", ctx, mock.Anything).Return(nil)
	suite.mockRepo.On("GetCommentById", ctx, commentID.Hex()).Return(testComment, nil)

	err := suite.blogUsecase.AddComment(ctx, testComment, userID.Hex())
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestAddComment_NilComment() {
	ctx := context.Background()
	userID := primitive.NewObjectID().Hex()

	err := suite.blogUsecase.AddComment(ctx, nil, userID)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "comment cannot be nil")
}

func (suite *BlogUsecaseTestSuite) TestAddLike_Success() {
	ctx := context.Background()
	likeID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	blogID := primitive.NewObjectID()

	testLike := &domain.Like{
		ID:     likeID,
		BlogID: blogID,
		UserID: userID,
	}

	suite.mockRepo.On("HasUserLikedBlog", ctx, userID.Hex(), blogID.Hex()).Return(false, nil)
	suite.mockRepo.On("AddLike", ctx, mock.Anything).Return(nil)

	err := suite.blogUsecase.AddLike(ctx, testLike, userID.Hex())
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestAddLike_NilLike() {
	ctx := context.Background()
	userID := primitive.NewObjectID().Hex()

	err := suite.blogUsecase.AddLike(ctx, nil, userID)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "like cannot be nil")
}

func (suite *BlogUsecaseTestSuite) TestAddView_Success() {
	ctx := context.Background()
	viewID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	blogID := primitive.NewObjectID()

	testView := &domain.View{
		ID:     viewID,
		BlogID: blogID,
		UserID: userID,
	}

	suite.mockRepo.On("HasUserViewedBlog", ctx, userID.Hex(), blogID.Hex()).Return(false, nil)
	suite.mockRepo.On("AddView", ctx, mock.Anything).Return(nil)

	err := suite.blogUsecase.AddView(ctx, testView, userID.Hex())
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestAddView_NilView() {
	ctx := context.Background()
	userID := primitive.NewObjectID().Hex()

	err := suite.blogUsecase.AddView(ctx, nil, userID)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "view cannot be nil")
}

func TestBlogUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseTestSuite))
}
