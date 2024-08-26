package repository_test

import (
	"meleket/domain"
	"meleket/mocks"
	"meleket/repository"
	"testing"

	// "time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuite struct {
	suite.Suite
	repo       *repository.UserRepository
	collection *mocks.Collection
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.collection = new(mocks.Collection)
	suite.repo = repository.NewUserRepository(suite.collection)
}

func (suite *UserRepositorySuite) TearDownSuite() {
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestCreate() {
	suite.Run("Create success", func() {
		expectedUser := &domain.User{
			ID:   primitive.NewObjectID(),
			Name: "testuser",
			Email: "testuser@example.com",
		}

		insertResult := &mongo.InsertOneResult{InsertedID: expectedUser.ID}
		suite.collection.On("InsertOne", mock.Anything, expectedUser).Return(insertResult, nil).Once()

		err := suite.repo.Create(expectedUser)

		suite.NoError(err)
	})

	suite.Run("Create failure", func() {
		expectedUser := &domain.User{
			ID:   primitive.NewObjectID(),
			Name: "testuser",
			Email: "testuser@example.com",
		}

		suite.collection.On("InsertOne", mock.Anything, expectedUser).Return(nil, mongo.ErrClientDisconnected).Once()

		err := suite.repo.Create(expectedUser)

		suite.Error(err)
	})
}

func (suite *UserRepositorySuite) TestGetUserByUsername() {
	suite.Run("GetUserByUsername success", func() {
		expectedUser := &domain.User{
			ID:   primitive.NewObjectID(),
			Name: "testuser",
			Email: "testuser@example.com",
		}

		singleResult := new(mocks.SingleResult)

		suite.collection.On("FindOne", mock.Anything, bson.M{"name": "testuser"}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Run(func(args mock.Arguments) {
			userptr := args.Get(0).(*domain.User)
			*userptr = *expectedUser
		}).Return(nil).Once()

		result, err := suite.repo.GetUserByUsername("testuser")

		suite.NoError(err)
		suite.Equal(expectedUser, result)
	})

	suite.Run("GetUserByUsername failure", func() {
		singleResult := new(mocks.SingleResult)

		suite.collection.On("FindOne", mock.Anything, bson.M{"name": "testuser"}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Return(mongo.ErrNoDocuments).Once()

		result, err := suite.repo.GetUserByUsername("testuser")

		suite.Error(err)
		suite.Nil(result)
	})
}

func (suite *UserRepositorySuite) TestGetUserByEmail() {
    suite.Run("GetUserByEmail_Success", func() {
        email := "testuser@example.com"
        expectedUser := &domain.User{ID: primitive.NewObjectID(), Name: "testuser", Email: email}
		
		singleResult := new(mocks.SingleResult)

        suite.collection.On("FindOne", mock.Anything, bson.M{"email": email}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Run(func(args mock.Arguments) {
			userptr := args.Get(0).(*domain.User)
			*userptr = *expectedUser
		}).Return(nil).Once()

        expectedUser, err := suite.repo.GetUserByEmail(email)

        suite.NoError(err)
        suite.Equal(expectedUser, expectedUser)
    })

    suite.Run("GetUserByEmail_NotFound", func() {
        email := "nonexistent@example.com"
		singleResult := new(mocks.SingleResult)

        suite.collection.On("FindOne", mock.Anything, bson.M{"email": email}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Return(mongo.ErrNoDocuments).Once()

        expectedUser, err := suite.repo.GetUserByEmail(email)

        suite.Error(err)
        suite.Nil(expectedUser)
    })
}

// Test for GetUserByID
func (suite *UserRepositorySuite) TestGetUserByID() {
    suite.Run("GetUserByID_Success", func() {
        id := primitive.NewObjectID()
        expectedUser := &domain.User{
			ID:   id,
			Name: "testuser",
			Email: "testuser@example.com",
		}

		singleResult := new(mocks.SingleResult)

		suite.collection.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Run(func(args mock.Arguments) {
			userptr := args.Get(0).(*domain.User)
			*userptr = *expectedUser
		}).Return(nil).Once()

        user, err := suite.repo.GetUserByID(id)

        suite.NoError(err)
        suite.Equal(expectedUser, user)
    })

    suite.Run("GetUserByID_NotFound", func() {
        id := primitive.NewObjectID()

        singleResult := new(mocks.SingleResult)

		suite.collection.On("FindOne", mock.Anything, bson.M{"_id": id}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Return(mongo.ErrNoDocuments).Once()

        expectedUser, err := suite.repo.GetUserByID(id)

        suite.Error(err)
        suite.Nil(expectedUser)
    })
}

// Test for GetAllUsers
func (suite *UserRepositorySuite) TestGetAllUsers() {
    suite.Run("GetAllUsers_Success", func() {
        expectedUsers := []*domain.User{
            {ID: primitive.NewObjectID(), Name: "user1"},
            {ID: primitive.NewObjectID(), Name: "user2"},
        }

        cursorMock := new(mocks.Cursor)
        suite.collection.On("Find", mock.Anything, bson.M{}).Return(cursorMock, nil).Once()
        cursorMock.On("All",mock.Anything, mock.AnythingOfType("*[]*domain.User")).Run(func(args mock.Arguments) {
			userptr := args.Get(1).(*[]*domain.User)
			*userptr = expectedUsers
		}).Return(nil).Once()
        cursorMock.On("Close", mock.Anything).Return(nil)

        users, err := suite.repo.GetAllUsers()

        suite.NoError(err)
        suite.Equal(expectedUsers, users)
    })

    suite.Run("GetAllUsers_Error", func() {
        suite.collection.On("Find", mock.Anything, bson.M{}).Return(nil, mongo.ErrClientDisconnected).Once()

        users, err := suite.repo.GetAllUsers()

        suite.Error(err)
        suite.Nil(users)
    })
}

// Test for UpdateUser
func (suite *UserRepositorySuite) TestUpdateUser() {
    suite.Run("UpdateUser_Success", func() {
        username := "testuser"
        expectedUser := &domain.User{Name: username}

		singleResult := new(mocks.SingleResult)
        suite.collection.On("FindOneAndUpdate", mock.Anything, bson.M{"name": username}, bson.M{"$set": expectedUser}).Return(singleResult).Once()
		singleResult.On("Err").Return(nil).Once()

        err := suite.repo.UpdateUser(username, expectedUser)

        suite.NoError(err)
    })

    suite.Run("UpdateUser_Error", func() {
        username := "testuser"
        expectedUser := &domain.User{Name: username}

		singleResult := new(mocks.SingleResult)
        suite.collection.On("FindOneAndUpdate", mock.Anything, bson.M{"name": username}, bson.M{"$set": expectedUser}).Return(singleResult).Once()
		singleResult.On("Err").Return(mongo.ErrNoDocuments).Once()

        err := suite.repo.UpdateUser(username, expectedUser)

        suite.Error(err)
    })
}

// Test for DeleteUser
func (suite *UserRepositorySuite) TestDeleteUser() {
    suite.Run("DeleteUser_Success", func() {
        id := primitive.NewObjectID()

		deleteResult := &mongo.DeleteResult{DeletedCount: 1}
        suite.collection.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(deleteResult, nil).Once()

        err := suite.repo.DeleteUser(id)

        suite.NoError(err)
    })

    suite.Run("DeleteUser_NotFound", func() {
        id := primitive.NewObjectID()

		deleteResult := &mongo.DeleteResult{DeletedCount: 1}
        suite.collection.On("DeleteOne", mock.Anything, bson.M{"_id": id}).Return(deleteResult, mongo.ErrNoDocuments).Once()

        err := suite.repo.DeleteUser(id)

        suite.Error(err)
    })
}

// Test for UpdateRole
func (suite *UserRepositorySuite) TestUpdateRole() {
    suite.Run("UpdateRole_Success", func() {
        username := "testuser"
        role := "admin"

		// singleResult := new(mocks.SingleResult)
        // suite.collection.On("UpdateOne", mock.Anything, bson.M{"name": username}, bson.M{"$set": bson.M{"role": role}}).Return(singleResult, nil).Once()
		updateResult := &mongo.UpdateResult{ModifiedCount: 1}
        suite.collection.On("UpdateOne", mock.Anything, bson.M{"name": username}, bson.M{"$set": bson.M{"role": role}}).Return(updateResult, nil).Once()

        err := suite.repo.UpdateRole(username, role)

        suite.NoError(err)
    })

    suite.Run("UpdateRole_Error", func() {
        username := "testuser"
        role := "admin"

        suite.collection.On("UpdateOne", mock.Anything, bson.M{"name": username}, bson.M{"$set": bson.M{"role": role}}).Return(nil, mongo.ErrClientDisconnected).Once()

        err := suite.repo.UpdateRole(username, role)

        suite.Error(err)
    })
}



func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
