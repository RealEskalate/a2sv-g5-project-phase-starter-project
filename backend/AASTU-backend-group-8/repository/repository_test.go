package repository_test

import (
	"meleket/domain"
	"meleket/repository"
	"meleket/mocks"
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
		user := &domain.User{
			ID:   primitive.NewObjectID(),
			Name: "testuser",
			Email: "testuser@example.com",
		}

		insertResult := &mongo.InsertOneResult{InsertedID: user.ID}
		suite.collection.On("InsertOne", mock.Anything, user).Return(insertResult, nil).Once()

		err := suite.repo.Create(user)

		suite.NoError(err)
	})

	suite.Run("Create failure", func() {
		user := &domain.User{
			ID:   primitive.NewObjectID(),
			Name: "testuser",
			Email: "testuser@example.com",
		}

		suite.collection.On("InsertOne", mock.Anything, user).Return(nil, mongo.ErrClientDisconnected).Once()

		err := suite.repo.Create(user)

		suite.Error(err)
	})
}

func (suite *UserRepositorySuite) TestGetUserByUsername() {
	suite.Run("GetUserByUsername success", func() {
		user := &domain.User{
			ID:   primitive.NewObjectID(),
			Name: "testuser",
			Email: "testuser@example.com",
		}

		singleResult := new(mocks.SingleResult)

		suite.collection.On("FindOne", mock.Anything, bson.M{"name": "testuser"}).Return(singleResult).Once()
		singleResult.On("Decode", mock.AnythingOfType("*domain.User")).Run(func(args mock.Arguments) {
			userptr := args.Get(0).(*domain.User)
			*userptr = *user
		}).Return(nil).Once()

		result, err := suite.repo.GetUserByUsername("testuser")

		suite.NoError(err)
		suite.Equal(user, result)
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


func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
