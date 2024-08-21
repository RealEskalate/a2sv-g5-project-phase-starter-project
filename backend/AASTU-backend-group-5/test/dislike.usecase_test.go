package test

import (
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DislikeUseCaseTestSuite struct {
	suite.Suite
	mockDislikeRepo *mocks.DisLike_Repository_interface
	dislikeUseCase  *usecase.DislikeUseCase
}

func (suite *DislikeUseCaseTestSuite) SetupTest() {
	suite.mockDislikeRepo = mocks.NewDisLike_Repository_interface(suite.T())
	suite.dislikeUseCase = &usecase.DislikeUseCase{
		DislikeRepo: suite.mockDislikeRepo,
	}
}

func (suite *DislikeUseCaseTestSuite) TestGetDisLikes() {
	// Using hex strings for postID and userID
	postID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	// Converting to ObjectIDs for the mock data
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	suite.Require().NoError(err)

	postObjectID, err := primitive.ObjectIDFromHex(postID)
	suite.Require().NoError(err)

	// Expected dislikes based on the mock data
	expectedDislikes := []domain.DisLike{
		{UserID: userObjectID, PostID: postObjectID},
	}

	// Setting up the mock for GetDisLikes
	suite.mockDislikeRepo.On("GetDisLikes", postID).Return(expectedDislikes, nil)

	// Calling the use case method
	dislikes, err := suite.dislikeUseCase.GetDisLikes(postID)

	// Assertions
	suite.NoError(err)
	suite.Len(dislikes, len(expectedDislikes))
	suite.Equal(expectedDislikes[0].UserID, dislikes[0].UserID)
	suite.Equal(expectedDislikes[0].PostID, dislikes[0].PostID)

	// Verify mock expectations
	suite.mockDislikeRepo.AssertExpectations(suite.T())
}

func (suite *DislikeUseCaseTestSuite) TestCreateDisLike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	suite.mockDislikeRepo.On("CreateDisLike", userID, postID).Return(nil)

	err := suite.dislikeUseCase.CreateDisLike(userID, postID)
	suite.NoError(err)

	suite.mockDislikeRepo.AssertExpectations(suite.T())
}

func (suite *DislikeUseCaseTestSuite) TestDeleteDisLike() {
	dislikeID := primitive.NewObjectID().Hex()

	suite.mockDislikeRepo.On("DeleteDisLike", dislikeID).Return(nil)

	err := suite.dislikeUseCase.DeleteDisLike(dislikeID)
	suite.NoError(err)

	suite.mockDislikeRepo.AssertExpectations(suite.T())
}

func TestDislikeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DislikeUseCaseTestSuite))
}
