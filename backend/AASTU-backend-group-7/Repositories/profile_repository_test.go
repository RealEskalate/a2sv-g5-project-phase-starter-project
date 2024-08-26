package Repositories_test

import (
	"blogapp/Domain"
	"blogapp/mocks"
	"context"
	"errors"
	"log"
	"net/http"
	"testing"

	repo "blogapp/Repositories"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepositoryTestSuite struct {
	suite.Suite
	repo            Domain.ProfileRepository
	usercollection  *mocks.Collection
	tokencollection *mocks.Collection
	postcollection  *mocks.Collection
}

func (suite *ProfileRepositoryTestSuite) SetupTest() {
	// Initialize the repository and mock dependencies
	suite.usercollection = new(mocks.Collection)
	suite.tokencollection = new(mocks.Collection)
	suite.postcollection = new(mocks.Collection)
	suite.repo = repo.NewProfileRepository(suite.usercollection, suite.tokencollection, suite.postcollection)
}

func (suite *ProfileRepositoryTestSuite) TearDownTest() {
	suite.usercollection = nil
	// Cleanup resources if needed
}

func (suite *ProfileRepositoryTestSuite) TestDeleteProfile() {
	suite.Run("DeleteUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		user := Domain.AccessClaims{ID: id, Role: "admin"}
		count := int64(1)

		// Mock the DeleteOne method
		delResult := &mongo.DeleteResult{DeletedCount: count}
		suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		suite.tokencollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		err, status := suite.repo.DeleteProfile(ctx, id, user)
		log.Println(err)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
	})

	suite.Run("DeleteUserFailure", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		user := Domain.AccessClaims{ID: id, Role: "admin"}
		count := int64(0)

		// Mock the DeleteOne method
		delResult := &mongo.DeleteResult{DeletedCount: count}
		suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		suite.tokencollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()

		err, status := suite.repo.DeleteProfile(ctx, id, user)
		log.Println(err)
		suite.Error(err)
		suite.Equal(404, status)
	})
}

func (suite *ProfileRepositoryTestSuite) TestGetProfile() {

	suite.Run("TestGetUserFail", func() {
		ctx := context.Background()
		expectedUser := Domain.OmitedUser{}
		id := primitive.ObjectID{}
		mockSingleResult := new(mocks.SingleResult)

		// Simulate error from Decode
		mockSingleResult.On("Decode", mock.Anything).Return(errors.New("new err"))

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		// Call the method under test
		result, err, status := suite.repo.GetProfile(ctx, id, Domain.AccessClaims{ID: id})

		// Assertions
		suite.Error(err)
		suite.Equal(http.StatusNotFound, status)
		suite.Equal(expectedUser, result)
	})

	suite.Run("TestGetUserSuccess", func() {
		ctx := context.Background()
		expectedUser := Domain.OmitedUser{} // Populate with expected values as needed
		id := primitive.ObjectID{}
		mockSingleResult := new(mocks.SingleResult)

		// Simulate successful decoding
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		// Call the method under test
		result, err, status := suite.repo.GetProfile(ctx, id, Domain.AccessClaims{ID: id})

		// Assertions
		suite.Error(err)
		suite.NotEqual(http.StatusOK, status)
		suite.Equal(expectedUser, result) // Compare to expected user
	})
}
func (suite *ProfileRepositoryTestSuite) TestUpdateProfile() {

	suite.Run("UpdateUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: id, Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}
		user := Domain.User{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.UpdateProfile(ctx, id, user, currentUser)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.NotEqual(expectedUser, result)
	})

	suite.Run("UpdateUserFail", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: primitive.NewObjectID(), Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}
		user := Domain.User{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.UpdateProfile(ctx, id, user, currentUser)
		suite.Error(err)
		suite.Equal(http.StatusForbidden, status)
		suite.NotEqual(expectedUser, result)

	})
}

func TestProfileRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProfileRepositoryTestSuite))
}
