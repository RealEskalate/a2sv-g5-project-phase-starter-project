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

var MockUserData = []domain.User{
	{Username: "testuser1", Email: "testuser1@gmail.com", Password: "password"},
	{Username: "testuser2", Email: "testuser2@gmail.com", Password: "password"},
	{Username: "testuser3", Email: "testuser3@gmail.com", Password: "password"},
	{Username: "testuser4", Email: "testuser4@gmail.com", Password: "password"},
	{Username: "testuser5", Email: "testuser5@gmail.com", Password: "password"},
	{Username: "testuser6", Email: "testuser6@gmail.com", Password: "password"},
}

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

func (suite *UserRepositoryTestSuite) TestCreateUser_Negative_DuplicateEmail() {
	user := domain.User{
		Username: "testuser",
		Email:    "user@gmail.com",
		Password: "password",
	}

	err := suite.UserRepository.CreateUser(context.Background(), &user)
	suite.Nil(err, "no error when creating user")

	newUser := domain.User{
		Username: "newtestuser",
		Email:    user.Email,
		Password: "password",
	}
	err = suite.UserRepository.CreateUser(context.Background(), &newUser)

	suite.NotNil(err, "error when creating user with duplicate email")
	suite.Equal(err.GetCode(), domain.ERR_CONFLICT, "error code is conflict")
}

func (suite *UserRepositoryTestSuite) TestCreateUser_Negative_DuplicateUsername() {
	user := domain.User{
		Username: "testuser",
		Email:    "user@gmail.com",
		Password: "password",
	}

	err := suite.UserRepository.CreateUser(context.Background(), &user)
	suite.Nil(err, "no error when creating user")

	newUser := domain.User{
		Username: user.Username,
		Email:    "newemail@gmail.com",
		Password: "password",
	}
	err = suite.UserRepository.CreateUser(context.Background(), &newUser)

	suite.NotNil(err, "error when creating user with duplicate username")
	suite.Equal(err.GetCode(), domain.ERR_CONFLICT, "error code is conflict")
}

func (suite *UserRepositoryTestSuite) TestFindUser_Positive() {
	for _, user := range MockUserData {
		suite.UserRepository.CreateUser(context.Background(), &user)
	}

	// find by both email and username
	for _, user := range MockUserData {
		foundUser, err := suite.UserRepository.FindUser(context.Background(), &user)
		suite.Nil(err, "no error when finding user")
		suite.Equal(user.Username, foundUser.Username, "username matches")
		suite.Equal(user.Email, foundUser.Email, "email matches")
		suite.Equal(user.Password, foundUser.Password, "password matches")
	}

	// find by username
	for _, user := range MockUserData {
		foundUser, err := suite.UserRepository.FindUser(context.Background(), &domain.User{Username: user.Username})
		suite.Nil(err, "no error when finding user")
		suite.Equal(user.Username, foundUser.Username, "username matches")
		suite.Equal(user.Email, foundUser.Email, "email matches")
		suite.Equal(user.Password, foundUser.Password, "password matches")
	}

	// find by email
	for _, user := range MockUserData {
		foundUser, err := suite.UserRepository.FindUser(context.Background(), &domain.User{Email: user.Email})
		suite.Nil(err, "no error when finding user")
		suite.Equal(user.Username, foundUser.Username, "username matches")
		suite.Equal(user.Email, foundUser.Email, "email matches")
		suite.Equal(user.Password, foundUser.Password, "password matches")
	}
}

func (suite *UserRepositoryTestSuite) TestFindUser_Negative_UserNotFound() {
	for _, user := range MockUserData {
		suite.UserRepository.CreateUser(context.Background(), &user)
	}

	_, err := suite.UserRepository.FindUser(context.Background(), &domain.User{Username: "testuser99"})
	suite.NotNil(err, "error when user not found")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")

	_, err = suite.UserRepository.FindUser(context.Background(), &domain.User{Email: "testuser99@gmail.com"})
	suite.NotNil(err, "error when user not found")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

func (suite *UserRepositoryTestSuite) TestSetRefreshToken_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)

	newRefreshToken := "this is a. kinda valid refresh token. it has the two dots"
	err := suite.UserRepository.SetRefreshToken(context.Background(), &user, newRefreshToken)
	suite.Nil(err, "no error when setting refresh token")
}

func (suite *UserRepositoryTestSuite) TestSetRefreshToken_Negative() {
	user := MockUserData[0]
	// user not created

	newRefreshToken := "this is a. kinda valid refresh token. it has the two dots"
	err := suite.UserRepository.SetRefreshToken(context.Background(), &user, newRefreshToken)
	suite.NotNil(err, "no error when setting refresh token")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

func (suite *UserRepositoryTestSuite) TeardownSuite() {
	initdb.DisconnectDB(suite.client)
}

func TestUserRepositry(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
