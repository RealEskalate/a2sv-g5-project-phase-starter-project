package test

import (
	"context"
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/repository"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepoTestSuite struct {
	suite.Suite
	mockColl         *mocks.CollectionInterface
	mockCursor       *mocks.CursorInterface
	mockSingleResult *mocks.SingleResultInterface
	mockDeleteResult *mocks.DeleteResultInterface
	repo             *repository.CommentRepository
}

func (suite *CommentRepoTestSuite) SetupTest() {
	suite.mockColl = mocks.NewCollectionInterface(suite.T())
	suite.mockCursor = mocks.NewCursorInterface(suite.T())
	suite.mockSingleResult = mocks.NewSingleResultInterface(suite.T())
	suite.mockDeleteResult = mocks.NewDeleteResultInterface(suite.T())

	suite.repo = repository.NewCommentRepository(suite.mockColl)
}

func (suite *CommentRepoTestSuite) TestGetComments() {
	postID := primitive.NewObjectID().Hex()

	suite.mockColl.On("Find", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)
		_, postIDExists := query["post_id"]
		return ok && postIDExists
	})).Return(suite.mockCursor, nil)

	suite.mockCursor.On("Next", context.TODO()).Return(true).Once()
	suite.mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		comment := args.Get(0).(*domain.Comment)
		*comment = domain.Comment{Content: "test comment"}
	}).Return(nil)
	suite.mockCursor.On("Next", context.TODO()).Return(false).Once()

	suite.mockCursor.On("Close", context.TODO()).Return(nil)

	comments, err := suite.repo.GetComments(postID)
	suite.NoError(err)
	suite.Len(comments, 1)
	suite.Equal("test comment", comments[0].Content)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockCursor.AssertExpectations(suite.T())
}

func (suite *CommentRepoTestSuite) TestCreateComment() {
	postID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	suite.mockColl.On("InsertOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		comment, ok := m.(domain.Comment)
		return ok && comment.PostID.Hex() == postID && comment.UserID.Hex() == userID
	})).Return(&mongo.InsertOneResult{}, nil)

	err := suite.repo.CreateComment(postID, userID)
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *CommentRepoTestSuite) TestDeleteComment() {
	commentID := primitive.NewObjectID().Hex()

	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)
		id, idExists := query["_id"]
		return ok && idExists && id.(primitive.ObjectID).Hex() == commentID
	})).Return(suite.mockDeleteResult, nil)

	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	err := suite.repo.DeleteComment(commentID)
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func (suite *CommentRepoTestSuite) TestDeleteCommentNotFound() {
	commentID := primitive.NewObjectID().Hex()

	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)
		id, idExists := query["_id"]
		return ok && idExists && id.(primitive.ObjectID).Hex() == commentID
	})).Return(suite.mockDeleteResult, nil)

	suite.mockDeleteResult.On("DeletedCount").Return(int64(0))

	err := suite.repo.DeleteComment(commentID)
	suite.Error(err)
	suite.EqualError(err, "no comment with this ID found")

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func (suite *CommentRepoTestSuite) TestUpdateComment() {
	commentID := primitive.NewObjectID().Hex()

	suite.mockColl.On("UpdateOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)
		id, idExists := query["_id"]
		return ok && idExists && id.(primitive.ObjectID).Hex() == commentID
	}), mock.MatchedBy(func(m interface{}) bool {
		update, ok := m.(bson.M)
		set, setExists := update["$set"].(bson.M)
		content, contentExists := set["content"]
		return ok && setExists && contentExists && content == "Updated content"
	})).Return(&mongo.UpdateResult{}, nil)

	err := suite.repo.UpdateComment(commentID)
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
}

func TestCommentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CommentRepoTestSuite))
}
