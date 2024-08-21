package Repositories_test

import (
	"blogapp/Domain"
	"blogapp/mocks"
	"net/http"

	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	repo "blogapp/Repositories"
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
	repo := repo.NewUserRepository(_usercollection, _refreshcollection)

	// suite.collection = collection
	suite.repo = repo
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// suite.usercollection.Drop(context.Background())
	// suite.refreshcollection.Drop(context.Background())
}
func (suite *UserRepositoryTestSuite) TestCreateUser() {
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
}

// func (suite *UserRepositoryTestSuite) TestCreateUser_NewUser() {
// 	ctx := context.Background()
// 	user := Domain.User{Email: "test@example.com", Password: "testpassword"}
// 	expectedUser := Domain.OmitedUser{ID: primitive.NewObjectID(), Email: "test@example.com"}
// 	expectedErr := nil
// 	expectedStatus := http.StatusOK

// 	// Mock the CountDocuments method
// 	mockCollection := new(mocks.Collection)
// 	suite.usercollection = mockCollection
// 	mockCollection.On("CountDocuments", ctx, bson.D{{"email", user.Email}}).Return(0, nil)

// 	result, err, status := suite.repo.CreateUser(ctx, &user)

// 	suite.NoError(err)
// 	suite.Equal(expectedStatus, status)
// 	suite.Equal(expectedUser, result)
// }

func (suite *UserRepositoryTestSuite) TestGetUsersByIdFail() {
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
}

func (suite *UserRepositoryTestSuite) TestGetUsersByIdPass() {
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
}

func (suite *UserRepositoryTestSuite) TestGetUsersById() {
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
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
