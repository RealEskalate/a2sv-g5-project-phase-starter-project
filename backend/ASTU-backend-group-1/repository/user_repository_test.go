package repository

import (
	"astu-backend-g1/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mongomocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	expectedUser = domain.User{
		ID:        "1",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "hashedpassword1",
		IsAdmin:   false,
	}
	expectedUsers = []domain.User{
		{
			Username:  "johndoe",
			Email:     "johndoe@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "hashedpassword1",
			IsAdmin:   false,
		},
		{
			Username:  "janedoe",
			Email:     "janedoe@example.com",
			FirstName: "Jane",
			LastName:  "Doe",
			Password:  "hashedpassword2",
			IsAdmin:   true,
		},
	}
)

type UserRespositoryTestSuite struct {
	suite.Suite
	coll           *mongomocks.Collection
	userRepository domain.UserRepository
}

func (suite *UserRespositoryTestSuite) SetupSuite() {
	suite.coll = &mongomocks.Collection{}
	suite.userRepository = NewUserRepository(suite.coll)
}

func (suite *UserRespositoryTestSuite) TearDownSuite() {
	suite.coll.AssertExpectations(suite.T())
}

func (suite *UserRespositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	suite.coll.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil)
	result, err := suite.userRepository.Create(expectedUser)
	assert.NoError(err)
	assert.Equal(result, expectedUser)
}

func (suite *UserRespositoryTestSuite) TestGet() {
	assert := assert.New(suite.T())
	suite.T().Parallel()
	suite.T().Run("Getting all users", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		for i, user := range expectedUsers {
			cur.On("Next", mock.Anything).Return(i < len(expectedUsers)).Once()
			cur.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*domain.User)
				*arg = user
			}).Return(nil).Once()
		}
		cur.On("Next", mock.Anything).Return(false).Once()
		suite.coll.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cur, nil)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.userRepository.Get(domain.UserFilterOption{})
		assert.NoError(err)
		assert.Equal(expectedUsers, result)
	})
	suite.T().Run("Getting by Username", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			*arg = expectedUser
		}).Return(nil)
		suite.coll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.userRepository.Get(domain.UserFilterOption{Username: expectedUser.Username})
		assert.NoError(err)
		assert.Equal(expectedUser, result[0])
	})
	suite.T().Run("Getting by Email", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			*arg = expectedUser
		}).Return(nil)
		suite.coll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.userRepository.Get(domain.UserFilterOption{Email: expectedUser.Email})
		assert.NoError(err)
		assert.Equal(expectedUser, result[0])
	})
	suite.T().Run("Getting by Id", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.User)
			*arg = expectedUser
		}).Return(nil)
		suite.coll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.userRepository.Get(domain.UserFilterOption{UserID: expectedUser.ID})
		assert.NoError(err)
		assert.Equal(expectedUser, result[0])
	})
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRespositoryTestSuite))
}
