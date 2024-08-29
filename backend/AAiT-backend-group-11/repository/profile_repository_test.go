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

func TestCreateUserProfile(t *testing.T) {
    mockCollection := new(mocks.Collection)
    repo := repository.NewProfileRepository(context.TODO(), nil, mockCollection)

    t.Run("User ID is missing", func(t *testing.T) {
        profile := &entities.Profile{
            UserID: primitive.NilObjectID,
        }

        createdProfile, err := repo.CreateUserProfile(profile)
        assert.Nil(t, createdProfile)
        assert.EqualError(t, err, "user id is required")
    })

    t.Run("Profile already exists", func(t *testing.T) {
        profile := &entities.Profile{
            UserID: primitive.NewObjectID(),
        }

        mockSingleResult := new(mocks.SingleResult)
        mockSingleResult.On("Err").Return(nil)

        mockCollection.On("FindOne", mock.Anything, bson.D{{"userId", profile.UserID}}, mock.Anything).Return(mockSingleResult)

        createdProfile, err := repo.CreateUserProfile(profile)
        assert.Nil(t, createdProfile)
        assert.EqualError(t, err, "profile already exists")

        mockCollection.AssertExpectations(t)
        mockSingleResult.AssertExpectations(t)
    })

    t.Run("Profile is successfully created", func(t *testing.T) {
        profile := &entities.Profile{
            UserID: primitive.NewObjectID(),
            Bio:    "This is a test bio",
            ContactInfo: entities.ContactInfo{
                Email:       "john.doe@example.com",
                PhoneNumber: "1234567890",
                Address:     "123 Main St",
            },
            ProfilePicture: "https://example.com/profile.jpg",
        }

        mockSingleResult := new(mocks.SingleResult)
        mockSingleResult.On("Err").Return(mongo.ErrNoDocuments)

        mockCollection.On("FindOne", mock.Anything, bson.D{{"userId", profile.UserID}}, mock.Anything).Return(mockSingleResult)
        mockCollection.On("InsertOne", mock.Anything, profile, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

        createdProfile, err := repo.CreateUserProfile(profile)
        assert.NoError(t, err)
        assert.NotNil(t, createdProfile)
        assert.Equal(t, profile, createdProfile)

        mockCollection.AssertExpectations(t)
        mockSingleResult.AssertExpectations(t)
    })
}

func TestUpdateUserProfile(t *testing.T) {
	mockCollection := new(mocks.Collection)
	repo := repository.NewProfileRepository(context.TODO(), nil, mockCollection)

	t.Run("User ID is missing", func(t *testing.T) {
		profile := &entities.Profile{
			UserID: primitive.NilObjectID,
		}

		updatedProfile, err := repo.UpdateUserProfile(profile)
		assert.Nil(t, updatedProfile)
		assert.EqualError(t, err, "user id is required")
	})

	t.Run("Profile is successfully updated", func(t *testing.T) {
		profile := &entities.Profile{
			UserID: primitive.NewObjectID(),
			Bio:    "This is a test bio",
			ContactInfo: entities.ContactInfo{
				Email:       "testupdate@gmail.com",
				PhoneNumber: "1234567890",
				Address:     "123 Main St",
			},
			ProfilePicture: "https://example.com/profile.jpg",
		}

		mockCollection.On("UpdateOne", mock.Anything, bson.D{{"userId", profile.UserID}}, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{}, nil)

		updatedProfile, err := repo.UpdateUserProfile(profile)
		assert.NoError(t, err)
		assert.NotNil(t, updatedProfile)
		assert.Equal(t, profile, updatedProfile)

		mockCollection.AssertExpectations(t)
	})
}

func TestDeleteUserProfile(t *testing.T) {
    mockCollection := new(mocks.Collection)
    repo := repository.NewProfileRepository(context.TODO(), nil, mockCollection)

    t.Run("Successful deletion", func(t *testing.T) {
        userID := primitive.NewObjectID().Hex()
        objectID, _ := primitive.ObjectIDFromHex(userID)
        filter := bson.D{{"userId", objectID}}

		mockDeleteCount := int64(1)
        mockCollection.On("DeleteOne", mock.Anything, filter, mock.Anything).Return(mockDeleteCount, nil)

        err := repo.DeleteUserProfile(userID)
        assert.NoError(t, err)

        mockCollection.AssertExpectations(t)
    })

    t.Run("Error during deletion", func(t *testing.T) {
        userID := primitive.NewObjectID().Hex()
        objectID, _ := primitive.ObjectIDFromHex(userID)
        filter := bson.D{{"userId", objectID}}

		mockDeleteCount := int64(0)
        mockCollection.On("DeleteOne", mock.Anything, filter, mock.Anything).Return(mockDeleteCount, mongo.ErrNoDocuments)

        err := repo.DeleteUserProfile(userID)
        assert.Error(t, err)
        assert.Equal(t, mongo.ErrNoDocuments, err)

        mockCollection.AssertExpectations(t)
    })
}