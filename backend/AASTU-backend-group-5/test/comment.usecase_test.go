package test

import (
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
)

type CommentUseCaseTestSuite struct {
	suite.Suite
	mockCommentRepo *mocks.Comment_Usecase_interface
	commentUseCase  *usecase.CommentUseCase
}

func (suite *CommentUseCaseTestSuite) SetupTest() {
	suite.mockCommentRepo = mocks.NewComment_Usecase_interface(suite.T())
	suite.commentUseCase = &usecase.CommentUseCase{
		CommentRepo: suite.mockCommentRepo,
	}
}

func (suite *CommentUseCaseTestSuite) TestGetComments() {
	postID := "post-id"
	expectedComments := []domain.Comment{{Content: "Test Comment"}}

	suite.mockCommentRepo.On("GetComments", postID).Return(expectedComments, nil)

	comments, err := suite.commentUseCase.GetComments(postID)
	suite.NoError(err)
	suite.Len(comments, 1)
	suite.Equal(expectedComments[0].Content, comments[0].Content)

	suite.mockCommentRepo.AssertExpectations(suite.T())
}

func (suite *CommentUseCaseTestSuite) TestCreateComment() {
	postID := "post-id"
	userID := "user-id"

	suite.mockCommentRepo.On("CreateComment", postID, userID).Return(nil)

	err := suite.commentUseCase.CreateComment(postID, userID)
	suite.NoError(err)

	suite.mockCommentRepo.AssertExpectations(suite.T())
}

func (suite *CommentUseCaseTestSuite) TestDeleteComment() {
	commentID := "comment-id"

	suite.mockCommentRepo.On("DeleteComment", commentID).Return(nil)

	err := suite.commentUseCase.DeleteComment(commentID)
	suite.NoError(err)

	suite.mockCommentRepo.AssertExpectations(suite.T())
}

func (suite *CommentUseCaseTestSuite) TestUpdateComment() {
	commentID := "comment-id"

	suite.mockCommentRepo.On("UpdateComment", commentID).Return(nil)

	err := suite.commentUseCase.UpdateComment(commentID)
	suite.NoError(err)

	suite.mockCommentRepo.AssertExpectations(suite.T())
}

func TestCommentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CommentUseCaseTestSuite))
}
