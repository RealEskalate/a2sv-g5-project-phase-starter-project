package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/domain/mocks"
	"context"
	"errors"
	"testing"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogCommentUseCaseTestSuite struct {
	suite.Suite
	commentUseCase    domain.CommentUseCase
	commentRepoMock   *mocks.CommentRepository
	blogRepoMock      *mocks.BlogRepository 
	timeout           time.Duration
}

func (suite *BlogCommentUseCaseTestSuite) SetupTest() {
	suite.commentRepoMock = new(mocks.CommentRepository)
	suite.blogRepoMock = new(mocks.BlogRepository)
	suite.timeout = time.Second * 2
	suite.commentUseCase = NewCommentUseCase(suite.commentRepoMock, suite.blogRepoMock, suite.timeout)
}

func (suite *BlogCommentUseCaseTestSuite) TestCreate_ValidComment() {
	commentRequest := &domain.CommentRequest{
		UserID:  "user123",
		BlogID:  "blog123",
		Content: "This is a valid comment.",
	}

	comment := &domain.Comment{
		CommentID: primitive.NewObjectID(),
		UserID:    commentRequest.UserID,
		BlogID:    commentRequest.BlogID,
		Content:   commentRequest.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.commentRepoMock.On("Create", mock.Anything, mock.Anything).Return(comment, nil)
	suite.blogRepoMock.On("UpdateCommentCount", mock.Anything, comment.BlogID, true).Return(nil)

	createdComment, err := suite.commentUseCase.Create(context.TODO(), commentRequest)

	suite.NoError(err)
	suite.Equal(comment, createdComment)
	suite.commentRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogCommentUseCaseTestSuite) TestCreate_InvalidContent() {
	commentRequest := &domain.CommentRequest{
		UserID:  "user123",
		BlogID:  "blog123",
		Content: "Short",
	}

	createdComment, err := suite.commentUseCase.Create(context.TODO(), commentRequest)

	suite.Error(err)
	suite.Nil(createdComment)
	suite.EqualError(err, "comment content too short")
}

func (suite *BlogCommentUseCaseTestSuite) TestDelete_ValidComment() {
	comment := &domain.Comment{
		CommentID: primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		Content:   "This is a comment.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.commentRepoMock.On("Delete", mock.Anything, comment.CommentID.Hex()).Return(comment, nil)
	suite.blogRepoMock.On("UpdateCommentCount", mock.Anything, comment.BlogID, false).Return(nil)

	deletedComment, err := suite.commentUseCase.Delete(context.TODO(), comment.CommentID.Hex())

	suite.NoError(err)
	suite.Equal(comment, deletedComment)
	suite.commentRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogCommentUseCaseTestSuite) TestDelete_CommentNotFound() {
	suite.commentRepoMock.On("Delete", mock.Anything, mock.Anything).Return(nil, errors.New("comment not found"))

	deletedComment, err := suite.commentUseCase.Delete(context.TODO(), "invalidID")

	suite.Error(err)
	suite.Nil(deletedComment)
	suite.EqualError(err, "comment not found")
}

func (suite *BlogCommentUseCaseTestSuite) TestGetCommentByID_ValidID() {
	comment := &domain.Comment{
		CommentID: primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		Content:   "This is a comment.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.commentRepoMock.On("GetCommentByID", mock.Anything, comment.CommentID.Hex()).Return(comment, nil)

	returnedComment, err := suite.commentUseCase.GetCommentByID(context.TODO(), comment.CommentID.Hex())

	suite.NoError(err)
	suite.Equal(comment, returnedComment)
	suite.commentRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogCommentUseCaseTestSuite) TestUpdate_ValidUpdate() {
	comment := &domain.Comment{
		CommentID: primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		Content:   "Updated comment.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.commentRepoMock.On("Update", mock.Anything, comment.Content, comment.CommentID.Hex()).Return(comment, nil)

	updatedComment, err := suite.commentUseCase.Update(context.TODO(), comment.Content, comment.CommentID.Hex())

	suite.NoError(err)
	suite.Equal(comment, updatedComment)
	suite.commentRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogCommentUseCaseTestSuite) TestUpdate_InvalidContent() {
	commentID := "someID"

	updatedComment, err := suite.commentUseCase.Update(context.TODO(), "Short", commentID)

	suite.Error(err)
	suite.Nil(updatedComment)
	suite.EqualError(err, "content too short")
}

func TestBlogCommentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogCommentUseCaseTestSuite))
}
