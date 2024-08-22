package tests

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	initdb "blog_api/infrastructure/db"
	"blog_api/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	client         *mongo.Client
	collection     *mongo.Collection
	UserRepository *repository.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	// setup the database connection
	err := env.LoadEnvironmentVariables("../.env")
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	client, err := initdb.ConnectDB(env.ENV.DB_ADDRESS, env.ENV.TEST_DB_NAME)
	if err != nil {
		suite.T().Fatal(err)
		return
	}

	suite.client = client
	suite.collection = client.Database(env.ENV.TEST_DB_NAME).Collection(domain.CollectionUsers)
	suite.UserRepository = repository.NewUserRepository(suite.collection)
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.collection.DeleteMany(context.Background(), bson.D{})
}

func (suite *UserRepositoryTestSuite) TestCreateUser_Positive() {
	user := domain.User{
		Username: "testuser",
		Email:    "user@gmail.com",
		Password: "password",
	}

	err := suite.UserRepository.CreateUser(context.Background(), &user)
	suite.Nil(err, "no error when creating user")

	var createdUser domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&createdUser)
	suite.Equal(user.Username, createdUser.Username, "username matches")
	suite.Equal(user.Email, createdUser.Email, "email matches")
	suite.Equal(user.Password, createdUser.Password, "password matches")
}

func (suite *UserRepositoryTestSuite) TeardownSuite() {
	initdb.DisconnectDB(suite.client)
}

func TestUserRepositry(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
