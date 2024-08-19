package test

import (
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeUseCaseTestSuite struct {
	suite.Suite
	mockLikeRepo *mocks.Like_Repositroy_interface
	likeUseCase  *usecase.LikeUseCase
}

func (suite *LikeUseCaseTestSuite) SetupTest() {
	suite.mockLikeRepo = mocks.NewLike_Repositroy_interface(suite.T())
	suite.likeUseCase = &usecase.LikeUseCase{
		LikeRepo: suite.mockLikeRepo,
	}
}

func (suite *LikeUseCaseTestSuite) TestGetLikes() {
	postID := primitive.NewObjectID().Hex()

	// Create expected likes with properly formatted object IDs
	expectedLikes := []domain.Like{
		{
			UserID: primitive.NewObjectID(),
			PostID: primitive.NewObjectID(),
		},
	}

	// Mock the GetLikes repository method
	suite.mockLikeRepo.On("GetLikes", postID).Return(expectedLikes, nil)

	// Call the use case method
	likes, err := suite.likeUseCase.GetLikes(postID)

	// Assertions
	suite.NoError(err)
	suite.Len(likes, len(expectedLikes))
	suite.Equal(expectedLikes[0].UserID, likes[0].UserID)
	suite.Equal(expectedLikes[0].PostID, likes[0].PostID)

	// Verify that the expectations on the mock repository were met
	suite.mockLikeRepo.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseTestSuite) TestCreateLike() {
	userID := primitive.NewObjectID().Hex()
	postID := primitive.NewObjectID().Hex()

	// Mock the CreateLike repository method
	suite.mockLikeRepo.On("CreateLike", userID, postID).Return(nil)

	// Call the use case method
	err := suite.likeUseCase.CreateLike(userID, postID)

	// Assertions
	suite.NoError(err)

	// Verify that the expectations on the mock repository were met
	suite.mockLikeRepo.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseTestSuite) TestDeleteLike() {
	likeID := primitive.NewObjectID().Hex()

	// Mock the DeleteLike repository method
	suite.mockLikeRepo.On("DeleteLike", likeID).Return(nil)

	// Call the use case method
	err := suite.likeUseCase.DeleteLike(likeID)

	// Assertions
	suite.NoError(err)

	// Verify that the expectations on the mock repository were met
	suite.mockLikeRepo.AssertExpectations(suite.T())
}

func TestLikeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(LikeUseCaseTestSuite))
}
