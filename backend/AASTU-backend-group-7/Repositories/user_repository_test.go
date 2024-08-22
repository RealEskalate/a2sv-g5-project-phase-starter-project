package Repositories_test

import (
	"blogapp/Domain"
	"blogapp/mocks"
	"errors"

	"log"
	"net/http"

	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	// repo "blogapp/Repositories"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	usercollection    *mocks.Collection
	refreshcollection *mocks.Collection
	repo              Domain.UserRepository
	mt                *mtest.T
}

func (suite *UserRepositoryTestSuite) SetupTest() {

	// _client := new(mocks.Client)
	suite.mt = mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock))
	_usercollection := new(mocks.Collection)
	suite.usercollection = _usercollection
	_refreshcollection := new(mocks.Collection)
	suite.refreshcollection = _refreshcollection
	// repo := repo.NewUserRepository(_usercollection, _refreshcollection)

	// suite.collection = collection
	// suite.repo = repo
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// suite.usercollection.Drop(context.Background())
	// suite.refreshcollection.Drop(context.Background())
}
func (suite *UserRepositoryTestSuite) TestCreateUser() {
	suite.Run("TestCreateUserFail", func() {
		ctx := context.Background()
		user := Domain.User{Email: "test@example.com", Password: "testpassword"}
		// expectedUser := Domain.OmitedUser{ID: primitive.NewObjectID(), Email: "test@example.com"}
		// expecte_dErr := errors.New("Email is already taken")
		expectedStatus := http.StatusBadRequest

		// Mock the CountDocuments method
		var count int64 = 1
		suite.usercollection.On("CountDocuments", ctx, bson.D{{"email", user.Email}}).Return(count, nil)

		result, err, status := suite.repo.CreateUser(ctx, &user)

		suite.Error(err)
		suite.Equal(expectedStatus, status)
		suite.Equal(Domain.OmitedUser{}, result)
	})
	suite.Run("TestCreateUseSuccess", func() {
		ctx := context.TODO()
		id := primitive.NewObjectID()
		user := Domain.User{ID: primitive.NewObjectID(), Email: "test@example.com", Password: "testpassword"}
		expectedUser := Domain.OmitedUser{ID: id, Email: "test@example.com"}
		expectedStatus := http.StatusOK

		// Mock the CountDocuments method
		var count int64 = 0
		suite.usercollection.On("CountDocuments", ctx, mock.Anything).Return(count, nil)

		//

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)
		suite.usercollection.On("InsertOne", ctx, mock.Anything).Return(&mongo.InsertOneResult{InsertedID: id}, nil)

		result, err, status := suite.repo.CreateUser(ctx, &user)
		suite.NoError(err)
		suite.Equal(expectedStatus, status)
		suite.Equal(expectedUser, result)

	})
}

func (suite *UserRepositoryTestSuite) TestGetUsersById() {
	suite.Run("Get_User_Success", func() {
		// match loged in user with the user to be fetched
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: id, Role: "user"}
		expectedUser := Domain.OmitedUser{ID: id}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, bson.D{{"_id", id}}).Return(mockSingleResult)

		result, err, status := suite.repo.GetUsersById(ctx, id, currentUser)

		suite.Nil(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)
	})

	suite.Run("Get_User_Fail", func() {
		// loged user isnot and admin and the user to be fetched is not the loged in user
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: primitive.NewObjectID(), Role: "user"}
		expectedUser := Domain.OmitedUser{ID: id}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, bson.D{{"_id", id}}).Return(mockSingleResult)

		result, err, status := suite.repo.GetUsersById(ctx, id, currentUser)

		suite.Error(err)
		suite.Equal(http.StatusForbidden, status)
		suite.Equal(Domain.OmitedUser{}, result)
	})
	suite.Run("Get_User_Success_admin", func() {
		// loged user is an admin
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: primitive.NewObjectID(), Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, bson.D{{"_id", id}}).Return(mockSingleResult)

		result, err, status := suite.repo.GetUsersById(ctx, id, currentUser)

		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)
	})
}

func (suite *UserRepositoryTestSuite) TestGetAllUsers() {
	// A testcase for the successful retrieval of users.
	suite.Run("GetAllUsersSuccess", func() {
		ctx := context.Background()
		users := []Domain.OmitedUser{
			{ID: primitive.NewObjectID(), Email: "user1@example.com"},
			{ID: primitive.NewObjectID(), Email: "user2@example.com"},
		}

		cursor := new(mocks.Cursor)
		cursor.On("All", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(1).(*[]Domain.OmitedUser)
			*userPtr = append(*userPtr, users...)
		})
		cursorCount := 0
		suite.usercollection.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cursor, nil).Once()
		cursor.On("Next", mock.Anything).Return(true).Twice()
		cursor.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = users[cursorCount]
			cursorCount++
		}).Times(len(users))
		cursor.On("Next", mock.Anything).Return(false).Once()
		cursor.On("Err").Return(nil)
		cursor.On("Close", mock.Anything).Return(nil)

		result, err, status := suite.repo.GetUsers(ctx)
		suite.NoError(err)
		suite.Equal(users[0].ID, result[0].ID)
		suite.Equal(users[1].ID, result[1].ID)
		suite.Equal(http.StatusOK, status)
	})

	// A testcase for the failure of retrieving users.
	suite.Run("GetAllUsersFailure", func() {
		ctx := context.Background()

		suite.usercollection.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(new(mocks.Cursor), mongo.ErrNoDocuments).Once()

		result, err, status := suite.repo.GetUsers(ctx)
		suite.Error(err)
		suite.Nil(result)
		suite.Equal(http.StatusNotFound, status)
	})
}

func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	// A testcase for the successful update of a task.
	suite.Run("UpdateUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: id, Role: "admin"}
		user := Domain.User{ID: id, Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, bson.D{{"_id", id}}).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.UpdateUsersById(ctx, id, user, currentUser)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)

	})

	suite.Run("UpdateUserFail", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: id, Role: "admin"}
		user := Domain.User{ID: id, Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, bson.D{{"_id", id}}).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.UpdateUsersById(ctx, id, user, currentUser)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)

	})
}
func (suite *UserRepositoryTestSuite) TestDemoteUser() {
	// A testcase for the successful update of a user.
	suite.Run("UpdateUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: primitive.NewObjectID(), Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.DemoteUser(ctx, id, currentUser)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)

	})

	suite.Run("UpdateUserFail", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: id, Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.DemoteUser(ctx, id, currentUser)
		suite.Error(err)
		suite.Equal(http.StatusForbidden, status)
		suite.NotEqual(expectedUser, result)

	})
}

func (suite *UserRepositoryTestSuite) TestPromoteUser() {
	// A testcase for the successful update of a task.
	suite.Run("UpdateUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: primitive.NewObjectID(), Role: "admin"}
		// user := Domain.User{ID: id, Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.PromoteUser(ctx, id, currentUser)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)

	})

	suite.Run("UpdateUserFail", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		currentUser := Domain.AccessClaims{ID: id, Role: "admin"}
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.PromoteUser(ctx, id, currentUser)
		suite.Error(err)
		suite.Equal(http.StatusForbidden, status)
		suite.NotEqual(expectedUser, result)

	})
}

// update by email
func (suite *UserRepositoryTestSuite) TestChangePassByEmail() {
	// A testcase for the successful update of a task.
	suite.Run("UpdateUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(1)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		result, err, status := suite.repo.ChangePassByEmail(ctx, "email", "password")
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(Domain.OmitedUser{}, result)

	})

	suite.Run("UpdateUserFail", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		expectedUser := Domain.OmitedUser{ID: id, Role: "admin"}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		count := int64(0)
		suite.usercollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: count}, nil).Once()

		_, err, status := suite.repo.ChangePassByEmail(ctx, "email", "password")
		suite.Error(err)
		suite.Equal(http.StatusBadRequest, status)

	})
}

// find by email
func (suite *UserRepositoryTestSuite) TestGetUsersByEmail() {
	suite.Run("Get_User_Success", func() {
		// match loged in user with the user to be fetched
		ctx := context.Background()
		id := primitive.NewObjectID()
		expectedUser := Domain.OmitedUser{ID: id}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})

		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)

		result, err, status := suite.repo.FindByEmail(ctx, "email")

		suite.Nil(err)
		suite.Equal(http.StatusOK, status)
		suite.Equal(expectedUser, result)
	})

	suite.Run("Get_User_Fail", func() {
		// logged user is not an admin and the user to be fetched is not the logged-in user
		ctx := context.Background()

		expectedUser := Domain.OmitedUser{}

		mockSingleResult := new(mocks.SingleResult)
		mockSingleResult.On("Decode", mock.Anything).Return(errors.New("new err")).Run(func(args mock.Arguments) {
			userPtr := args.Get(0).(*Domain.OmitedUser)
			*userPtr = expectedUser
		})
		suite.usercollection.On("FindOne", ctx, mock.Anything).Return(mockSingleResult)
		result, err, status := suite.repo.FindByEmail(ctx, "email")

		suite.Error(err)
		suite.Equal(http.StatusInternalServerError, status)
		suite.Equal(expectedUser, result) // Check that the result is an empty struct
	})
}

func (suite *UserRepositoryTestSuite) TestDeleteUser() {
	// A testcase for the successful deletion of a user.
	suite.Run("DeleteUserSuccess", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		user := Domain.AccessClaims{ID: id, Role: "admin"}
		count := int64(1)
		// suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{DeletedCount: count}, nil).Once()
		// Mock the DeleteOne method
		delResult := &mongo.DeleteResult{DeletedCount: count}
		suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		suite.refreshcollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()

		err, status := suite.repo.DeleteUsersById(ctx, id, user)
		log.Println(err)
		suite.NoError(err)
		suite.Equal(http.StatusOK, status)
	})

	// A testcase for the failure of deleting a user.
	suite.Run("DeleteUserFailure", func() {
		ctx := context.Background()
		id := primitive.NewObjectID()
		user := Domain.AccessClaims{ID: id, Role: "admin"}
		count := int64(0)
		// suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(&mongo.DeleteResult{DeletedCount: count}, nil).Once()
		// Mock the DeleteOne method
		delResult := &mongo.DeleteResult{DeletedCount: count}
		suite.usercollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()
		suite.refreshcollection.On("DeleteOne", mock.Anything, mock.Anything).Return(delResult, nil).Once()

		err, status := suite.repo.DeleteUsersById(ctx, id, user)
		log.Println(err)
		suite.Error(err)
		suite.Equal(404, status)
	})
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
