package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/domain/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRatingUseCaseSuite struct {
	suite.Suite
	blogRatingUseCase domain.BlogRatingUseCase
	ratingRepoMock    *mocks.BlogRatingRepository
	blogRepoMock      *mocks.BlogRepository
}

func (suite *BlogRatingUseCaseSuite) SetupTest() {
	// Initialize the mocks
	suite.ratingRepoMock = &mocks.BlogRatingRepository{}
	suite.blogRepoMock = &mocks.BlogRepository{}

	// Initialize the use case with the mocked repositories
	suite.blogRatingUseCase = NewBlogRatingUseCase(suite.ratingRepoMock, suite.blogRepoMock, 2*time.Second)
}

func (suite *BlogRatingUseCaseSuite) TestInsertRating_Success() {
	// Arrange
	ratingRequest := &domain.BlogRatingRequest{
		RatingID: "",
		UserID:   "user123",
		BlogID:   "blog123",
		Rating:   5,
	}

	rating := &domain.BlogRating{
		UserID:    "user123",
		BlogID:    "blog123",
		Rating:    5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.ratingRepoMock.On("InsertRating", mock.Anything, mock.AnythingOfType("*domain.BlogRating")).Return(rating, nil)
	suite.blogRepoMock.On("InsertRating", mock.Anything, mock.AnythingOfType("*domain.BlogRating")).Return(nil)

	// Act
	insertedRating, err := suite.blogRatingUseCase.InsertRating(context.Background(), ratingRequest)

	// Assert
	suite.NoError(err)
	suite.Equal(rating, insertedRating)
	suite.ratingRepoMock.AssertCalled(suite.T(), "InsertRating", mock.Anything, mock.AnythingOfType("*domain.BlogRating"))
	suite.blogRepoMock.AssertCalled(suite.T(), "InsertRating", mock.Anything, mock.AnythingOfType("*domain.BlogRating"))
}

func (suite *BlogRatingUseCaseSuite) TestInsertRating_Failure() {
	// Arrange
	ratingRequest := &domain.BlogRatingRequest{
		RatingID: "",
		UserID:   "user123",
		BlogID:   "blog123",
		Rating:   5,
	}

	suite.ratingRepoMock.On("InsertRating", mock.Anything, mock.AnythingOfType("*domain.BlogRating")).Return(nil, assert.AnError)

	// Act
	insertedRating, err := suite.blogRatingUseCase.InsertRating(context.Background(), ratingRequest)

	// Assert
	suite.Error(err)
	suite.Nil(insertedRating)
	suite.ratingRepoMock.AssertCalled(suite.T(), "InsertRating", mock.Anything, mock.AnythingOfType("*domain.BlogRating"))
}

func (suite *BlogRatingUseCaseSuite) TestGetRatingByBlogID_Success() {
	// Arrange
	expectedRatings := []*domain.BlogRating{
		{RatingID: primitive.NewObjectID(), BlogID: "blog123", UserID: "user123", Rating: 4},
	}
	suite.ratingRepoMock.On("GetRatingByBlogID", mock.Anything, "blog123").Return(expectedRatings, nil)

	// Act
	ratings, err := suite.blogRatingUseCase.GetRatingByBlogID(context.Background(), "blog123")

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRatings, ratings)
	suite.ratingRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestGetRatingByBlogID_Failure() {
	// Arrange
	suite.ratingRepoMock.On("GetRatingByBlogID", mock.Anything, "blog123").Return(nil, errors.New("fetch error"))

	// Act
	ratings, err := suite.blogRatingUseCase.GetRatingByBlogID(context.Background(), "blog123")

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), ratings)
	suite.ratingRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestGetRatingByID_Success() {
	// Arrange
	expectedRating := &domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "blog123",
		UserID:   "user123",
		Rating:   4,
	}
	suite.ratingRepoMock.On("GetRatingByID", mock.Anything, "rating123").Return(expectedRating, nil)

	// Act
	rating, err := suite.blogRatingUseCase.GetRatingByID(context.Background(), "rating123")

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRating, rating)
	suite.ratingRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestGetRatingByID_Failure() {
	// Arrange
	suite.ratingRepoMock.On("GetRatingByID", mock.Anything, "rating123").Return(nil, errors.New("fetch error"))

	// Act
	rating, err := suite.blogRatingUseCase.GetRatingByID(context.Background(), "rating123")

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), rating)
	suite.ratingRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestUpdateRating_Success() {
	// Arrange
	ratingID := "rating123"
	newRating := 5
	expectedRating := &domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "blog123",
		UserID:   "user123",
		Rating:   newRating,
	}
	prevRating := 4

	suite.ratingRepoMock.On("UpdateRating", mock.Anything, newRating, ratingID).Return(expectedRating, prevRating, nil)
	suite.blogRepoMock.On("UpdateRating", mock.Anything, expectedRating, prevRating).Return(nil)

	// Act
	updatedRating, err := suite.blogRatingUseCase.UpdateRating(context.Background(), newRating, ratingID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRating, updatedRating)
	suite.ratingRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestUpdateRating_Failure() {
	// Arrange
	ratingID := "rating123"
	newRating := 5

	suite.ratingRepoMock.On("UpdateRating", mock.Anything, newRating, ratingID).Return(nil, 0, errors.New("update error"))

	// Act
	updatedRating, err := suite.blogRatingUseCase.UpdateRating(context.Background(), newRating, ratingID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedRating)
	suite.ratingRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestDeleteRating_Success() {
	// Arrange
	ratingID := "rating123"
	expectedRating := &domain.BlogRating{
		RatingID: primitive.NewObjectID(),
		BlogID:   "blog123",
		UserID:   "user123",
		Rating:   4,
	}

	suite.ratingRepoMock.On("DeleteRating", mock.Anything, ratingID).Return(expectedRating, nil)
	suite.blogRepoMock.On("DeleteRating", mock.Anything, expectedRating).Return(nil)

	// Act
	deletedRating, err := suite.blogRatingUseCase.DeleteRating(context.Background(), ratingID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRating, deletedRating)
	suite.ratingRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *BlogRatingUseCaseSuite) TestDeleteRating_Failure() {
	// Arrange
	ratingID := "rating123"

	suite.ratingRepoMock.On("DeleteRating", mock.Anything, ratingID).Return(nil, errors.New("delete error"))

	// Act
	deletedRating, err := suite.blogRatingUseCase.DeleteRating(context.Background(), ratingID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), deletedRating)
	suite.ratingRepoMock.AssertExpectations(suite.T())
}

func TestBlogRatingUseCaseSuite(t *testing.T) {
	suite.Run(t, new(BlogRatingUseCaseSuite))
}
