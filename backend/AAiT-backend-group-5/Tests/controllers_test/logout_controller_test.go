package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Test Suite
type LogoutControllerTestSuite struct {
	suite.Suite
	mockUsecase    *mocks.LogoutUsecase
	mockJwtService *mocks.JwtService
	controller     *controllers.LogoutController
	router         *gin.Engine
}

func (suite *LogoutControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.LogoutUsecase)
	suite.mockJwtService = new(mocks.JwtService)
	suite.controller = controllers.NewLogoutController(suite.mockUsecase, suite.mockJwtService)
	suite.router = gin.Default()

	// Define the route
	suite.router.POST("/logout", suite.controller.Logout)
}

func (suite *LogoutControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *LogoutControllerTestSuite) TestLogout_Success() {
	suite.mockUsecase.On("LogoutUser", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

	request, _ := http.NewRequest(http.MethodPost, "/logout", nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "logged out successfully")
}

func (suite *LogoutControllerTestSuite) TestLogout_InternalServerError() {
	expectedError := &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}

	suite.mockUsecase.On("LogoutUser", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("string")).Return(expectedError).Once()

	request, _ := http.NewRequest(http.MethodPost, "/logout", nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "internal server error")
}

func TestUserLogoutControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutControllerTestSuite))
}
