package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	repository "github.com/aait.backend.g5.main/backend/Repository"
)

type BlogCommentRepositorySuite struct {
	suite.Suite

	blogCommentCollection *mocks.Collection
	repository            *repository.BlogCommentRepository
}

func (suite *BlogCommentRepositorySuite) SetupTest() {
	suite.blogCommentCollection = &mocks.Collection{}
	suite.repository = &repository.BlogCommentRepository{
		BlogCommentCollection: suite.blogCommentCollection,
	}
}

func (suite *BlogCommentRepositorySuite) TestAddComment_Success() {
	comment := models.Comment{
		Content: "Test Comment",
	}

	suite.blogCommentCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)

	err := suite.repository.AddComment(context.Background(), comment)

	suite.Nil(err)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestAddComment_Failure() {
	comment := models.Comment{
		Content: "Test Comment",
	}

	suite.blogCommentCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to insert"))

	err := suite.repository.AddComment(context.Background(), comment)

	suite.NotNil(err)
	suite.Equal("Failed to create comment", err.Message)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestGetComments_Success() {
	cursor := &mocks.Cursor{}
	comments := []models.Comment{
		{ID: "1", Content: "First Comment"},
		{ID: "2", Content: "Second Comment"},
	}

	suite.blogCommentCollection.On("Find", mock.Anything, map[string]string{"blog_id": "123"}).Return(cursor, nil)
	// Mock the behavior of cursor.All to populate the passed slice pointer
	cursor.On("All", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		// Extract the pointer to the slice and set it with the test comments
		arg := args.Get(1).(*[]models.Comment)
		*arg = comments
	}).Return(nil)
	cursor.On("Close", mock.Anything).Return(nil)

	result, err := suite.repository.GetComments(context.Background(), "123")

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(len(comments), len(result))
	suite.blogCommentCollection.AssertExpectations(suite.T())
	cursor.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestGetComments_Failure() {
	suite.blogCommentCollection.On("Find", mock.Anything, map[string]string{"blog_id": "123"}).Return(nil, errors.New("Find error"))

	result, err := suite.repository.GetComments(context.Background(), "123")

	suite.NotNil(err)
	suite.Nil(result)
	suite.Equal("Failed to retrieve comments", err.Message)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestUpdateComment_Success() {
	commentUpdate := dtos.CommentUpdateRequest{
		Content: "Updated Content",
	}

	filter := bson.M{"_id": "123"}
	update := bson.M{"$set": commentUpdate}

	suite.blogCommentCollection.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil)

	err := suite.repository.UpdateComment(context.Background(), "123", commentUpdate)

	suite.Nil(err)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestUpdateComment_Failure() {
	commentUpdate := dtos.CommentUpdateRequest{
		Content: "Updated Content",
	}

	filter := bson.M{"_id": "123"}
	update := bson.M{"$set": commentUpdate}

	suite.blogCommentCollection.On("UpdateOne", mock.Anything, filter, update).Return(nil, errors.New("Update error"))

	err := suite.repository.UpdateComment(context.Background(), "123", commentUpdate)

	suite.NotNil(err)
	suite.Equal("Failed to update comment", err.Message)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestDeleteComment_Success() {
	filter := bson.M{"_id": "123"}

	suite.blogCommentCollection.On("DeleteOne", mock.Anything, filter).Return(int64(1), nil)

	err := suite.repository.DeleteComment(context.Background(), "123")

	suite.Nil(err)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestDeleteComment_Failure() {
	filter := bson.M{"_id": "123"}

	suite.blogCommentCollection.On("DeleteOne", mock.Anything, filter).Return(int64(0), errors.New("Delete error"))

	err := suite.repository.DeleteComment(context.Background(), "123")

	suite.NotNil(err)
	suite.Equal("Failed to delete comment", err.Message)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestGetComment_Success() {
	comment := models.Comment{ID: "123", Content: "Test Comment"}

	// Mock the SingleResult
	singleResult := &mocks.SingleResult{}

	// Mock the FindOne call to return the SingleResult mock
	suite.blogCommentCollection.On("FindOne", mock.Anything, bson.M{"_id": "123"}).Return(singleResult)

	// Mock the Decode method on the SingleResult
	singleResult.On("Decode", mock.AnythingOfType("*models.Comment")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Comment)
		*arg = comment
	}).Return(nil)

	result, err := suite.repository.GetComment(context.Background(), "123")

	// Assertions
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(comment.ID, result.ID)
	suite.blogCommentCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestGetComment_Failure() {
	// Mock the SingleResult
	singleResult := &mocks.SingleResult{}

	// Mock the FindOne call to return the SingleResult mock
	suite.blogCommentCollection.On("FindOne", mock.Anything, bson.M{"_id": "123"}).Return(singleResult)

	// Mock the Decode method on the SingleResult to return an error
	singleResult.On("Decode", mock.AnythingOfType("*models.Comment")).Return(errors.New("Decode error"))

	result, err := suite.repository.GetComment(context.Background(), "123")

	// Assertions
	suite.NotNil(err)
	suite.Nil(result)
	suite.Equal(err.Message, "Failed to retrieve comment")
	suite.blogCommentCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestDeleteComments_Success() {
	filter := bson.M{"blog_id": "123"}

	suite.blogCommentCollection.On("DeleteMany", mock.Anything, filter).Return(int64(1), nil)

	err := suite.repository.DeleteComments(context.Background(), "123")

	suite.Nil(err)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func (suite *BlogCommentRepositorySuite) TestDeleteComments_Failure() {
	filter := bson.M{"blog_id": "123"}

	suite.blogCommentCollection.On("DeleteMany", mock.Anything, filter).Return(int64(0), errors.New("DeleteMany error"))

	err := suite.repository.DeleteComments(context.Background(), "123")

	suite.NotNil(err)
	suite.Equal("Failed to delete comments", err.Message)
	suite.blogCommentCollection.AssertExpectations(suite.T())
}

func TestBlogCommentRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlogCommentRepositorySuite))
}
