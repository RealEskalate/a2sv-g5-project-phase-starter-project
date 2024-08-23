package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateUser(t *testing.T) {

	// Create a mock collection
	mockCollection := mongo.Database("mock-db").Collection("mock-collection")

	// Create an instance of userRepository with the mock collection
	userRepo := repository.NewUserRepository(mockCollection)

	// Create a dummy user entity
	user := &entities.User{
		ID:         primitive.NewObjectID(),
		Username:   "testuser",
		Email:      "testuser@example.com",
		Password:   "password123",
		Profile:    entities.Profile{},
		IsVerified: false,
		Role:       "user",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Set up the expectations for the mock
	insertResult := &mongo.InsertOneResult{InsertedID: user.ID}
	mockCollection.On("InsertOne", mock.Anything, user, mock.Anything).Return(insertResult, nil)

	// Call the CreateUser method
	createdUser, err := userRepo.CreateUser(user)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, user.ID, createdUser.ID)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Email, createdUser.Email)

	// Ensure that the mock expectations were met
	mockCollection.AssertExpectations(t)
}

func TestFindUserByEmail(t *testing.T) {
	
	// Create a mock collection
	mockCollection := mongo.Database("mock-db").Collection("mock-collection")

	// Create an instance of userRepository with the mock collection
	userRepo := repository.NewUserRepository(mockCollection)

	// Create a dummy user entity
	user := &entities.User{
		ID:         primitive.NewObjectID(),
		Username:   "testuser",
		Email:      "test@gmail.com",
		Password:   "password123",
		Profile:    entities.Profile{},
		IsVerified: false,
		Role:       "user",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Set up the expectations for the mock
	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(nil)

	// Call the FindUserByEmail method
	foundUser, err := userRepo.FindUserByEmail(user.Email)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Email, foundUser.Email)
	
	// Ensure that the mock expectations were met
	mockCollection.AssertExpectations(t)
}
