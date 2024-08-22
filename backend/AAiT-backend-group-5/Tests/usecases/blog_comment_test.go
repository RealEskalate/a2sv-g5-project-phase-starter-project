package usecases_test

import (
	"context"
	"testing"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type CommentUsecaseTestSuite struct {
	suite.Suite
	commentRepoMock  *mocks.MockBlogCommentRepository
	blogRepoMock     *mocks.MockBlogRepository
	cacheServiceMock *mocks.MockRedisCache
	commentUsecase   interfaces.BlogCommentUsecase
	ctr              *gomock.Controller
}

func (suite *CommentUsecaseTestSuite) SetupSuite() {
	suite.ctr = gomock.NewController(suite.T())
	suite.commentRepoMock = mocks.NewMockBlogCommentRepository(suite.ctr)
	suite.blogRepoMock = mocks.NewMockBlogRepository(suite.ctr)
	suite.cacheServiceMock = mocks.NewMockRedisCache(suite.ctr)
	suite.commentUsecase = usecases.NewCommentUsecase(
		suite.commentRepoMock,
		suite.blogRepoMock,
		suite.cacheServiceMock,
	)
}

func (suite *CommentUsecaseTestSuite) TearDownSuite() {
	suite.ctr.Finish()
}

func (suite *CommentUsecaseTestSuite) newComment() *models.Comment {
	return &models.Comment{
		ID:      "comment1",
		BlogID:  "blog1",
		UserID:  "user1",
		Content: "Great post!",
	}
}

func (suite *CommentUsecaseTestSuite) existingComment() *models.Comment {
	return &models.Comment{
		ID:      "comment1",
		BlogID:  "blog1",
		UserID:  "user1",
		Content: "Old content",
	}
}

func (suite *CommentUsecaseTestSuite) TestAddComment_Success() {
	ctx := context.Background()
	comment := *suite.newComment()

	suite.blogRepoMock.
		EXPECT().
		GetBlog(ctx, comment.BlogID).
		Return(&models.Blog{}, nil)
	suite.commentRepoMock.
		EXPECT().
		AddComment(ctx, comment).
		Return(nil)

	err := suite.commentUsecase.AddComment(ctx, comment)
	if err != nil {
		suite.T().Errorf("expected no error, got %v", err)
	}
}

func (suite *CommentUsecaseTestSuite) TestAddComment_BlogNotFound() {
	ctx := context.Background()
	comment := *suite.newComment()

	suite.blogRepoMock.
		EXPECT().
		GetBlog(ctx, comment.BlogID).
		Return(nil, models.NotFound("Blog not found"))

	err := suite.commentUsecase.AddComment(ctx, comment)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.NotFound("Blog not found"), err)
}

func (suite *CommentUsecaseTestSuite) TestGetComments_Success() {
	ctx := context.Background()
	blogID := "blog1"
	comments := []models.Comment{
		*suite.newComment(),
	}

	suite.blogRepoMock.
		EXPECT().
		GetBlog(ctx, blogID).
		Return(&models.Blog{}, nil)
	suite.commentRepoMock.
		EXPECT().
		GetComments(ctx, blogID).
		Return(comments, nil)

	result, err := suite.commentUsecase.GetComments(ctx, blogID)
	if err != nil {
		suite.T().Errorf("expected no error, got %v", err)
	}
	suite.Equal(comments, result)
}

func (suite *CommentUsecaseTestSuite) TestGetComments_BlogNotFound() {
	ctx := context.Background()
	blogID := "blog1"

	suite.blogRepoMock.
		EXPECT().
		GetBlog(ctx, blogID).
		Return(nil, models.NotFound("Blog not found"))

	result, err := suite.commentUsecase.GetComments(ctx, blogID)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.NotFound("Blog not found"), err)
	suite.Nil(result)
}

func (suite *CommentUsecaseTestSuite) TestGetComment_Success() {
	ctx := context.Background()
	commentID := "comment1"
	comment := suite.newComment()

	suite.commentRepoMock.
		EXPECT().
		GetComment(ctx, commentID).
		Return(comment, nil)

	result, err := suite.commentUsecase.GetComment(ctx, commentID)
	if err != nil {
		suite.T().Errorf("expected no error, got %v", err)
	}
	suite.Equal(comment, result)
}

func (suite *CommentUsecaseTestSuite) TestGetComment_CommentNotFound() {
	ctx := context.Background()
	commentID := "comment1"

	suite.commentRepoMock.
		EXPECT().
		GetComment(ctx, commentID).
		Return(nil, models.NotFound("Comment not found"))

	result, err := suite.commentUsecase.GetComment(ctx, commentID)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.NotFound("Comment not found"), err)
	suite.Nil(result)
}

func (suite *CommentUsecaseTestSuite) TestUpdateComment_Success() {
	ctx := context.Background()
	commentID := "comment1"
	userID := "user1"
	updateRequest := dtos.CommentUpdateRequest{Content: "Updated content"}

	existingComment := suite.existingComment()

	suite.commentRepoMock.
		EXPECT().
		GetComment(ctx, commentID).
		Return(existingComment, nil)
	suite.commentRepoMock.
		EXPECT().
		UpdateComment(ctx, commentID, updateRequest).
		Return(nil)

	err := suite.commentUsecase.UpdateComment(ctx, commentID, userID, updateRequest)
	if err != nil {
		suite.T().Errorf("expected no error, got %v", err)
	}
}

func (suite *CommentUsecaseTestSuite) TestUpdateComment_Unauthorized() {
	ctx := context.Background()
	commentID := "comment1"
	userID := "user2" 
	updateRequest := dtos.CommentUpdateRequest{Content: "Updated content"}

	existingComment := &models.Comment{
		ID:      commentID,
		BlogID:  "blog1",
		UserID:  "user1", 
		Content: "Old content",
	}

	suite.commentRepoMock.
		EXPECT().
		GetComment(ctx, commentID).
		Return(existingComment, nil)

	err := suite.commentUsecase.UpdateComment(ctx, commentID, userID, updateRequest)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.Unauthorized("You are not authorized to update this comment"), err)
}

func (suite *CommentUsecaseTestSuite) TestDeleteComment_Success() {
	ctx := context.Background()
	commentID := "comment1"
	userID := "user1"
	existingComment := suite.existingComment()

	suite.commentRepoMock.
		EXPECT().
		GetComment(ctx, commentID).
		Return(existingComment, nil)
	suite.commentRepoMock.
		EXPECT().
		DeleteComment(ctx, commentID).
		Return(nil)

	err := suite.commentUsecase.DeleteComment(ctx, userID, commentID)
	if err != nil {
		suite.T().Errorf("expected no error, got %v", err)
	}
}

func (suite *CommentUsecaseTestSuite) TestDeleteComment_Unauthorized() {
	ctx := context.Background()
	commentID := "comment1"
	userID := "user2" 
	existingComment := suite.existingComment()

	suite.commentRepoMock.
		EXPECT().
		GetComment(ctx, commentID).
		Return(existingComment, nil)

	err := suite.commentUsecase.DeleteComment(ctx, userID, commentID)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.Unauthorized("You are not authorized to delete this comment"), err)
}

func TestCommentUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CommentUsecaseTestSuite))
}
