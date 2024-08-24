package tests

import (
	"blog_api/delivery/controllers"
	"blog_api/domain"
	"blog_api/mocks"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	authController  controllers.AuthController
	mockUserUsecase *mocks.UserUsecaseInterface
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.mockUserUsecase = new(mocks.UserUsecaseInterface)
	suite.authController = *controllers.NewAuthController(suite.mockUserUsecase, nil)
}

func (suite *AuthControllerTestSuite) TestLogin() {
	// Prepare the request body directly
	requestBody := `{"username":"testuser", "password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up expected response
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNjkxMjM5MjAwfQ.sK5NzGzZMjcM_E0Kni3z6IZxwslaz1w0bQetfKBvQ1M"
	expectedRefreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNjk1MjM1NjAwfQ.XGFLg-5eTgDl_2LrkDRdG0z3rVZOXH3XkOZk1E6QjOQ"

	suite.mockUserUsecase.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).
		Return(expectedToken, expectedRefreshToken, nil)

	// Set up the Gin context and recorder
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Call the handler
	suite.authController.HandleLogin(ctx)

	// Assert the response
	assert.Equal(suite.T(), http.StatusCreated, w.Code) // Update this line
	assert.Contains(suite.T(), w.Body.String(), expectedToken)
	assert.Contains(suite.T(), w.Body.String(), expectedRefreshToken)
	suite.mockUserUsecase.AssertExpectations(suite.T())
}
func (suite *AuthControllerTestSuite) TestSignup() {
	// Prepare the request body directly
	requestBody := `{"username":"newuser", "email":"newuser@example.com", "password":"password123"}`
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up the expected behavior for the mock use case
	suite.mockUserUsecase.On("Signup", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("string")).
		Return(nil)

	// Set up the Gin context and recorder
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Call the handler
	suite.authController.HandleSignup(ctx)

	// Assert the response
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockUserUsecase.AssertExpectations(suite.T())
}
func (suite *AuthControllerTestSuite) TestHandleRenewAccessToken() {
	// Setup mock
	refreshToken := "refresh-token"
	newAccessToken := "new-access-token"
	suite.mockUserUsecase.On("RenewAccessToken", mock.Anything, refreshToken).Return(newAccessToken, nil)

	// Create request
	req, _ := http.NewRequest(http.MethodPost, "/renew-token", nil)
	req.Header.Set("Authorization", "Bearer "+refreshToken)
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call handler
	suite.authController.HandleRenewAccessToken(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"accessToken":"new-access-token"}`, w.Body.String())
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestHandleInitResetPassword() {
	// Setup mock
	user := domain.User{
		Username: "user123",
		Email:    "user123@example.com",
	}
	suite.mockUserUsecase.On("InitResetPassword", mock.Anything, user.Username, user.Email, mock.Anything).Return(nil)

	// Create request
	resetData := `{"username":"user123","email":"user123@example.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/init-reset-password", bytes.NewBufferString(resetData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call handler
	suite.authController.HandleInitResetPassword(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"A reset password token has been sent to your email."}`, w.Body.String())
	suite.mockUserUsecase.AssertExpectations(suite.T())
}
func (suite *AuthControllerTestSuite) TestHandleLogout() {
	// Setup mock
	username := "user123"
	token := "logout-token"
	suite.mockUserUsecase.On("Logout", mock.Anything, username, token).Return(nil)

	// Create request
	req, _ := http.NewRequest(http.MethodPost, "/logout", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Keys = map[string]interface{}{"username": username}

	// Call handler
	suite.authController.HandleLogout(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestHandleDemoteUser() {
	// Setup mock
	username := "user123"
	suite.mockUserUsecase.On("DemoteUser", mock.Anything, username).Return(nil)

	// Create request
	req, _ := http.NewRequest(http.MethodPost, "/demote/user123", nil)
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "username", Value: username},
	}

	// Call handler
	suite.authController.HandleDemoteUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"User demoted"}`, w.Body.String())
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestHandlePromoteUser() {
	// Setup mock
	username := "user123"
	suite.mockUserUsecase.On("PromoteUser", mock.Anything, username).Return(nil)

	// Create request
	req, _ := http.NewRequest(http.MethodPost, "/promote/user123", nil)
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "username", Value: username},
	}

	// Call handler
	suite.authController.HandlePromoteUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"User promoted"}`, w.Body.String())
	suite.mockUserUsecase.AssertExpectations(suite.T())
}
func (suite *AuthControllerTestSuite) TestHandleVerifyEmail() {
	// Setup mock
	username := "user123"
	token := "verify-token"
	suite.mockUserUsecase.On("VerifyEmail", mock.Anything, username, token, mock.Anything).Return(nil)

	// Create request
	req, _ := http.NewRequest(http.MethodPost, "/verify/user123/verify-token", nil)
	req.Header.Set("Authorization", "Bearer token") // If your handler requires an Authorization header
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "username", Value: username},
		{Key: "token", Value: token},
	}

	// Call handler
	suite.authController.HandleVerifyEmail(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"User verified"}`, w.Body.String())
	suite.mockUserUsecase.AssertExpectations(suite.T())
}
func (suite *AuthControllerTestSuite) TestHandleUpdateUser() {
	// Setup mock
	reqUsername := "user123"
	tokenUsername := "authUser"
	mockResponse := map[string]string{"bio": "UpdatedBio", "phone_number": "1234567890"}
	suite.mockUserUsecase.On("UpdateUser", mock.Anything, reqUsername, tokenUsername, mock.Anything).Return(mockResponse, nil)
	suite.authController.DeleteFile = func(filePath string) error { return nil }

	// Create request with form data
	form := `bio=UpdatedBio&phone_number=1234567890`
	req, _ := http.NewRequest(http.MethodPost, "/update/user123", strings.NewReader(form))
	req.Header.Set("Authorization", "Bearer token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	// Create context
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "username", Value: reqUsername},
	}
	c.Keys = map[string]interface{}{"username": tokenUsername}

	// Call handler
	suite.authController.HandleUpdateUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	expectedResponse := `{"message":"User updated","data":{"bio":"UpdatedBio","phone_number":"1234567890"}}`
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
