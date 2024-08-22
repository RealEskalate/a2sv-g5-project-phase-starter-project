package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	repository "github.com/aait.backend.g5.main/backend/Repository"
)

type BlogPopularityActionRepositorySuite struct {
	suite.Suite

	blogUserActionCollection *mocks.Collection
	blogActionCollection     *mocks.Collection
	repository               *repository.BlogPupularityActionRepo
}

func (suite *BlogPopularityActionRepositorySuite) SetupTest() {
	suite.blogUserActionCollection = new(mocks.Collection)
	suite.blogActionCollection = new(mocks.Collection)
	suite.repository = &repository.BlogPupularityActionRepo{
		BlogUserActionCollection: suite.blogUserActionCollection,
		BlogActionCollection:     suite.blogActionCollection,
	}
}

// Test Like Success
func (suite *BlogPopularityActionRepositorySuite) TestLike_Success() {
	popularityAction := dtos.TrackPopularityRequest{
		BlogID: "blog1",
		Action: "like",
		UserID: "user1",
	}

	// Mocking the first UpdateOne call for the BlogUserActionCollection
	suite.blogUserActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
		"user_id": popularityAction.UserID,
	}, bson.M{
		"$set": bson.M{"action": popularityAction.Action},
	}, options.Update().SetUpsert(true)).Return(nil, nil)

	// Mocking the second UpdateOne call for the BlogActionCollection
	suite.blogActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
	}, bson.M{
		"$inc": bson.M{"like_count": 1},
	}).Return(nil, nil)

	err := suite.repository.Like(context.Background(), popularityAction)

	suite.Nil(err)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	suite.blogActionCollection.AssertExpectations(suite.T())
}

// Test Like Failure
func (suite *BlogPopularityActionRepositorySuite) TestLike_Failure() {
	popularityAction := dtos.TrackPopularityRequest{
		BlogID: "blog1",
		Action: "like",
		UserID: "user1",
	}

	// Mocking the first UpdateOne call to return an error
	suite.blogUserActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
		"user_id": popularityAction.UserID,
	}, bson.M{
		"$set": bson.M{"action": popularityAction.Action},
	}, options.Update().SetUpsert(true)).Return(nil, errors.New("update error"))

	err := suite.repository.Like(context.Background(), popularityAction)

	suite.NotNil(err)
	suite.Equal("update error", err.Message)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	suite.blogActionCollection.AssertNotCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything, mock.Anything)
}

// Test Dislike Success
func (suite *BlogPopularityActionRepositorySuite) TestDislike_Success() {
	popularityAction := dtos.TrackPopularityRequest{
		BlogID: "blog1",
		Action: "dislike",
		UserID: "user1",
	}

	// Mocking the first UpdateOne call for the BlogUserActionCollection
	suite.blogUserActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
		"user_id": popularityAction.UserID,
	}, bson.M{
		"$set": bson.M{"action": popularityAction.Action},
	}, options.Update().SetUpsert(true)).Return(nil, nil)

	// Mocking the second UpdateOne call for the BlogActionCollection
	suite.blogActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
	}, bson.M{
		"$inc": bson.M{"dislike_count": 1},
	}).Return(nil, nil)

	err := suite.repository.Dislike(context.Background(), popularityAction)

	suite.Nil(err)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	suite.blogActionCollection.AssertExpectations(suite.T())
}

// Test Dislike Failure
func (suite *BlogPopularityActionRepositorySuite) TestDislike_Failure() {
	popularityAction := dtos.TrackPopularityRequest{
		BlogID: "blog1",
		Action: "dislike",
		UserID: "user1",
	}

	// Mocking the first UpdateOne call to return an error
	suite.blogUserActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
		"user_id": popularityAction.UserID,
	}, bson.M{
		"$set": bson.M{"action": popularityAction.Action},
	}, options.Update().SetUpsert(true)).Return(nil, errors.New("update error"))

	err := suite.repository.Dislike(context.Background(), popularityAction)

	suite.NotNil(err)
	suite.Equal("update error", err.Message)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	suite.blogActionCollection.AssertNotCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything, mock.Anything)
}

// Test GetBlogPopularityAction Success
func (suite *BlogPopularityActionRepositorySuite) TestGetBlogPopularityAction_Success() {
	blogID := "blog1"
	userID := "user1"
	expectedResult := models.PopularityAction{
		BlogID: blogID,
		Action: "like",
		UserID: userID,
	}

	// Mocking FindOne to return a successful SingleResult with the expected PopularityAction
	singleResult := &mocks.SingleResult{}
	suite.blogUserActionCollection.On("FindOne", mock.Anything, bson.M{
		"blog_id": blogID,
		"user_id": userID,
	}).Return(singleResult)
	singleResult.On("Decode", mock.AnythingOfType("*models.PopularityAction")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.PopularityAction)
		*arg = expectedResult
	}).Return(nil)

	result, err := suite.repository.GetBlogPopularityAction(context.Background(), blogID, userID)

	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(expectedResult, *result)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

// Test GetBlogPopularityAction Failure
func (suite *BlogPopularityActionRepositorySuite) TestGetBlogPopularityAction_Failure() {
	blogID := "blog1"
	userID := "user1"

	// Mocking FindOne to return a SingleResult that triggers an error
	singleResult := &mocks.SingleResult{}
	suite.blogUserActionCollection.On("FindOne", mock.Anything, bson.M{
		"blog_id": blogID,
		"user_id": userID,
	}).Return(singleResult)
	singleResult.On("Decode", mock.AnythingOfType("*models.PopularityAction")).Return(errors.New("decoding error"))

	result, err := suite.repository.GetBlogPopularityAction(context.Background(), blogID, userID)

	suite.NotNil(err)
	suite.Nil(result)
	suite.Equal("decoding error", err.Message)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	singleResult.AssertExpectations(suite.T())
}

// Test UndoLike Success
func (suite *BlogPopularityActionRepositorySuite) TestUndoLike_Success() {
	popularityAction := dtos.TrackPopularityRequest{
		BlogID: "blog1",
		UserID: "user1",
	}

	// Mocking DeleteOne for BlogUserActionCollection
	suite.blogUserActionCollection.On("DeleteOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
		"user_id": popularityAction.UserID,
	}).Return(int64(1), nil)

	// Mocking UpdateOne for BlogActionCollection
	suite.blogActionCollection.On("UpdateOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
	}, bson.M{
		"$inc": bson.M{"like_count": -1},
	}).Return(nil, nil)

	err := suite.repository.UndoLike(context.Background(), popularityAction)

	suite.Nil(err)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	suite.blogActionCollection.AssertExpectations(suite.T())
}

// Test UndoLike Failure
func (suite *BlogPopularityActionRepositorySuite) TestUndoLike_Failure() {
	popularityAction := dtos.TrackPopularityRequest{
		BlogID: "blog1",
		UserID: "user1",
	}

	// Mocking DeleteOne to return an error
	suite.blogUserActionCollection.On("DeleteOne", mock.Anything, bson.M{
		"blog_id": popularityAction.BlogID,
		"user_id": popularityAction.UserID,
	}).Return(int64(0), errors.New("delete error"))

	err := suite.repository.UndoLike(context.Background(), popularityAction)

	suite.NotNil(err)
	suite.Equal("delete error", err.Message)
	suite.blogUserActionCollection.AssertExpectations(suite.T())
	suite.blogActionCollection.AssertNotCalled(suite.T(), "UpdateOne", mock.Anything, mock.Anything, mock.Anything)
}

func TestBlogPopularityActionRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlogPopularityActionRepositorySuite))
}
