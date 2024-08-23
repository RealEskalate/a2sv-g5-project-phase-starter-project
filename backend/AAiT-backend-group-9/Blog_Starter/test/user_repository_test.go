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

type UserRepositorySuit struct {
	suite.Suite
	// the funcionalities we need to test
	repository domain.UserRepository
	db         *mongo.Database
}

func (suite *UserRepositorySuit) SetupSuite() {
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
	repository := repository.NewUserRepository(db, "user")

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.repository = repository
	suite.db = db
}

func (suite *UserRepositorySuit) TearDownSuite() {
	// we need to drop the table we used in the tests
	defer suite.db.Drop(context.Background())
}

func (suite *UserRepositorySuit) SetupTest() {
	// this function runs before every test in the suite
	// we need to clear the table before every test
	_, err := suite.db.Collection("user").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

// implement the test cases here

// test the UpdateSignup function
func (suite *UserRepositorySuit) TestUpdateSignup() {
	// first create a user then update it and check it

	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmai.com",
		Password: "password",
	}

	_, err := suite.repository.CreateUser(context.Background(), user)
	if err != nil {
		suite.T().Fatal(err)
	}

	// now update the user
	user.Username = "newTest"
	err = suite.repository.UpdateSignup(context.Background(), user)
	if err != nil {
		suite.T().Fatal(err)
	}
}

// test the CreateUser function
func (suite *UserRepositorySuit) TestCreateUser() {
	// testing the create user functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)
	suite.NotNil(user.CreatedAt)

}

// test the DeleteUser function
func (suite *UserRepositorySuit) TestDeleteUser() {
	// testing the delete user functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	err = suite.repository.DeleteUser(context.Background(), user.UserID.Hex())
	suite.NoError(err)
}

// test the GetAllUser function
func (suite *UserRepositorySuit) TestGetAllUser() {
	// testing the get all user functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob.@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	users, err := suite.repository.GetAllUser(context.Background())
	suite.NoError(err)
	suite.NotNil(users)
	// check length of the ueers
	suite.Equal(1, len(users))
}

// test the GetUserByEmail function
func (suite *UserRepositorySuit) TestGetUserByEmail() {
	// testing the get user by email functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	foundUser, err := suite.repository.GetUserByEmail(context.Background(), user.Email)
	suite.NoError(err)
	suite.NotNil(foundUser)
	suite.Equal(user.Email, foundUser.Email)
}

// test the GetUserByID function

func (suite *UserRepositorySuit) TestGetUserByID() {
	// testing the get user by ID functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	foundUser, err := suite.repository.GetUserByID(context.Background(), user.UserID.Hex())
	suite.NoError(err)
	suite.NotNil(foundUser)
	suite.Equal(user.UserID, foundUser.UserID)
}

// test the UpdateProfile function
func (suite *UserRepositorySuit) TestUpdateProfile() {
	// testing the update profile functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	userUpdate := &domain.UserUpdate{
		Username: "newTest",
		Name:     "newName",
		Bio:      "newBio",
		ContactInfo: domain.ContactInfo{
			Phone:   "123456789",
			Address: "newAddress",
		},
	}

	updatedUser, err := suite.repository.UpdateProfile(context.Background(), userUpdate, user.UserID.Hex())
	suite.NoError(err)
	suite.NotNil(updatedUser)
	suite.Equal(userUpdate.Username, updatedUser.Username)
	suite.Equal(userUpdate.Name, updatedUser.Name)
	suite.Equal(userUpdate.Bio, updatedUser.Bio)
	suite.Equal(userUpdate.ContactInfo, updatedUser.ContactInfo)
}

// test the UpdateProfilePicture function
func (suite *UserRepositorySuit) TestUpdateProfilePicture() {
	// testing the update profile picture functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	profilePicPath := "path/to/profile/picture"
	updatedUser, err := suite.repository.UpdateProfilePicture(context.Background(), profilePicPath, user.UserID.Hex())
	suite.NoError(err)
	suite.NotNil(updatedUser)
	suite.Equal(profilePicPath, updatedUser.ProfilePicture)
}

// test the UpdateToken function
func (suite *UserRepositorySuit) TestUpdateToken() {
	// testing the update token functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	accessToken := "newAccessToken"
	refreshToken := "newRefreshToken"
	updatedUser, err := suite.repository.UpdateToken(context.Background(), accessToken, refreshToken, user.UserID.Hex())
	suite.NoError(err)
	suite.NotNil(updatedUser)
	suite.Equal(accessToken, updatedUser.AccessToken)
	suite.Equal(refreshToken, updatedUser.RefreshToken)
}

// test the UpdateRole function
func (suite *UserRepositorySuit) TestUpdateRole() {
	// testing the update role functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyob@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	role := "newRole"
	updatedUser, err := suite.repository.UpdateRole(context.Background(), role, user.UserID.Hex())
	suite.NoError(err)
	suite.NotNil(updatedUser)
	suite.Equal(role, updatedUser.Role)
}

// test the UpdatePassword function
func (suite *UserRepositorySuit) TestUpdatePassword() {
	// testing the update password functionality
	user := &domain.User{
		UserID:   primitive.NewObjectID(),
		Username: "test",
		Email:    "eyda@gmail.com",
		Password: "password",
	}

	user, err := suite.repository.CreateUser(context.Background(), user)
	suite.NoError(err)
	suite.NotNil(user)

	password := "newPassword"
	updatedUser, err := suite.repository.UpdatePassword(context.Background(), password, user.UserID.Hex())
	suite.NoError(err)
	suite.NotNil(updatedUser)
	suite.Equal(password, updatedUser.Password)
}

// run the siuite
func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuit))
}
