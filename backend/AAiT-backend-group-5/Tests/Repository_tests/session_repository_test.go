package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	repository "github.com/aait.backend.g5.main/backend/Repository"
)

type SessionRepositorySuite struct {
	suite.Suite

	db         *mocks.Database
	collection *mocks.Collection
	repository *repository.SessionRepo
}

func (suite *SessionRepositorySuite) SetupTest() {
	suite.db = new(mocks.Database)
	suite.collection = new(mocks.Collection)
	suite.repository = &repository.SessionRepo{
		Collection: suite.collection,
	}
}

// Test SaveToken function
func (suite *SessionRepositorySuite) TestSaveToken_Success_UpdateExistingToken() {
	session := &models.Session{
		UserID:       "user123",
		RefreshToken: "new-refresh-token",
		AccessToken:  "new-access-token",
	}

	// Mock the FindOne operation to simulate an existing session
	mockResult := new(mocks.SingleResult)
	suite.collection.On("FindOne", mock.Anything, bson.M{"user_id": session.UserID}).Return(mockResult)

	// Mock the Decode operation to simulate successful decoding
	mockResult.On("Decode", mock.AnythingOfType("*models.Session")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Session)
		arg.UserID = session.UserID
		arg.RefreshToken = "old-refresh-token"
		arg.AccessToken = "old-access-token"
	}).Return(nil)

	// Mock the UpdateOne operation
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"user_id": session.UserID}, mock.Anything).Return(&mongo.UpdateResult{}, nil)

	// Call the SaveToken method
	errResp := suite.repository.SaveToken(context.TODO(), session)

	// Assertions
	suite.Nil(errResp)
	suite.collection.AssertExpectations(suite.T())
}

func (suite *SessionRepositorySuite) TestSaveToken_Success_InsertNewToken() {
	session := &models.Session{
		UserID:       "user123",
		RefreshToken: "new-refresh-token",
		AccessToken:  "new-access-token",
	}

	// Mock the FindOne operation to simulate no existing session
	mockResult := new(mocks.SingleResult)
	suite.collection.On("FindOne", mock.Anything, bson.M{"user_id": session.UserID}).Return(mockResult)

	// Mock the Decode operation to simulate a "no documents found" error
	mockResult.On("Decode", mock.AnythingOfType("*models.Session")).Return(mongo.ErrNoDocuments)

	// Mock the InsertOne operation
	suite.collection.On("InsertOne", mock.Anything, session).Return(&mongo.InsertOneResult{}, nil)

	// Call the SaveToken method
	errResp := suite.repository.SaveToken(context.TODO(), session)

	// Assertions
	suite.Nil(errResp)
	suite.collection.AssertExpectations(suite.T())
}

// Test UpdateToken function
func (suite *SessionRepositorySuite) TestUpdateToken_Success() {
	session := &models.Session{
		UserID:       "user123",
		RefreshToken: "updated-refresh-token",
		AccessToken:  "updated-access-token",
	}

	// Mock the UpdateOne operation
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"user_id": session.UserID}, mock.Anything).Return(&mongo.UpdateResult{}, nil)

	// Call the UpdateToken method
	errResp := suite.repository.UpdateToken(context.TODO(), session)

	// Assertions
	suite.Nil(errResp)
	suite.collection.AssertExpectations(suite.T())
}

func (suite *SessionRepositorySuite) TestUpdateToken_Failure() {
	session := &models.Session{
		UserID:       "user123",
		RefreshToken: "updated-refresh-token",
		AccessToken:  "updated-access-token",
	}

	// Mock the UpdateOne operation to return an error
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"user_id": session.UserID}, mock.Anything).Return(nil, errors.New("update error"))

	// Call the UpdateToken method
	errResp := suite.repository.UpdateToken(context.TODO(), session)

	// Assertions
	suite.NotNil(errResp)
	suite.Equal("update error", errResp.Message)
	suite.collection.AssertExpectations(suite.T())
}

// Test RemoveToken function
func (suite *SessionRepositorySuite) TestRemoveToken_Success() {
	userID := "user123"

	// Mock the DeleteOne operation
	suite.collection.On("DeleteOne", mock.Anything, bson.M{"user_id": userID}).Return(int64(1), nil)

	// Call the RemoveToken method
	errResp := suite.repository.RemoveToken(context.TODO(), userID)

	// Assertions
	suite.Nil(errResp)
	suite.collection.AssertExpectations(suite.T())
}

func (suite *SessionRepositorySuite) TestRemoveToken_Failure() {
	userID := "user123"

	// Mock the DeleteOne operation to return an error
	suite.collection.On("DeleteOne", mock.Anything, bson.M{"user_id": userID}).Return(int64(0), errors.New("delete error"))

	// Call the RemoveToken method
	errResp := suite.repository.RemoveToken(context.TODO(), userID)

	// Assertions
	suite.NotNil(errResp)
	suite.Equal("delete error", errResp.Message)
	suite.collection.AssertExpectations(suite.T())
}

// Test GetToken function
func (suite *SessionRepositorySuite) TestGetToken_Success() {
	userID := "user123"
	expectedSession := &models.Session{
		UserID:       userID,
		RefreshToken: "refresh-token",
		AccessToken:  "access-token",
	}

	// Mock the FindOne operation
	mockResult := new(mocks.SingleResult)
	suite.collection.On("FindOne", mock.Anything, bson.M{"user_id": userID}).Return(mockResult)

	// Mock the Decode operation to simulate successful decoding
	mockResult.On("Decode", mock.AnythingOfType("*models.Session")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Session)
		arg.UserID = expectedSession.UserID
		arg.RefreshToken = expectedSession.RefreshToken
		arg.AccessToken = expectedSession.AccessToken
	}).Return(nil)

	// Call the GetToken method
	result, errResp := suite.repository.GetToken(context.TODO(), userID)

	// Assertions
	suite.Nil(errResp)
	suite.Equal(expectedSession, result)
	suite.collection.AssertExpectations(suite.T())
}

func (suite *SessionRepositorySuite) TestGetToken_Failure() {
	userID := "user123"

	// Mock the FindOne operation
	mockResult := new(mocks.SingleResult)
	suite.collection.On("FindOne", mock.Anything, bson.M{"user_id": userID}).Return(mockResult)

	// Mock the Decode operation to return an error
	mockResult.On("Decode", mock.AnythingOfType("*models.Session")).Return(mongo.ErrNoDocuments)

	// Call the GetToken method
	result, errResp := suite.repository.GetToken(context.TODO(), userID)

	// Assertions
	suite.Nil(result)
	suite.NotNil(errResp)
	suite.Equal("Session not found", errResp.Message)
	suite.collection.AssertExpectations(suite.T())
}

func TestSessionRepositorySuite(t *testing.T) {
	suite.Run(t, new(SessionRepositorySuite))
}
