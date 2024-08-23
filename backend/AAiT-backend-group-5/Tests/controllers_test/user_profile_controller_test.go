package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
// 	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
// 	models "github.com/aait.backend.g5.main/backend/Domain/Models"
// 	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type UserProfileControllerTestSuite struct {
// 	suite.Suite
// 	mockUsecase *mocks.UserProfileUpdateUsecase
// 	controller  *controllers.UserProfileController
// 	router      *gin.Engine
// }

// func (suite *UserProfileControllerTestSuite) SetupSuite() {
// 	suite.mockUsecase = new(mocks.UserProfileUpdateUsecase)
// 	suite.controller = controllers.NewUserProfileController(suite.mockUsecase)
// 	suite.router = gin.Default()

// 	// Define the route
// 	suite.router.PUT("/profile/update", suite.controller.ProfileUpdate)
// }

// func (suite *UserProfileControllerTestSuite) TearDownSuite() {
// 	suite.mockUsecase.AssertExpectations(suite.T())
// }

// func (suite *UserProfileControllerTestSuite) TestProfileUpdate_Success() {
// 	// Prepare request data
// 	updatedUser := &dtos.ProfileUpdateRequest{
// 		Name:     "test_name",
// 		Username: "test_username",
// 		Password: "test_password",
// 	}

// 	jsonData, _ := json.Marshal(updatedUser)
// 	request, _ := http.NewRequest(http.MethodPut, "/profile/update", bytes.NewBuffer(jsonData))
// 	request.Header.Set("Content-Type", "application/json")

// 	// Use Gin's context and manually set the user ID
// 	responseWriter := httptest.NewRecorder()

// 	// Directly set the "id" in the context
// 	ctx, _ := gin.CreateTestContext(responseWriter)
// 	ctx.Request = request
// 	ctx.Set("id", "user-123")

// 	// Call the controller directly with the context
// 	suite.mockUsecase.On("UpdateUserProfile", mock.AnythingOfType("*gin.Context"), "user-123", updatedUser).Return(nil).Once()
// 	suite.controller.ProfileUpdate(ctx)

// 	suite.Equal(http.StatusOK, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "user profile successfully updated")
// }

// func (suite *UserProfileControllerTestSuite) TestProfileUpdate_BindError() {
// 	// Prepare invalid request data
// 	invalidData := "invalid request"
// 	request, _ := http.NewRequest(http.MethodPut, "/profile/update", bytes.NewBufferString(invalidData))
// 	request.Header.Set("Content-Type", "application/json")

// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusInternalServerError, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "invalid request")
// }

// func (suite *UserProfileControllerTestSuite) TestProfileUpdate_UpdateError() {
// 	// Prepare request data
// 	updatedUser := &dtos.ProfileUpdateRequest{
// 		Name:     "test_name",
// 		Username: "test_username",
// 		Password: "test_password",
// 	}

// 	jsonData, _ := json.Marshal(updatedUser)
// 	request, _ := http.NewRequest(http.MethodPut, "/profile/update", bytes.NewBuffer(jsonData))
// 	request.Header.Set("Content-Type", "application/json")

// 	// Use Gin's context and manually set the user ID
// 	responseWriter := httptest.NewRecorder()

// 	// Set user ID in context
// 	ctx, _ := gin.CreateTestContext(responseWriter)
// 	ctx.Request = request
// 	ctx.Set("id", "user-123")

// 	expectedError := &models.ErrorResponse{
// 		Code:    http.StatusInternalServerError,
// 		Message: "update failed",
// 	}

// 	suite.mockUsecase.On("UpdateUserProfile", mock.Anything, "user-123", updatedUser).Return(expectedError).Once()
// 	suite.controller.ProfileUpdate(ctx)

// 	suite.Equal(expectedError.Code, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "update failed")
// }

// func TestUserProfileControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserProfileControllerTestSuite))
// }
