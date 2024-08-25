package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/domain/mocks" // Adjust this import path based on your project structure
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeUseCaseSuite struct {
	suite.Suite
	likeUseCase  domain.LikeUseCase
	likeRepoMock *mocks.LikeRepository
	blogRepoMock *mocks.BlogRepository
}

func (suite *LikeUseCaseSuite) SetupTest() {
	// Initialize the mocks
	suite.likeRepoMock = &mocks.LikeRepository{}
	suite.blogRepoMock = &mocks.BlogRepository{}

	// Initialize the use case with the mocked repositories
	suite.likeUseCase = NewLikeUseCase(suite.likeRepoMock, suite.blogRepoMock, 2*time.Second)
}

func (suite *LikeUseCaseSuite) TestGetByID_Success() {
	// Arrange
	expectedLike := &domain.Like{
		LikeID:    primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		CreatedAt: time.Now(),
	}

	suite.likeRepoMock.On("GetByID", mock.Anything, "user123", "blog123").Return(expectedLike, nil)

	// Act
	like, err := suite.likeUseCase.GetByID(context.Background(), "user123", "blog123")

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLike, like)
	suite.likeRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestGetByID_Failure() {
	// Arrange
	suite.likeRepoMock.On("GetByID", mock.Anything, "user123", "blog123").Return(nil, errors.New("not found"))

	// Act
	like, err := suite.likeUseCase.GetByID(context.Background(), "user123", "blog123")

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), like)
	suite.likeRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestLikeBlog_Success() {
	// Arrange
	like := &domain.Like{
		LikeID:    primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		CreatedAt: time.Now(),
	}

	suite.likeRepoMock.On("LikeBlog", mock.Anything, like).Return(like, nil)
	suite.blogRepoMock.On("UpdateLikeCount", mock.Anything, "blog123", true).Return(nil)

	// Act
	liked, err := suite.likeUseCase.LikeBlog(context.Background(), like)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), like, liked)
	suite.likeRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestLikeBlog_Failure_LikeRepo() {
	// Arrange
	like := &domain.Like{
		LikeID:    primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		CreatedAt: time.Now(),
	}

	suite.likeRepoMock.On("LikeBlog", mock.Anything, like).Return(nil, errors.New("like repo error"))

	// Act
	liked, err := suite.likeUseCase.LikeBlog(context.Background(), like)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), liked)
	suite.likeRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestLikeBlog_Failure_BlogRepo() {
	// Arrange
	like := &domain.Like{
		LikeID:    primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		CreatedAt: time.Now(),
	}

	suite.likeRepoMock.On("LikeBlog", mock.Anything, like).Return(like, nil)
	suite.blogRepoMock.On("UpdateLikeCount", mock.Anything, "blog123", true).Return(errors.New("blog repo error"))

	// Act
	liked, err := suite.likeUseCase.LikeBlog(context.Background(), like)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), liked)
	suite.likeRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestUnlikeBlog_Success() {
	// Arrange
	likeID := "like123"
	like := &domain.Like{
		LikeID:    primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		CreatedAt: time.Now(),
	}

	suite.likeRepoMock.On("UnlikeBlog", mock.Anything, likeID).Return(like, nil)
	suite.blogRepoMock.On("UpdateLikeCount", mock.Anything, like.BlogID, false).Return(nil)

	// Act
	unliked, err := suite.likeUseCase.UnlikeBlog(context.Background(), likeID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), like, unliked)
	suite.likeRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestUnlikeBlog_Failure_LikeRepo() {
	// Arrange
	likeID := "like123"

	suite.likeRepoMock.On("UnlikeBlog", mock.Anything, likeID).Return(nil, errors.New("like repo error"))

	// Act
	unliked, err := suite.likeUseCase.UnlikeBlog(context.Background(), likeID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), unliked)
	suite.likeRepoMock.AssertExpectations(suite.T())
}

func (suite *LikeUseCaseSuite) TestUnlikeBlog_Failure_BlogRepo() {
	// Arrange
	likeID := "like123"
	like := &domain.Like{
		LikeID:    primitive.NewObjectID(),
		UserID:    "user123",
		BlogID:    "blog123",
		CreatedAt: time.Now(),
	}

	suite.likeRepoMock.On("UnlikeBlog", mock.Anything, likeID).Return(like, nil)
	suite.blogRepoMock.On("UpdateLikeCount", mock.Anything, like.BlogID, false).Return(errors.New("blog repo error"))

	// Act
	unliked, err := suite.likeUseCase.UnlikeBlog(context.Background(), likeID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), unliked)
	suite.likeRepoMock.AssertExpectations(suite.T())
	suite.blogRepoMock.AssertExpectations(suite.T())
}

// Run the test suite
func TestLikeUseCaseSuite(t *testing.T) {
	suite.Run(t, new(LikeUseCaseSuite))
}
