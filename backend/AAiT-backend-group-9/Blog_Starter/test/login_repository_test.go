package test

import (
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoginRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.LoginRepository
	db         *mongo.Database
}

func (suite *LoginRepositorySuit) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	repository := repository.NewLoginRepository(db, "user")

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *LoginRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}

func (suite *LoginRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("user").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// create a user using
	userRepo := repository.NewUserRepository(suite.db, "user")
	user := domain.User{
		UserID:   primitive.NewObjectID(),
		Email:    "eyob@gmail.com",
		Password: "password",
	}
	userRepo.CreateUser(context.Background(), &user)

}

// implement the test cases here

// test the Login function
func (suite *LoginRepositorySuit) TestLogin() {
	// first create a user then update it and check it

	user := domain.UserLogin{
		Email:    "eyob@gmail.com",
		Password: "Password",
	}

	loginUser, err := suite.repository.Login(context.Background(), &user)
	if err != nil {
		suite.T().Fatal(err)
	}

	// check access token and refresh token is exist

	suite.NotNil(loginUser.AccessToken)
	suite.NotNil(loginUser.RefreshToken)
}

func TestLoginRepositorySuit(t *testing.T) {
	suite.Run(t, new(LoginRepositorySuit))
}
