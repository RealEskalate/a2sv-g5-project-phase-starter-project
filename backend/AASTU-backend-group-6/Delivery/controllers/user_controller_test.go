package controllers_test

import (
	"blogs/Delivery/controllers"
	domain "blogs/Domain"
	"blogs/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerTestSuite struct {
	suite.Suite
	mockUserUsecase *mocks.UserUseCase
	userController  *controllers.NewUserController
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUserUsecase = new(mocks.UserUseCase)
	suite.userController = &controllers.NewUserController{
		UserUsecase: suite.mockUserUsecase,
	}
}

func (suite *UserControllerTestSuite) TestUpdateUser_Success() {
	// Prepare a valid user update request
	updateUserRequest := domain.UserUpdateRequest{
		Full_Name:         "Updated Name",
		Username:          "updatedUsername",
		Password:          "newPassword123",
		Profile_image_url: "http://example.com/new-image.jpg",
		Contact:           "1234567890",
		Bio:               "Updated bio",
	}
	body, _ := json.Marshal(updateUserRequest)
	userID := "authenticatedUserId"

	// Create a request with the authenticated user's ID
	req, err := http.NewRequest(http.MethodPut, "/user/update/"+userID, bytes.NewReader(body))
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.AddParam("id", userID)
	c.Set("user_id", userID)

	// Mock the FindUserByID method to return an existing user
	existingUser := domain.User{
		ID:                primitive.NewObjectID(),
		Full_Name:         "Original Name",
		Username:          "originalUsername",
		Password:          "originalHashedPassword",
		Profile_image_url: "http://example.com/original-image.jpg",
		Contact:           "0987654321",
		Bio:               "Original bio",
	}
	suite.mockUserUsecase.On("FindUserByID", mock.Anything, userID).Return(&existingUser, nil).Once()

	// Mock the UpdateUser method to return the updated user
	updatedUser := existingUser
	updatedUser.Full_Name = updateUserRequest.Full_Name
	updatedUser.Username = updateUserRequest.Username
	updatedUser.Profile_image_url = updateUserRequest.Profile_image_url
	updatedUser.Contact = updateUserRequest.Contact
	updatedUser.Bio = updateUserRequest.Bio

	suite.mockUserUsecase.On("UpdateUser", mock.Anything, updateUserRequest).Return(&domain.SuccessResponse{
		Message: "User updated successfully",
		Data:    updatedUser,
		Status:  http.StatusOK,
	}).Once()

	// Call the UpdateUser method
	suite.userController.UpdateUser(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	var response domain.SuccessResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("User updated successfully", response.Message)
	suite.Equal(http.StatusOK, response.Status)

	// Convert response.Data to map[string]interface{} to verify fields
	dataMap, ok := response.Data.(map[string]interface{})
	if !ok {
		suite.T().Fatal("expected Data to be of type map[string]interface{}")
	}

	// Extract and verify user fields from the map
	fullName, _ := dataMap["full_name"].(string)
	username, _ := dataMap["username"].(string)
	profileImageURL, _ := dataMap["profile_image"].(string)
	contact, _ := dataMap["contact"].(string)
	bio, _ := dataMap["bio"].(string)

	// Create the decoded user for comparison
	decodedUser := domain.User{
		Full_Name:         fullName,
		Username:          username,
		Profile_image_url: profileImageURL,
		Contact:           contact,
		Bio:               bio,
	}

	// Assert that the updated user details match
	suite.Equal(updatedUser.Full_Name, decodedUser.Full_Name)
	suite.Equal(updatedUser.Username, decodedUser.Username)
	suite.Equal(updatedUser.Profile_image_url, decodedUser.Profile_image_url)
	suite.Equal(updatedUser.Contact, decodedUser.Contact)
	suite.Equal(updatedUser.Bio, decodedUser.Bio)

	// Assert that the mock was called with the correct parameters
	suite.mockUserUsecase.AssertCalled(suite.T(), "UpdateUser", mock.Anything, updateUserRequest)
}

func (suite *UserControllerTestSuite) TestUpdateUser_Unauthorized() {
	// Prepare a valid user update request
	updateUserRequest := domain.UserUpdateRequest{
		Full_Name:         "Unauthorized Name",
		Username:          "unauthorizedUsername",
		Password:          "unauthorizedPassword123",
		Profile_image_url: "http://example.com/unauthorized-image.jpg",
		Contact:           "0987654321",
		Bio:               "Unauthorized bio",
	}
	body, _ := json.Marshal(updateUserRequest)
	userID := "targetUserId"
	authenticatedUserID := "authenticatedUserId"

	// Create a request with a different user ID than the authenticated user's ID
	req, err := http.NewRequest(http.MethodPut, "/user/update/"+userID, bytes.NewReader(body))
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.AddParam("id", userID)
	c.Set("user_id", authenticatedUserID)

	// Call the UpdateUser method
	suite.userController.UpdateUser(c)

	// Verify the response
	suite.Equal(http.StatusUnauthorized, w.Code)
	var response map[string]string
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("Unauthorized to update user Information", response["error"])

	// Assert that no update was attempted on the usecase
	suite.mockUserUsecase.AssertNotCalled(suite.T(), "UpdateUser", mock.Anything, updateUserRequest)
}

func (suite *UserControllerTestSuite) TestPromoteUser_Success() {
	// Mock request and context
	role := "admin"
	userID := "someUserId"
	promotionRequest := domain.UserPromotionRequest{
		Action: "promote",
	}
	body, _ := json.Marshal(promotionRequest)
	req, err := http.NewRequest(http.MethodPut, "/user/promote/"+userID, bytes.NewReader(body))
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.AddParam("id", userID)
	c.Set("role", role)

	// Mock the PromoteandDemoteUser method to return a successful response
	suite.mockUserUsecase.On("PromoteandDemoteUser", mock.Anything, userID, promotionRequest, role).Return(&domain.SuccessResponse{
		Message: "User promoted successfully",
		Status:  http.StatusOK,
	}).Once()

	// Call the PromoteUser method
	suite.userController.PromoteUser(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	var response domain.SuccessResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("User promoted successfully", response.Message)

	// Assert that the mock was called with the correct parameters
	suite.mockUserUsecase.AssertCalled(suite.T(), "PromoteandDemoteUser", mock.Anything, userID, promotionRequest, role)
}

func (suite *UserControllerTestSuite) TestPromoteUser_Unauthorized() {
	// Mock request and context
	role := "user" // Non-admin role
	userID := "someUserId"
	promotionRequest := domain.UserPromotionRequest{
		Action: "promote",
	}
	body, _ := json.Marshal(promotionRequest)
	req, err := http.NewRequest(http.MethodPut, "/user/promote/"+userID, bytes.NewReader(body))
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.AddParam("id", userID)
	c.Set("role", role)

	// Mock the PromoteandDemoteUser method to return an unauthorized response
	suite.mockUserUsecase.On("PromoteandDemoteUser", mock.Anything, userID, promotionRequest, role).Return(&domain.ErrorResponse{
		Message: "Unauthorized to promote/demote user",
		Status:  http.StatusForbidden,
	}).Once()

	// Call the PromoteUser method
	suite.userController.PromoteUser(c)

	// Verify the response
	suite.Equal(http.StatusForbidden, w.Code)
	var response domain.ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("Unauthorized to promote/demote user", response.Message)

	// Assert that the mock was called with the correct parameters
	suite.mockUserUsecase.AssertCalled(suite.T(), "PromoteandDemoteUser", mock.Anything, userID, promotionRequest, role)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
