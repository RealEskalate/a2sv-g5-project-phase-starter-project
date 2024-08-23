package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	config "github.com/aait.backend.g5.main/backend/Config"
	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Test Suite
type LoginControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.LoginUsecase
	controller  *controllers.LoginController
	router      *gin.Engine
	env         *config.Env
}

func (suite *LoginControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.LoginUsecase)
	suite.env = &config.Env{}
	suite.controller = controllers.NewLoginController(suite.mockUsecase, suite.env)
	suite.router = gin.Default()

	// Define the route
	suite.router.POST("/login", suite.controller.Login)
}

func (suite *LoginControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *LoginControllerTestSuite) TestLogin_Success() {
	loginRequest := dtos.LoginRequest{
		UsernameOrEmail: "test_user",
		Password:        "test_password",
	}
	loginResponse := &dtos.LoginResponse{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
	}

	suite.mockUsecase.On("LoginUser", mock.Anything, loginRequest).Return(loginResponse, nil).Once()

	body, _ := json.Marshal(loginRequest)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "test_access_token")
	suite.Contains(responseWriter.Body.String(), "test_refresh_token")
}

func (suite *LoginControllerTestSuite) TestLogin_BadRequest() {
	request, _ := http.NewRequest(http.MethodPost, "/login", nil)
	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusBadRequest, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "missing form body")
}

func (suite *LoginControllerTestSuite) TestLogin_InternalServerError() {
	loginRequest := dtos.LoginRequest{
		UsernameOrEmail: "test_user",
		Password:        "test_password",
	}
	expectedError := &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}

	suite.mockUsecase.On("LoginUser", mock.Anything, loginRequest).Return(&dtos.LoginResponse{}, expectedError).Once()

	body, _ := json.Marshal(loginRequest)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "internal server error")
}

func TestLoginControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LoginControllerTestSuite))
}
