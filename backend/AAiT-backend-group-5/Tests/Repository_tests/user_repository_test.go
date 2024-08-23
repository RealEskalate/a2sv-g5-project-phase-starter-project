package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	repository "github.com/aait.backend.g5.main/backend/Repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuite struct {
	suite.Suite

	db         *mocks.Database
	collection *mocks.Collection
	repository *repository.UserMongoRepository
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.db = new(mocks.Database)
	suite.collection = new(mocks.Collection)
	suite.repository = &repository.UserMongoRepository{
		Collection: suite.collection,
	}
}

func (suite *UserRepositorySuite) TestCreateUser_Success() {
	// Prepare input data
	user := &models.User{
		Username: "test-username",
		Email:    "email@email.com",
	}

	// Mock the InsertOne operation
	mockInsertResult := &mongo.InsertOneResult{InsertedID: "some-id"} // Simulating the InsertOne result
	suite.collection.On("InsertOne", mock.Anything, user).Return(mockInsertResult, nil)

	// Call the CreateUser method
	errResp := suite.repository.CreateUser(context.TODO(), user)

	// Assertions
	suite.Nil(errResp) // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestCreateUser_Failure() {
	// Prepare input data
	user := &models.User{
		Username: "test-username",
		Email:    "email@email.com",
	}

	// Mock the InsertOne operation to return an error
	mockErr := errors.New("failed to insert document")
	suite.collection.On("InsertOne", mock.Anything, user).Return(nil, mockErr)

	// Call the CreateUser method
	errResp := suite.repository.CreateUser(context.TODO(), user)

	// Assertions
	suite.NotNil(errResp)                                 // Expecting an error
	suite.Equal("failed to insert document", errResp.Message) // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetUserByEmailOrUsername_Success() {
	// Prepare input data
	username := "test-username"
	email := "email@email.com"
	user := &models.User{
		Username: username,
		Email:    email,
	}

	// Create a mock SingleResult
	mockSingleResult := mocks.NewSingleResult(suite.T())
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		*(args[0].(*models.User)) = *user
	}).Return(nil)

	// Mock the FindOne operation
	suite.collection.On("FindOne", mock.Anything, bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}).Return(mockSingleResult, nil)

	// Call the GetUserByEmailOrUsername method
	result, errResp := suite.repository.GetUserByEmailOrUsername(context.TODO(), username, email)

	// Assertions
	suite.NotNil(result)                                    // Expecting a user object
	suite.Equal(user, result)                              // Check if the user object matches
	suite.Nil(errResp)                                     // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetUserByEmailOrUsername_NotFound() {
	// Prepare input data
	username := "test-username"
	email := "email@email.com"
	user := &models.User{
	}

	// Create a mock SingleResult
	mockSingleResult := mocks.NewSingleResult(suite.T())
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		*(args[0].(*models.User)) = *user
	}).Return(mongo.ErrNoDocuments)

	// Mock the FindOne operation to return an error (not found)
	suite.collection.On("FindOne", mock.Anything, bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}).Return(mockSingleResult, mongo.ErrNoDocuments)

	// Call the GetUserByEmailOrUsername method
	result, errResp := suite.repository.GetUserByEmailOrUsername(context.TODO(), username, email)

	// Assertions
	suite.Empty(result)               
	suite.NotEmpty(*errResp)                     // 
	suite.Equal("user not found", errResp.Message)      // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetUserByName_Success() {
	// Prepare input data
	name := "test-name"
	user := &models.User{
		Name: name,
	}

	// Create a mock SingleResult
	mockSingleResult := mocks.NewSingleResult(suite.T())
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		*(args[0].(*models.User)) = *user
	}).Return(nil)

	// Mock the FindOne operation
	suite.collection.On("FindOne", mock.Anything, bson.M{"name": name}).Return(mockSingleResult, nil)

	// Call the GetUserByName method
	result, errResp := suite.repository.GetUserByName(context.TODO(), name)

	// Assertions
	suite.NotNil(result)                                    // Expecting a user object
	suite.Equal(user, result)                              // Check if the user object matches
	suite.Nil(errResp)                                     // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetUserByName_NotFound() {
	// Prepare input data
	name := "test-name"
	user := &models.User{
		Name: "wrong-name",
	}

	// Create a mock SingleResult
	mockSingleResult := mocks.NewSingleResult(suite.T())
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		*(args[0].(*models.User)) = *user
	}).Return(mongo.ErrNoDocuments)
	// Mock the FindOne operation to return an error (not found)
	suite.collection.On("FindOne", mock.Anything, bson.M{"name": name}).Return(mockSingleResult, mongo.ErrNoDocuments)

	// Call the GetUserByName method
	result, errResp := suite.repository.GetUserByName(context.TODO(), name)

	// Assertions
	suite.Nil(result)                                    // Expecting no user object
	suite.NotNil(errResp)                               // Expecting an error
	suite.Equal("user not found", errResp.Message)      // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}


func (suite *UserRepositorySuite) TestGetUserByID_Success() {
	// Prepare input data
	id := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(id)
	user := &models.User{
		ID:       id,
		Username: "test-username",
	}

	// Create a mock SingleResult
	mockSingleResult := mocks.NewSingleResult(suite.T())
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		*(args[0].(*models.User)) = *user
	}).Return(nil)

	// Mock the FindOne operation
	suite.collection.On("FindOne", mock.Anything, bson.M{"_id": objID}).Return(mockSingleResult, nil)

	// Call the GetUserByID method
	result, errResp := suite.repository.GetUserByID(context.TODO(), id)

	// Assertions
	suite.NotNil(result)                                     // Expecting a user object
	suite.Equal(user.ID, result.ID)                         // Check if the user ID matches
	suite.Nil(errResp)                                      // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetUserByID_Failure() {
	// Prepare input data
	id := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(id)
	user := &models.User{
		ID:       id,
		Username: "test-username",
	}

	// Create a mock SingleResult
	mockSingleResult := mocks.NewSingleResult(suite.T())
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		*(args[0].(*models.User)) = *user
	}).Return(mongo.ErrNoDocuments)
	// Mock the FindOne operation to return an error (not found)
	suite.collection.On("FindOne", mock.Anything, bson.M{"_id": objID}).Return(mockSingleResult, mongo.ErrNoDocuments)

	// Call the GetUserByID method
	result, errResp := suite.repository.GetUserByID(context.TODO(), id)

	// Assertions
	suite.Nil(result)                                     // Expecting no user object
	suite.NotNil(errResp)                                // Expecting an error
	suite.Equal("user with the given ID not found", errResp.Message) // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestUpdateUser_Success() {
	// Prepare input data
	id := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(id)
	user := &models.User{
		Username: "new-username",
		Name: "new-name",
		Bio: "new-bio",
		Email: "new-email",
		Password: "new-password",
		ImageKey: "new-image-key",
		PhoneNumber: "new-phone-number",
	}

	change := bson.M{
		"username": "new-username",
		"name": "new-name",
		"bio": "new-bio",
		"email": "new-email",
		"password": "new-password",
		"image_key": "new-image-key",
		"phone_number": "new-phone-number",
	}
	// Mock the UpdateOne operation
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"_id": objID}, bson.M{"$set": change}).Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil)

	// Call the UpdateUser method
	errResp := suite.repository.UpdateUser(context.TODO(), user, id)

	// Assertions
	suite.Nil(errResp) // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestUpdateUser_Failure() {
	// Prepare input data
	id := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(id)
	user := &models.User{
		Username: "new-username",
	}

	// Mock the UpdateOne operation to return an error
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"_id": objID}, bson.M{"$set": bson.M{
		"username": "new-username",
	}}).Return(nil, errors.New("update error"))

	// Call the UpdateUser method
	errResp := suite.repository.UpdateUser(context.TODO(), user, id)

	// Assertions
	suite.NotNil(errResp)                          // Expecting an error
	suite.Equal("update error", errResp.Message)  // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestDeleteUser_Success() {
	// Prepare input data
	id := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(id)

	// Mock the DeleteOne operation
	suite.collection.On("DeleteOne", mock.Anything, bson.M{"_id": objID}).Return(int64(1), nil)

	// Call the DeleteUser method
	errResp := suite.repository.DeleteUser(context.TODO(), id)

	// Assertions
	suite.Nil(errResp) // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestDeleteUser_Failure() {
	// Prepare input data
	id := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(id)

	// Mock the DeleteOne operation to return an error
	suite.collection.On("DeleteOne", mock.Anything, bson.M{"_id": objID}).Return(int64(0), errors.New("delete error"))

	// Call the DeleteUser method
	errResp := suite.repository.DeleteUser(context.TODO(), id)

	// Assertions
	suite.NotNil(errResp)                         // Expecting an error
	suite.Equal("delete error", errResp.Message) // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}



func (suite *UserRepositorySuite) TestPromoteUser_Success() {
	// Prepare input data
	userID := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(userID)

	// Mock the UpdateOne operation
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"_id": objID}, bson.M{"$set": bson.M{"role": "admin"}}).Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil)

	// Call the PromoteUser method
	errResp := suite.repository.PromoteUser(context.TODO(), userID)

	// Assertions
	suite.Nil(errResp) // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestPromoteUser_Failure() {
	// Prepare input data
	userID := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(userID)

	// Mock the UpdateOne operation to return an error
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"_id": objID}, bson.M{"$set": bson.M{"role": "admin"}}).Return(nil, errors.New("update error"))

	// Call the PromoteUser method
	errResp := suite.repository.PromoteUser(context.TODO(), userID)

	// Assertions
	suite.NotNil(errResp)                          // Expecting an error
	suite.Equal("update error", errResp.Message)  // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestDemoteUser_Success() {
	// Prepare input data
	userID := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(userID)

	// Mock the UpdateOne operation
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"_id": objID}, bson.M{"$set": bson.M{"role": "user"}}).Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil)

	// Call the DemoteUser method
	errResp := suite.repository.DemoteUser(context.TODO(), userID)

	// Assertions
	suite.Nil(errResp) // Expecting no error
	suite.collection.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestDemoteUser_Failure() {
	// Prepare input data
	userID := "605c72ef8f4c7a9f1d9a5c4f"
	objID, _ := primitive.ObjectIDFromHex(userID)

	// Mock the UpdateOne operation to return an error
	suite.collection.On("UpdateOne", mock.Anything, bson.M{"_id": objID}, bson.M{"$set": bson.M{"role": "user"}}).Return(nil, errors.New("update error"))

	// Call the DemoteUser method
	errResp := suite.repository.DemoteUser(context.TODO(), userID)

	// Assertions
	suite.NotNil(errResp)                          // Expecting an error
	suite.Equal("update error", errResp.Message)  // Check if the error message matches
	suite.collection.AssertExpectations(suite.T())
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
