package tests

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/domain/dtos"
	initdb "blog_api/infrastructure/db"
	"blog_api/repository"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockUserData is a slice of mock user data used through out the tests
// The first three users are normal users, the next three are admins, and the last one is the root user
var MockUserData = []domain.User{
	{Username: "testuser1", Email: "testuser1@gmail.com", Password: "password", Role: domain.RoleUser},
	{Username: "testuser2", Email: "testuser2@gmail.com", Password: "password", Role: domain.RoleUser},
	{Username: "testuser3", Email: "testuser3@gmail.com", Password: "password", Role: domain.RoleUser},
	{Username: "testuser4", Email: "testuser4@gmail.com", Password: "password", Role: domain.RoleAdmin},
	{Username: "testuser5", Email: "testuser5@gmail.com", Password: "password", Role: domain.RoleAdmin},
	{Username: "testuser6", Email: "testuser6@gmail.com", Password: "password", Role: domain.RoleAdmin},
	{Username: "rootUser", Email: "testuser7@gmail.com", Password: "password", Role: domain.RoleRoot},
}

// UserRepositoryTestSuite is a test suite for the UserRepository
type UserRepositoryTestSuite struct {
	suite.Suite
	client         *mongo.Client
	collection     *mongo.Collection
	UserRepository *repository.UserRepository
}

// SetupSuite initializes the database connection and the UserRepository
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

// Delete all documents from the collection before each test
func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.collection.DeleteMany(context.Background(), bson.D{})
}

// TestCreateUser_Positive tests the CreateUser method with a new user
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

// TestCreateUser_Negative_DuplicateEmail tests the CreateUser method with a user that has a duplicate email
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

// TestCreateUser_Negative_DuplicateUsername tests the CreateUser method with a user that has a duplicate username
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

// TestFindUser_Positive tests the FindUser method with a user that exists
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

// TestFindUser_Negative_UserNotFound tests the FindUser method with a user that does not exist
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

// TestSetRefreshToken_Positive tests the SetRefreshToken method with a valid user
func (suite *UserRepositoryTestSuite) TestSetRefreshToken_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)

	newRefreshToken := "this is a. kinda valid refresh token. it has the two dots"
	err := suite.UserRepository.SetRefreshToken(context.Background(), &user, newRefreshToken)
	suite.Nil(err, "no error when setting refresh token")
}

// TestSetRefreshToken_Negative tests the SetRefreshToken method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestSetRefreshToken_Negative_UsetNotFound() {
	user := MockUserData[0]
	// user not created

	newRefreshToken := "this is a. kinda valid refresh token. it has the two dots"
	err := suite.UserRepository.SetRefreshToken(context.Background(), &user, newRefreshToken)
	suite.NotNil(err, "no error when setting refresh token")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TestUpdateUser_Positive_NonlocalProfilePicture tests the UpdateUser method with a user that has a non-local profile picture
func (suite *UserRepositoryTestSuite) TestUpdateUser_Positive_NonlocalProfilePicture() {
	user := MockUserData[0]
	originalImage := "oldfile.jpg"
	user.ProfilePicture.IsLocal = false
	user.ProfilePicture.FileName = originalImage
	suite.UserRepository.CreateUser(context.Background(), &user)

	updates := dtos.UpdateUser{
		PhoneNumber: "2511234567890",
		Bio:         "new bio",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "newfile.jpg",
			IsLocal:  true,
		},
	}

	// check with the updatedData map
	updatedData, oldFile, err := suite.UserRepository.UpdateUser(context.Background(), user.Username, &updates)
	suite.Nil(err, "no error when updating user")
	suite.Equal(updates.PhoneNumber, updatedData["phonenumber"], "phone number matches")
	suite.Equal(updates.Bio, updatedData["bio"], "bio matches")
	suite.Equal(updates.ProfilePicture.FileName, updatedData["profilepicture"], "profile picture matches")
	suite.Equal(updates.ProfilePicture.IsLocal, true, "profile picture is local")
	suite.Equal("", oldFile, "old file name matches")

	// check with the user in the DB
	var updatedDataFromDB domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&updatedDataFromDB)
	suite.Equal(updates.PhoneNumber, updatedDataFromDB.PhoneNumber, "phone number matches")
	suite.Equal(updates.Bio, updatedDataFromDB.Bio, "bio matches")
	suite.Equal(updates.ProfilePicture.FileName, updatedDataFromDB.ProfilePicture.FileName, "profile picture matches")
	suite.Equal(updates.ProfilePicture.IsLocal, updatedDataFromDB.ProfilePicture.IsLocal, "profile picture is local")
}

// TestUpdateUser_Positive_LocalProfilePicture tests the UpdateUser method with a user that has a local profile picture
func (suite *UserRepositoryTestSuite) TestUpdateUser_Positive_LocalProfilePicture() {
	user := MockUserData[0]
	originalImage := "oldfile.jpg"
	user.ProfilePicture.IsLocal = true
	user.ProfilePicture.FileName = originalImage
	suite.UserRepository.CreateUser(context.Background(), &user)

	updates := dtos.UpdateUser{
		PhoneNumber: "2511234567890",
		Bio:         "new bio",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "newfile.jpg",
			IsLocal:  true,
		},
	}

	// check with the updatedData map
	updatedData, oldFile, err := suite.UserRepository.UpdateUser(context.Background(), user.Username, &updates)
	suite.Nil(err, "no error when updating user")
	suite.Equal(updates.PhoneNumber, updatedData["phonenumber"], "phone number matches")
	suite.Equal(updates.Bio, updatedData["bio"], "bio matches")
	suite.Equal(updates.ProfilePicture.FileName, updatedData["profilepicture"], "profile picture matches")
	suite.Equal(updates.ProfilePicture.IsLocal, true, "profile picture is local")
	suite.Equal(originalImage, oldFile, "old file name matches")

	// check with the user in the DB
	var updatedDataFromDB domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&updatedDataFromDB)
	suite.Equal(updates.PhoneNumber, updatedDataFromDB.PhoneNumber, "phone number matches")
	suite.Equal(updates.Bio, updatedDataFromDB.Bio, "bio matches")
	suite.Equal(updates.ProfilePicture.FileName, updatedDataFromDB.ProfilePicture.FileName, "profile picture matches")
	suite.Equal(updates.ProfilePicture.IsLocal, updatedDataFromDB.ProfilePicture.IsLocal, "profile picture is local")
}

// TestUpdateUser_Negative_UserNotFound tests the UpdateUser method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestUpdateUser_Negative_UserNotFound() {
	username := "this one doesnt exist"
	updates := dtos.UpdateUser{
		PhoneNumber: "2511234567890",
		Bio:         "new bio",
		ProfilePicture: dtos.ProfilePicture{
			FileName: "newfile.jpg",
			IsLocal:  true,
		},
	}

	// check with the updatedData map
	updatedData, oldFile, err := suite.UserRepository.UpdateUser(context.Background(), username, &updates)
	suite.Equal("", oldFile, "old file name matches")
	suite.Equal(len(updatedData), 0, "no data updated")
	suite.NotNil(err, "error when updating user")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TestChangeRole_Positive tests the ChangeRole method with a valid user
func (suite *UserRepositoryTestSuite) TestChangeRole_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)

	newRole := "custom_rolies"
	err := suite.UserRepository.ChangeRole(context.Background(), user.Username, newRole)
	suite.Nil(err, "no error when changing role")

	// check with the user in the DB
	var updatedDataFromDB domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&updatedDataFromDB)
	suite.Equal(newRole, updatedDataFromDB.Role, "role matches")
	suite.Equal(user.Username, updatedDataFromDB.Username, "username matches")
	suite.Equal(user.Email, updatedDataFromDB.Email, "email matches")
	suite.Equal(user.Password, updatedDataFromDB.Password, "password matches")
}

// TestChangeRole_Negative_CantChangeRoot tests the ChangeRole method with the root user
func (suite *UserRepositoryTestSuite) TestChangeRole_Negative_CantChangeRoot() {
	user := MockUserData[len(MockUserData)-1]
	suite.UserRepository.CreateUser(context.Background(), &user)

	newRole := "custom_rolies"
	err := suite.UserRepository.ChangeRole(context.Background(), user.Username, newRole)
	suite.NotNil(err, "error when changing role")

	// check with the user in the DB
	var updatedDataFromDB domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&updatedDataFromDB)
	suite.NotEqual(newRole, updatedDataFromDB.Role, "role not updated")
	suite.Equal(user.Role, updatedDataFromDB.Role, "old role unaffected")
	suite.Equal(user.Username, updatedDataFromDB.Username, "username matches")
	suite.Equal(user.Email, updatedDataFromDB.Email, "email matches")
	suite.Equal(user.Password, updatedDataFromDB.Password, "password matches")
}

// TestChangeRole_Negative_UserNotFound tests the ChangeRole method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestChangeRole_Negative_UserNotFound() {
	user := MockUserData[0]
	newRole := "custom_rolies"
	err := suite.UserRepository.ChangeRole(context.Background(), user.Username, newRole)
	suite.NotNil(err, "error when changing role")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TestUpdateVerificationDetails_Positive tests the UpdateVerificationDetails method with a valid user
func (suite *UserRepositoryTestSuite) TestUpdateVerificationDetails_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)

	verificationDetails := domain.VerificationData{
		Token:     "pretend this is a very long random string",
		ExpiresAt: time.Now().Round(time.Second),
		Type:      domain.VerifyEmailType,
	}
	err := suite.UserRepository.UpdateVerificationDetails(context.Background(), user.Username, verificationDetails)
	suite.Nil(err, "no error when updating verification details")

	// check with the user in the DB
	var updatedDataFromDB domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&updatedDataFromDB)
	suite.Equal(verificationDetails.Token, updatedDataFromDB.VerificationData.Token, "token matches")
	suite.Equal(verificationDetails.ExpiresAt, updatedDataFromDB.VerificationData.ExpiresAt.Local(), "expires at matches")
	suite.Equal(verificationDetails.Type, updatedDataFromDB.VerificationData.Type, "type matches")
	suite.Equal(user.Username, updatedDataFromDB.Username, "username matches")
	suite.Equal(user.Email, updatedDataFromDB.Email, "email matches")
	suite.Equal(user.Password, updatedDataFromDB.Password, "password matches")
}

// TestUpdateVerificationDetails_Negative_UserNotFound tests the UpdateVerificationDetails method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestUpdateVerificationDetails_Negative_UserNotFound() {
	user := MockUserData[0]

	verificationDetails := domain.VerificationData{
		Token:     "pretend this is a very long random string",
		ExpiresAt: time.Now().Round(time.Second),
		Type:      domain.VerifyEmailType,
	}
	err := suite.UserRepository.UpdateVerificationDetails(context.Background(), user.Username, verificationDetails)
	suite.NotNil(err, "no error when updating verification details")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TestVerifyUser_Positive tests the VerifyUser method with a valid user
func (suite *UserRepositoryTestSuite) TestVerifyUser_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)

	// check before verification
	var userBefore domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&userBefore)
	suite.False(userBefore.IsVerified, "user is not verified")

	err := suite.UserRepository.VerifyUser(context.Background(), user.Username)
	suite.Nil(err, "no error when verifying user")

	// check after verification
	var userAfter domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&userAfter)
	suite.True(userAfter.IsVerified, "user is verified")
}

// TestVerifyUser_Negative_UserNotFound tests the VerifyUser method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestVerifyUser_Negative_UserNotFound() {
	user := MockUserData[0]
	// suite.UserRepository.CreateUser(context.Background(), &user)

	err := suite.UserRepository.VerifyUser(context.Background(), user.Username)
	suite.NotNil(err, "error when verifying user")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TestUpdatePassword_Positive tests the UpdatePassword method with a valid user
func (suite *UserRepositoryTestSuite) TestUpdatePassword_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)
	newPassword := "newpassword121983"

	err := suite.UserRepository.UpdatePassword(context.Background(), user.Username, newPassword)
	suite.Nil(err, "no error when updating password")

	// check from the DB
	var userAfter domain.User
	suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&userAfter)
	suite.Equal(newPassword, userAfter.Password, "password matches")
}

// TestUpdatePassword_Negative_UserNotFound tests the UpdatePassword method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestUpdatePassword_Negative_UserNotFound() {
	user := MockUserData[0]
	newPassword := "newpassword121983"

	err := suite.UserRepository.UpdatePassword(context.Background(), user.Username, newPassword)
	suite.NotNil(err, "error when updating password")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TestDeleteUser_Positive tests the DeleteUser method with a valid user
func (suite *UserRepositoryTestSuite) TestDeleteUser_Positive() {
	user := MockUserData[0]
	suite.UserRepository.CreateUser(context.Background(), &user)

	err := suite.UserRepository.DeleteUser(context.Background(), user.Username)
	suite.Nil(err, "no error when deleting user")

	// check from the DB
	var userAfter domain.User
	qErr := suite.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&userAfter)
	suite.NotNil(qErr, "user not found")
	suite.Equal(qErr, mongo.ErrNoDocuments, "user not found")
}

// TestDeleteUser_Negative_UserNotFound tests the DeleteUser method with a user that does not exist
func (suite *UserRepositoryTestSuite) TestDeleteUser_Negative_UserNotFound() {
	user := MockUserData[0]

	err := suite.UserRepository.DeleteUser(context.Background(), user.Username)
	suite.NotNil(err, "error when deleting user")
	suite.Equal(err.GetCode(), domain.ERR_NOT_FOUND, "error code is not found")
}

// TeardownSuite disconnects the database connection
func (suite *UserRepositoryTestSuite) TeardownSuite() {
	initdb.DisconnectDB(suite.client)
}

func TestUserRepositry(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
