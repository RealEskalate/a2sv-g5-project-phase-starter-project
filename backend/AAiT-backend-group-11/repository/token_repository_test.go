package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateRefreshToken(t *testing.T) {
	// Test case: Successfully create a new refresh token
	t.Run("Success", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		mockToken := &entities.RefreshToken{
			ID:     primitive.NewObjectID(),
			UserID: "user123",
			Token:  "some-refresh-token",
		}

		// Define the filter to check for existing tokens
		filter := bson.D{{"userId", mockToken.UserID}}

		mockSingleResult := new(mocks.SingleResult)
		// Mock the FindOne method to return ErrNoDocuments, simulating no existing token
		mockCollection.On("FindOne", mock.Anything, filter).Return(mockSingleResult, mongo.ErrNoDocuments)
		
		mockSingleResult.On("Err").Return(errors.New("mongo: no documents in result"))
		
		// Mock the InsertOne method to simulate inserting the token
		mockCollection.On("InsertOne", mock.Anything, mockToken).Return(nil, nil)

		// Run the function
		result, err := repo.CreateRefreshToken(mockToken)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockToken, result)

		// Verify the expectations
		mockCollection.AssertExpectations(t)
	})

	// Test case: Refresh token already exists
	t.Run("AlreadyExists", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		mockToken := &entities.RefreshToken{
			ID:     primitive.NewObjectID(),
			UserID: "user123",
			Token:  "some-refresh-token",
		}

		// Define the filter to check for existing tokens
		filter := bson.D{{"userId", mockToken.UserID}}

		mockSingleResult := new(mocks.SingleResult)
		// Mock the FindOne method to return a valid result, simulating an existing token
		mockCollection.On("FindOne", mock.Anything, filter).Return(mockSingleResult, nil)

		mockSingleResult.On("Err").Return(nil)
		// Run the function
		result, err := repo.CreateRefreshToken(mockToken)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "refresh token already exists", err.Error())

		// Verify the expectations
		mockCollection.AssertExpectations(t)
	})
}

func TestFindRefreshTokenByUserId(t *testing.T) {
	// Test case: Successfully find a refresh token by user ID
	t.Run("Success", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Create a mock token
		mockToken := &entities.RefreshToken{
			ID:     primitive.NewObjectID(),
			UserID: primitive.NewObjectID().Hex(),
			Token:  "some-refresh-token",
		}

		// Convert user ID to ObjectID
		userID, _ := primitive.ObjectIDFromHex(mockToken.UserID)

		// Define the filter
		filter := bson.D{{"userId", userID}}

		// Mock the FindOne method
		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Err").Return(nil)
		mockSingleResult.On("Decode", mock.Anything).Return(func(v interface{}) error {
			*v.(*entities.RefreshToken) = *mockToken
			return nil
		})
		mockCollection.On("FindOne", mock.Anything, filter).Return(mockSingleResult)

		// Run the function
		result, err := repo.FindRefreshTokenByUserId(mockToken.UserID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockToken, result)

		// Verify the expectations
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	// Test case: Invalid user ID format
	t.Run("InvalidUserID", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Run the function with an invalid user ID
		result, err := repo.FindRefreshTokenByUserId("invalidUserID")

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "the provided hex string is not a valid ObjectID", err.Error())

		// No need to verify expectations as the function should return before interacting with the collection
	})

	// Test case: No refresh token found (ErrNoDocuments)
	t.Run("NoDocumentFound", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Create a mock token
		mockToken := &entities.RefreshToken{
			ID:     primitive.NewObjectID(),
			UserID: primitive.NewObjectID().Hex(),
			Token:  "some-refresh-token",
		}

		// Convert user ID to ObjectID
		userID, _ := primitive.ObjectIDFromHex(mockToken.UserID)

		// Define the filter
		filter := bson.D{{"userId", userID}}

		// Mock the FindOne method to simulate no document found
		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Err").Return(errors.New("mongo: no documents in result"))
		mockCollection.On("FindOne", mock.Anything, filter).Return(mockSingleResult)

		// Run the function
		result, err := repo.FindRefreshTokenByUserId(mockToken.UserID)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "mongo: no documents in result", err.Error())

		// Verify the expectations
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	// Test case: Decode error
	t.Run("DecodeError", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Create a mock token
		mockToken := &entities.RefreshToken{
			ID:     primitive.NewObjectID(),
			UserID: primitive.NewObjectID().Hex(),
			Token:  "some-refresh-token",
		}

		// Convert user ID to ObjectID
		userID, _ := primitive.ObjectIDFromHex(mockToken.UserID)

		// Define the filter
		filter := bson.D{{"userId", userID}}

		// Mock the FindOne method and simulate a Decode error
		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Err").Return(nil)
		mockSingleResult.On("Decode", mock.Anything).Return(errors.New("decode error"))
		mockCollection.On("FindOne", mock.Anything, filter).Return(mockSingleResult)

		// Run the function
		result, err := repo.FindRefreshTokenByUserId(mockToken.UserID)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "decode error", err.Error())

		// Verify the expectations
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDeleteRefreshTokenByUserId(t *testing.T) {
	// Test case: Successfully delete a refresh token by user ID
	t.Run("Success", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)

		// Convert user ID to ObjectID
		userID, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")

		// Define the filter
		filter := bson.D{{"userId", userID}}

		// Mock the FindOneAndDelete method
		mockSingleResult := new(mocks.SingleResult)
		mockCollection.On("FindOneAndDelete", mock.Anything, filter).Return(mockSingleResult)
		
		mockSingleResult.On("Err").Return(nil)

		// Run the function
		err := repo.DeleteRefreshTokenByUserId("507f1f77bcf86cd799439011")

		// Assertions
		assert.NoError(t, err)

		// Verify the expectations
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	// Test case: Invalid user ID format
	t.Run("InvalidUserID", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Run the function with an invalid user ID
		err := repo.DeleteRefreshTokenByUserId("invalidUserID")

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "the provided hex string is not a valid ObjectID", err.Error())

		// No need to verify expectations as the function should return before interacting with the collection
	})

	// Test case: No refresh token found (ErrNoDocuments)
	t.Run("NoDocumentFound", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Convert user ID to ObjectID
		userID, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")

		// Define the filter
		filter := bson.D{{"userId", userID}}

		// Mock the FindOneAndDelete method to simulate no document found
		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Err").Return(errors.New("mongo: no documents in result"))
		mockCollection.On("FindOneAndDelete", mock.Anything, filter).Return(mockSingleResult)

		// Run the function
		err := repo.DeleteRefreshTokenByUserId("507f1f77bcf86cd799439011")

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "mongo: no documents in result", err.Error())

		// Verify the expectations
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	// Test case: Error during deletion
	t.Run("DeleteError", func(t *testing.T) {
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		repo := repository.NewTokenRepository(mockDatabase, mockCollection)
	
		// Convert user ID to ObjectID
		userID, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")

		// Define the filter
		filter := bson.D{{"userId", userID}}

		// Mock the FindOneAndDelete method to simulate an error
		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Err").Return(errors.New("deletion error"))
		mockCollection.On("FindOneAndDelete", mock.Anything, filter).Return(mockSingleResult)

		// Run the function
		err := repo.DeleteRefreshTokenByUserId("507f1f77bcf86cd799439011")

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "deletion error", err.Error())

		// Verify the expectations
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}