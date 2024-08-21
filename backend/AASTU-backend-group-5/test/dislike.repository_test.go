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

type DislikeRepoTestSuite struct {
	suite.Suite
	mockColl         *mocks.CollectionInterface
	mockCursor       *mocks.CursorInterface
	mockDeleteResult *mocks.DeleteResultInterface
	repo             *repository.DislikeRepository
}

func (suite *DislikeRepoTestSuite) SetupTest() {
	suite.mockColl = mocks.NewCollectionInterface(suite.T())
	suite.mockCursor = mocks.NewCursorInterface(suite.T())
	suite.mockDeleteResult = mocks.NewDeleteResultInterface(suite.T())

	suite.repo = repository.NewDislikeRepository(suite.mockColl)
}

func (suite *DislikeRepoTestSuite) TestGetDisLikes() {
	postID := primitive.NewObjectID().Hex()

	// Setting up the mock for the Find method
	suite.mockColl.On("Find", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		query, ok := m.(bson.M)
		_, postIDExists := query["post_id"]
		return ok && postIDExists
	})).Return(suite.mockCursor, nil)

	// Simulating cursor behavior
	suite.mockCursor.On("Next", context.TODO()).Return(true).Once()
	suite.mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		// Simulating the decoded dislike document
		dislike := args.Get(0).(*domain.DisLike)
		*dislike = domain.DisLike{
			UserID: primitive.NewObjectID(),
			PostID: primitive.NewObjectID(),
		}
	}).Return(nil)
	suite.mockCursor.On("Next", context.TODO()).Return(false).Once()
	suite.mockCursor.On("Close", context.TODO()).Return(nil)

	// Call the repository method
	dislikes, err := suite.repo.GetDisLikes(postID)

	// Assertions
	suite.NoError(err)
	suite.Len(dislikes, 1)
	suite.NotEmpty(dislikes[0].UserID)
	suite.NotEmpty(dislikes[0].PostID)

	// Ensure expectations are met
	suite.mockColl.AssertExpectations(suite.T())
	suite.mockCursor.AssertExpectations(suite.T())
}

func (suite *DislikeRepoTestSuite) TestCreateDisLike() {
	suite.mockColl.On("InsertOne", context.TODO(), mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	err := suite.repo.CreateDisLike(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
}

func (suite *DislikeRepoTestSuite) TestDeleteDisLike() {
	suite.mockColl.On("DeleteOne", context.TODO(), mock.MatchedBy(func(m interface{}) bool {
		_, ok := m.(bson.M)
		return ok
	})).Return(suite.mockDeleteResult, nil)
	suite.mockDeleteResult.On("DeletedCount").Return(int64(1))

	err := suite.repo.DeleteDisLike(primitive.NewObjectID().Hex())
	suite.NoError(err)

	suite.mockColl.AssertExpectations(suite.T())
	suite.mockDeleteResult.AssertExpectations(suite.T())
}

func TestDislikeRepoTestSuite(t *testing.T) {
	suite.Run(t, new(DislikeRepoTestSuite))
}
