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
	suite.mockColl.On("InsertOne", context.TODO(), mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	err := suite.repo.CreateComment(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *CommentRepoTestSuite) TestDeleteComment() {
	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	})).Return(suite.mockDeleteResult, nil)
	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	err := suite.repo.DeleteComment(primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func (suite *CommentRepoTestSuite) TestUpdateComment() {
	suite.mockColl.On("UpdateOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	}), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	})).Return(&mongo.UpdateResult{}, nil)

	err := suite.repo.UpdateComment(primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
}

func TestCommentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CommentRepoTestSuite))
}
