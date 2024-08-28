package repository_test

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/mongo/mocks"
	"backend-starter-project/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T)  {
	
	mockCollection := new(mocks.Collection)
	repo := repository.NewUserRepository(mockCollection)

	mockUser := &entities.User{
		ID: primitive.NewObjectID(),
		Username: "test",
		Email: "test@gmail.com",
		Password: "password",
		Profile: entities.Profile{},
	}

	mockCollection.On("InsertOne", mock.Anything, mockUser).Return(nil, nil)

	result, err := repo.CreateUser(mockUser)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockCollection.AssertExpectations(t)
}

func TestFindUserByEmail(t *testing.T) {
    mockCollection := new(mocks.Collection)
    repo := repository.NewUserRepository(mockCollection)

    email := "test@example.com"

    expectedUser := &entities.User{
        Email: email,
        Username:  "Test User",
    }

	mockSingleresult := new(mocks.SingleResult)
    mockCollection.On("FindOne", mock.Anything, bson.M{"email": email}).Return(mockSingleresult, nil)
    mockSingleresult.On("Decode", mock.AnythingOfType("*entities.User")).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*entities.User)
        *arg = *expectedUser
    }).Return(nil)

    user, err := repo.FindUserByEmail(email)

    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, expectedUser.Email, user.Email)
    assert.Equal(t, expectedUser.Username, user.Username)

    mockCollection.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockCollection := new(mocks.Collection)
	repo := repository.NewUserRepository(mockCollection)

	userId := primitive.NewObjectID().Hex()

	mockDeleteCount := int64(1)
	mockCollection.On("DeleteOne", mock.Anything, bson.M{"_id": userId}).Return(mockDeleteCount, nil)

	err := repo.DeleteUser(userId)

	assert.NoError(t, err)

	mockCollection.AssertExpectations(t)
}

func TestFindUserById(t *testing.T) {
	mockCollection := new(mocks.Collection)
	repo := repository.NewUserRepository(mockCollection)

	userId := primitive.NewObjectID().Hex()

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil{
		t.Fatalf("error occurred while converting userId to objectId")
	}

	expectedUser := &entities.User{
		ID: objectId,
		Username: "test",
		Email: "test@gmail.com",
		Password: "password",
		Profile: entities.Profile{},
	}

	mockSingleresult := new(mocks.SingleResult)
	mockCollection.On("FindOne", mock.Anything, bson.M{"_id": expectedUser.ID}).Return(mockSingleresult, nil)
	mockSingleresult.On("Decode", mock.AnythingOfType("*entities.User")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.User)
		*arg = *expectedUser
	}).Return(nil)

	user, err := repo.FindUserById(userId)

	assert.NoError(t, err)

	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Password, user.Password)
	assert.Equal(t, expectedUser.Profile, user.Profile)

	mockCollection.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T)  {
	mockCollection := new(mocks.Collection)
	repo := repository.NewUserRepository(mockCollection)

	userID := "60c72b2f9b1e8f3b5b7c8f8d"

	objectID, _ := primitive.ObjectIDFromHex(userID)

	updateUser := &entities.User{
		ID: objectID,
		Username: "test",
		Email: "test@gmail.com",
		Password: "password",
		Profile: entities.Profile{},
	}

	filter := bson.M{
		"_id": objectID,
	}

	update := bson.M{
		"$set": updateUser,
	}

    mockCollection.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	_, err := repo.UpdateUser(updateUser)

	assert.NoError(t, err)
	
	mockCollection.AssertExpectations(t)
}

func TestMarkUserAsVerified(t *testing.T) {
    mockCollection := new(mocks.Collection)
    repo := repository.NewUserRepository(mockCollection)

    email := "test@example.com"
    userID := primitive.NewObjectID()
    expectedUser := &entities.User{
        ID:    userID,
        Email: email,
    }

	mockSingleresult := new(mocks.SingleResult)
    // Mock FindUserByEmail
    mockCollection.On("FindOne", mock.Anything, bson.M{"email": email}).Return(mockSingleresult, nil)

    // Mock UpdateOne
    update := bson.M{"$set": bson.M{"isVerified": true}}
    mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": userID}, update).Return(nil, nil).Once()
	mockSingleresult.On("Decode", mock.AnythingOfType("*entities.User")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.User)
		*arg = *expectedUser
	}).Return(nil)
	
    err := repo.MarkUserAsVerified(email)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}