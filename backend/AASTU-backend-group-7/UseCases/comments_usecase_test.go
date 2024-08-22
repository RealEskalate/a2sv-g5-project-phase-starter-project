package usecases_test

import (
	"blogapp/Domain"
	usecases "blogapp/UseCases"
	"blogapp/mocks"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUsecaseSuite struct {
	suite.Suite
	context        context.Context
	commentUsecase *usecases.CommentUsecase
	repo           *mocks.CommentRepository
}

func (suite *CommentUsecaseSuite) SetupTest() {
	suite.repo = new(mocks.CommentRepository)
	suite.commentUsecase = usecases.NewCommentUseCase(suite.repo)
	suite.context = context.Background()
}

func (suite *CommentUsecaseSuite) TestCommentOnPost() {
	c, _ := gin.CreateTestContext(nil)
	comment := Domain.Comment{}
	objID := primitive.NewObjectID()
	suite.repo.On("CommentOnPost", mock.Anything, mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.commentUsecase.CommentOnPost(c, &comment, objID)
	suite.Nil(err)
	suite.Equal(200, status)

}

func (suite *CommentUsecaseSuite) TestGetCommentByID() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("GetCommentByID", mock.Anything, mock.Anything).Return(&Domain.Comment{}, nil, 200)
	_, err, status := suite.commentUsecase.GetCommentByID(c, id)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *CommentUsecaseSuite) TestEditComment() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	comment := Domain.Comment{}
	suite.repo.On("EditComment", mock.Anything, mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.commentUsecase.EditComment(c, id, &comment)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *CommentUsecaseSuite) TestGetUserComments() {
	c, _ := gin.CreateTestContext(nil)
	authorID := primitive.NewObjectID()
	suite.repo.On("GetUserComments", mock.Anything, mock.Anything).Return([]*Domain.Comment{}, nil, 200)
	_, err, status := suite.commentUsecase.GetUserComments(c, authorID)
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *CommentUsecaseSuite) TestDeleteComment() {
	c, _ := gin.CreateTestContext(nil)
	id := primitive.NewObjectID()
	suite.repo.On("DeleteComment", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.commentUsecase.DeleteComment(c, id)
	suite.Nil(err)
	suite.Equal(200, status)
}

func TestCommentUsecaseSuite(t *testing.T) {
	suite.Run(t, new(CommentUsecaseSuite))
}
