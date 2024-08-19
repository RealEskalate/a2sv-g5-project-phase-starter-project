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

type LikeRepoTestSuite struct {
	suite.Suite
	mockColl         *mocks.CollectionInterface
	mockCursor       *mocks.CursorInterface
	mockDeleteResult *mocks.DeleteResultInterface
	repo             *repository.LikeRepository
}

func (suite *LikeRepoTestSuite) SetupTest() {
	suite.mockColl = mocks.NewCollectionInterface(suite.T())
	suite.mockCursor = mocks.NewCursorInterface(suite.T())
	suite.mockDeleteResult = mocks.NewDeleteResultInterface(suite.T())

	suite.repo = repository.NewLikeRepository(suite.mockColl)
}

func (suite *LikeRepoTestSuite) TestGetLikes() {

	suite.mockColl.On("Find", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)

		_, postIDExists := query["post_id"]
		return ok && postIDExists
	})).Return(suite.mockCursor, nil)

	suite.mockCursor.On("Next", context.TODO()).Return(true).Once()
	suite.mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		like := args.Get(0).(*domain.Like)
		*like = domain.Like{UserID: primitive.NewObjectID(), PostID: primitive.NewObjectID()}
	}).Return(nil)

	suite.mockCursor.On("Next", context.TODO()).Return(false).Once()

	suite.mockCursor.On("Close", context.TODO()).Return(nil)

	likes, err := suite.repo.GetLikes(primitive.NewObjectID().Hex())

	suite.NoError(err)
	suite.Len(likes, 1)
	suite.NotEmpty(likes[0].UserID)
	suite.NotEmpty(likes[0].PostID)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockCursor.AssertExpectations(suite.T())
}

func (suite *LikeRepoTestSuite) TestCreateLike() {
	suite.mockColl.On("InsertOne", context.TODO(), mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	err := suite.repo.CreateLike(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *LikeRepoTestSuite) TestDeleteLike() {
	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	})).Return(suite.mockDeleteResult, nil)
	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	err := suite.repo.DeleteLike(primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func TestLikeRepoTestSuite(t *testing.T) {
	suite.Run(t, new(LikeRepoTestSuite))
}
