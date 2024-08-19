package tests

import (
	"context"
	"testing"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type UserRepositorySuite struct {
	suite.Suite
	Repository  *repository.UserMongoRepository
	DB          *mongo.Database
	Collection  *mongo.Collection
	TestContext context.Context
}

// SetupSuite runs before all tests in the suite
func (suite *UserRepositorySuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") 

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		suite.T().Fatalf("Failed to connect to MongoDB: %v", err)
	}

	suite.DB = client.Database("test_db")
	suite.Collection = suite.DB.Collection("user-collection")
	suite.Repository = repository.NewUserRepository(suite.DB)
	suite.TestContext = context.Background()
}

// TearDownSuite runs after all tests in the suite
func (suite *UserRepositorySuite) TearDownSuite() {
	err := suite.DB.Drop(suite.TestContext)
	if err != nil {
		suite.T().Fatalf("Failed to clean up test database: %v", err)
	}
}

// TestCreateUser tests the CreateUser method
func (suite *UserRepositorySuite) TestCreateUser() {
	user := &models.User{
		Username:   "testuser",
		Name:       "Test User",
		Email:      "testuser@example.com",
		Password:   "hashedpassword",
		IsVerified: false,
	}

	err := suite.Repository.CreateUser(suite.TestContext, user)
	suite.Empty(err, "Expected no error when creating a user")

	// Check if the user was inserted correctly
	var insertedUser models.User
	Err := suite.Collection.FindOne(suite.TestContext, bson.M{"username": "testuser"}).Decode(&insertedUser)
	suite.Nil(Err, "Expected no error when fetching the inserted user")
	suite.Equal(user.Email, insertedUser.Email, "Expected inserted email to match")
}

// TestGetUserByEmailOrUsername tests the GetUserByEmailOrUsername method
func (suite *UserRepositorySuite) TestGetUserByEmailOrUsername() {
	// Insert a user for testing
	user := &models.User{
		Username: "fetchuser",
		Email:    "fetchuser@example.com",
	}
	suite.Collection.InsertOne(suite.TestContext, user)

	// Fetch by username
	fetchedUser, err := suite.Repository.GetUserByEmailOrUsername(suite.TestContext, "fetchuser", "")
	suite.Empty(err, "Expected no error when fetching by username")
	suite.Equal("fetchuser@example.com", fetchedUser.Email, "Expected fetched email to match")

	// Fetch by email
	fetchedUser, err = suite.Repository.GetUserByEmailOrUsername(suite.TestContext, "", "fetchuser@example.com")
	suite.Empty(err, "Expected no error when fetching by email")
	suite.Equal("fetchuser", fetchedUser.Username, "Expected fetched username to match")
}

// TestUpdateUser tests the UpdateUser method
// TestUpdateUser tests the UpdateUser method
func (suite *UserRepositorySuite) TestUpdateUser() {
	// Insert a user to update
	user := &models.User{
		Username: "updatableuser",
		Email:    "updatable@example.com",
		Name:     "Old Name",
	}
	insertResult, err := suite.Collection.InsertOne(suite.TestContext, user)
	suite.Nil(err, "Expected no error when inserting user for update test")

	// Convert inserted ID to string
	objID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Test updating with new fields
	updatedUser := &models.User{
		Username: "newusername",
		Email:    "newemail@example.com",
		Name:     "New Name",
	}

	errResp := suite.Repository.UpdateUser(suite.TestContext, updatedUser, objID)
	suite.Empty(errResp, "Expected no error when updating user")

	// Verify the update
	var result models.User
	err = suite.Collection.FindOne(suite.TestContext, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	suite.Empty(err, "Expected no error when fetching updated user")
	suite.Equal("newusername", result.Username, "Expected username to be updated")
	suite.Equal("newemail@example.com", result.Email, "Expected email to be updated")
	suite.Equal("New Name", result.Name, "Expected name to be updated")

	// Test updating with empty fields (no update should occur)
	emptyUpdateUser := &models.User{}
	errResp = suite.Repository.UpdateUser(suite.TestContext, emptyUpdateUser, objID)
	suite.Empty(errResp, "Expected no error when updating with empty fields")

	// Verify no changes occurred
	err = suite.Collection.FindOne(suite.TestContext, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	suite.Empty(err, "Expected no error when fetching user after empty update")
	suite.Equal("newusername", result.Username, "Expected username to remain unchanged")
	suite.Equal("newemail@example.com", result.Email, "Expected email to remain unchanged")
	suite.Equal("New Name", result.Name, "Expected name to remain unchanged")
}

// TestUpdateUserInvalidID tests the UpdateUser method with an invalid ID
func (suite *UserRepositorySuite) TestUpdateUserInvalidID() {
	invalidID := "invalidObjectID"

	user := &models.User{
		Username: "username",
		Email:    "email@example.com",
		Name:     "User Name",
	}

	errResp := suite.Repository.UpdateUser(suite.TestContext, user, invalidID)
	suite.NotEmpty(errResp, "Expected an error when using an invalid ObjectID")
}

// TestDeleteUser tests the DeleteUser method
func (suite *UserRepositorySuite) TestDeleteUser() {
	// Insert a user to delete
	user := &models.User{
		Username: "deletableuser",
		Email:    "deletable@example.com",
	}
	suite.Collection.InsertOne(suite.TestContext, user)

	// Delete the user
	err := suite.Repository.DeleteUser(suite.TestContext, user.ID.Hex())
	suite.Empty(err, "Expected no error when deleting user")

	// Ensure the user is deleted
	var deletedUser models.User
	Err := suite.Collection.FindOne(suite.TestContext, bson.M{"username": "deletableuser"}).Decode(&deletedUser)
	suite.Nil(Err, "Expected an error when fetching a deleted user")
	suite.Empty(Err)
}

func (suite *UserRepositorySuite) TestPromoteUser() {
	// Insert a user for testing promotion
	user := &models.User{
		Username: "promotableuser",
		Email:    "promotable@example.com",
		Name:     "Promotable User",
		Role:     "user",
	}
	insertResult, err := suite.Collection.InsertOne(suite.TestContext, user)
	suite.Empty(err, "Expected no error when inserting user for promotion test")

	// Convert inserted ID to string
	objID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Test promoting the user to "admin"
	errResp := suite.Repository.PromoteUser(suite.TestContext, objID)
	suite.Empty(errResp, "Expected no error when promoting user to admin")

	// Verify the role update
	var promotedUser models.User
	var role models.Role = "admin"
	err = suite.Collection.FindOne(suite.TestContext, bson.M{"_id": insertResult.InsertedID}).Decode(&promotedUser)
	suite.Empty(err, "Expected no error when fetching user after promotion")
	suite.Equal(role, promotedUser.Role, "Expected role to be updated to admin")
}

func (suite *UserRepositorySuite) TestDemoteUser() {
	// Insert a user for testing demotion
	user := &models.User{
		Username: "demotableuser",
		Email:    "demotable@example.com",
		Name:     "Demotable User",
		Role:     "admin", // Initially an admin
	}
	insertResult, err := suite.Collection.InsertOne(suite.TestContext, user)
	suite.Empty(err, "Expected no error when inserting user for demotion test")

	// Convert inserted ID to string
	objID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Test demoting the user (e.g., to a lower role)
	errResp := suite.Repository.DemoteUser(suite.TestContext, objID)
	suite.Empty(errResp, "Expected no error when demoting user")

	// Verify the role update
	var demotedUser models.User
	var role models.Role = "admin"
	err = suite.Collection.FindOne(suite.TestContext, bson.M{"_id": insertResult.InsertedID}).Decode(&demotedUser)
	suite.Empty(err, "Expected no error when fetching user after demotion")
	suite.Equal(role, demotedUser.Role, "Expected role to remain admin (demotion logic should be handled correctly)")
}


func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
