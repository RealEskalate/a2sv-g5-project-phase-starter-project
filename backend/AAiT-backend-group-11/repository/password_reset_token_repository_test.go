package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreatePasswordResetToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)

		// Instantiate the repository with the mock collection
		repo := repository.NewPasswordTokenRepository(mockDatabase, mockCollection)

		// Create a sample PasswordResetToken entity
		token := &entities.PasswordResetToken{
			Token: "sample_token",
			// Add other necessary fields here
		}

		mockSingleResult := new(mocks.SingleResult)
		// Mock the InsertOne method to return a success response
		mockCollection.On("InsertOne", mock.Anything, token).Return(mockSingleResult, nil)

		// Call the CreatePasswordResetToken method
		result, err := repo.CreatePasswordResetToken(token)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, token, result)

		// Verify that the mock methods were called
		mockCollection.AssertExpectations(t)
	})

	t.Run("Insert_Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		// Instantiate the repository with the mock collection
		repo := repository.NewPasswordTokenRepository(mockDatabase, mockCollection)

		// Create a sample PasswordResetToken entity
		token := &entities.PasswordResetToken{
			Token: "sample_token",
			// Add other necessary fields here
		}

		// Mock the InsertOne method to return an error
		mockCollection.On("InsertOne", mock.Anything, token).Return(nil, assert.AnError)

		// Call the CreatePasswordResetToken method
		result, err := repo.CreatePasswordResetToken(token)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, assert.AnError, err)

		// Verify that the mock methods were called
		mockCollection.AssertExpectations(t)
	})
}

func TestDeletePasswordResetTokenByUserId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		// Instantiate the repository with the mock collection
		repo := repository.NewPasswordTokenRepository(mockDatabase, mockCollection)

		mockDeleteCount := int64(1)
		// Mock the DeleteOne method to return a success response
		mockCollection.On("DeleteOne", mock.Anything, mock.Anything).Return(mockDeleteCount, nil)

		// Call the DeletePasswordResetTokenByUserId method
		err := repo.DeletePasswordResetTokenByUserId("sample_user_id")

		// Assertions
		assert.NoError(t, err)

		// Verify that the mock methods were called
		mockCollection.AssertExpectations(t)
	})

	t.Run("Delete_Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		// Instantiate the repository with the mock collection
		repo := repository.NewPasswordTokenRepository(mockDatabase, mockCollection)

		mockDeleteCount := int64(0)
		// Mock the DeleteOne method to return an error
		mockCollection.On("DeleteOne", mock.Anything, mock.Anything).Return(mockDeleteCount, assert.AnError)

		// Call the DeletePasswordResetTokenByUserId method
		err := repo.DeletePasswordResetTokenByUserId("sample_user_id")

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)

		// Verify that the mock methods were called
		mockCollection.AssertExpectations(t)
	})
}

func TestFindPasswordReset(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)
		// Instantiate the repository with the mock collection
		repo := repository.NewPasswordTokenRepository(mockDatabase, mockCollection)

		// Create a sample PasswordResetToken entity
		expectedToken := &entities.PasswordResetToken{
			ID:   primitive.NewObjectID(),
			UserID: primitive.NewObjectID(),
			Token: "sample_token",
		}

		mockSingleResult := new(mocks.SingleResult)

		// Mock the FindOne method to return a success response
		mockCollection.On("FindOne", mock.Anything, mock.MatchedBy(func(filter map[string]string) bool {
			return filter["token"] == "sample_token"
		})).Return(mockSingleResult)

		mockSingleResult.On("Decode", mock.AnythingOfType("*entities.PasswordResetToken")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*entities.PasswordResetToken)
			*arg = *expectedToken
		}).Return(nil)

		// Call the FindPasswordReset method
		result, err := repo.FindPasswordReset("sample_token")

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedToken, result)

		// Verify that the mock methods were called
		mockCollection.AssertExpectations(t)
	})

	t.Run("FindOne_Error", func(t *testing.T) {
		// Create a mock collection
		mockCollection := new(mocks.Collection)
		mockDatabase := new(mocks.Database)

		// Instantiate the repository with the mock collection
		repo := repository.NewPasswordTokenRepository(mockDatabase, mockCollection)

		// Sample token string
		tokenString := "sample_token"

		// Mock the FindOne method to return an error
		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(assert.AnError)

		mockCollection.On("FindOne", mock.Anything, mock.MatchedBy(func(filter map[string]string) bool {
			return filter["token"] == tokenString
		})).Return(mockSingleResult)

		// Call the FindPasswordReset method
		result, err := repo.FindPasswordReset(tokenString)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, assert.AnError, err)

		// Verify that the mock methods were called
		mockCollection.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}