package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"context"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetUserProfile(t *testing.T) {
    mockCollection := new(mocks.Collection)
    mockContext := context.TODO()
    mockDatabase := new(mocks.Database)

    repo := repository.NewProfileRepository(mockContext, mockDatabase, mockCollection)

    // Create a mock profile document
    expectedProfile := &entities.Profile{
		UserID: primitive.NewObjectID(),
		Bio: "This is a test bio",	
		ContactInfo: entities.ContactInfo{
			Email: "john.doe@example.com",
			PhoneNumber: "1234567890",
			Address: "123 Main St",
		},
		ProfilePicture: "https://example.com/profile.jpg",
	}

	userIDHex := expectedProfile.UserID.Hex()

    // Create a single result with the mock profile
    mockSingleResult := new(mocks.SingleResult)
	mockSingleResult.On("Decode", mock.AnythingOfType("*entities.Profile")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.Profile)
		*arg = *expectedProfile
	}).Return(nil)

   mockCollection.On("FindOne", mock.Anything, bson.D{{Key: "userId", Value: expectedProfile.UserID}}, mock.Anything).Return(mockSingleResult)

    // Call the GetUserProfile method
	actualProfile, err := repo.GetUserProfile(userIDHex)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the returned profile matches the expected one
	assert.Equal(t, expectedProfile, actualProfile)

	// Ensure the mock methods were called as expected
	mockSingleResult.AssertExpectations(t)

	mockCollection.AssertExpectations(t)
}

func TestCreateUserProfile(t *testing.T) {
	mockCollection := new(mocks.Collection)
	mockContext := context.TODO()
	mockDatabase := new(mocks.Database)

	repo := repository.NewProfileRepository(mockContext, mockDatabase, mockCollection)

	userID := primitive.NewObjectID()

	// Create a mock profile document
	mockProfile := &entities.Profile{
		UserID: userID,
		Bio: "This is a test bio",	
		ContactInfo: entities.ContactInfo{
			Email: "johndoe@gmail.com",
			PhoneNumber: "1234567890",
			Address: "123 Main St",
		},
		ProfilePicture: "https://example.com/profile.jpg",
	}

	// Mock the FindOne method to return an error
	mockCollection.On("FindOne", mockContext, bson.D{{Key: "userId", Value: userID}}).Return(nil)

	// Mock the InsertOne method to return the mockProfile
	mockCollection.On("InsertOne", mockContext, mockProfile).Return(&mongo.SingleResult{}, nil)

	result, err := repo.CreateUserProfile(mockProfile)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockCollection.AssertExpectations(t)
}
