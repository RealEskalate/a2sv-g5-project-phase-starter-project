package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"net/http"
	"net/http/httptest"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthControllerSuite struct {
	suite.Suite
	mockUsecase    *mocks.IAuthUsecase
	authController *controllers.AuthController
}

func (suite *AuthControllerSuite) SetupTest() {
	suite.mockUsecase = new(mocks.IAuthUsecase)
	googleConfig := &oauth2.Config{} // Mock or setup a minimal OAuth2 config if needed

	suite.authController = controllers.NewAuthController(suite.mockUsecase, googleConfig)
}

func (suite *AuthControllerSuite) TestRegister_Success() {
	userDTO := dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}
	user := &dto.CreatedResponseDto{
		ID:       uuid.New(),
		FullName: "John Doe",
		Email:    "john@example.com",
		Bio:      "",
		ImageUrl: "",
	}

	suite.mockUsecase.On("RegisterUser", &userDTO).Return(user, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(userDTO)
	c.Request = httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Register(c)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "John Doe")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestRegister_InvalidJson() {
	userDTO := dto.RegisterUserDTO{
		FullName: "",
		Email:    "invalid-email",
		Password: "short",
	}

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(userDTO)
	c.Request = httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Register(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "'FullName' failed on the 'required'")
	assert.Contains(suite.T(), w.Body.String(), "'Email' failed on the 'email'")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestRegister_EmailAlreadyExists() {
	userDTO := dto.RegisterUserDTO{
		FullName: "John Doe",
		Email:    "existing@email.com",
		Password: "password",
	}
	suite.mockUsecase.On("RegisterUser", &userDTO).Return(nil, domain.ErrUserEmailExists)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(userDTO)
	c.Request = httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Register(c)

	assert.Equal(suite.T(), domain.ErrUserEmailExists.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Email already exists")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestLogin_Success() {
	loginDTO := dto.LoginUserDTO{
		Email:    "john@example.com",
		Password: "password",
	}
	tokens := &dto.TokenResponseDTO{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
	}

	suite.mockUsecase.On("LoginUser", &loginDTO).Return(tokens, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(loginDTO)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Login(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "access-token")
	assert.Contains(suite.T(), w.Body.String(), "refresh-token")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestLogin_InvalidCredentials() {
	loginDTO := dto.LoginUserDTO{
		Email:    "invalid@example.com",
		Password: "wrongpassword",
	}

	suite.mockUsecase.On("LoginUser", &loginDTO).Return(nil, domain.ErrInvalidCredentials)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(loginDTO)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Login(c)

	assert.Equal(suite.T(), domain.ErrInvalidCredentials.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid email or password")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestLogin_UserNotFound() {
	loginDTO := dto.LoginUserDTO{
		Email:    "john@example.com",
		Password: "password",
	}

	suite.mockUsecase.On("LoginUser", &loginDTO).Return(nil, domain.ErrUserNotFound)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(loginDTO)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Login(c)

	assert.Equal(suite.T(), domain.ErrUserNotFound.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "User not found")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestLogin_InvalidJson() {
	loginDTO := dto.LoginUserDTO{
		Email:    "",
		Password: "",
	}

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(loginDTO)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.Login(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "'Email' failed on the 'required'")
	assert.Contains(suite.T(), w.Body.String(), "'Password' failed on the 'required'")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestRefreshToken_Success() {
	refreshDTO := dto.RefreshTokenDTO{
		RefreshToken: "valid-refresh-token",
	}
	tokens := &dto.TokenResponseDTO{
		AccessToken:  "new-access-token",
		RefreshToken: "new-refresh-token",
	}

	suite.mockUsecase.On("RefreshTokens", "valid-refresh-token").Return(tokens, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(refreshDTO)
	c.Request = httptest.NewRequest("POST", "/refresh-token", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.RefreshToken(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "new-access-token")
	assert.Contains(suite.T(), w.Body.String(), "new-refresh-token")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestRefreshToken_InvalidRefreshToken() {
	refreshDTO := dto.RefreshTokenDTO{
		RefreshToken: "invalid-refresh-token",
	}

	suite.mockUsecase.On("RefreshTokens", "invalid-refresh-token").Return(nil, domain.ErrInvalidToken)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(refreshDTO)
	c.Request = httptest.NewRequest("POST", "/refresh-token", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.RefreshToken(c)

	assert.Equal(suite.T(), domain.ErrInvalidToken.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid token")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestRefreshToken_InvalidJson() {
	refreshDTO := dto.RefreshTokenDTO{
		RefreshToken: "",
	}

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(refreshDTO)
	c.Request = httptest.NewRequest("POST", "/refresh-token", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.RefreshToken(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "'RefreshToken' failed on the 'required'")
	
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestForgotPassword_Success() {
	forgotPasswordDTO := dto.ForgotPasswordRequestDTO{
		Email: "john@example.com",
	}

	suite.mockUsecase.On("ForgotPassword", &forgotPasswordDTO).Return(nil)
	
	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(forgotPasswordDTO)
	c.Request = httptest.NewRequest("POST", "/forgot-password", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.ForgotPassword(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Password reset link sent to your email")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestForgotPassword_InvalidJson() {
	forgotPasswordDTO := dto.ForgotPasswordRequestDTO{
		Email: "",
	}

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(forgotPasswordDTO)
	c.Request = httptest.NewRequest("POST", "/forgot-password", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.ForgotPassword(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "'Email' failed on the 'required'")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestResetPassword_Success() {
	resetPasswordDTO := dto.ResetPasswordRequestDTO{
		Token: "valid-token",
		NewPassword: "new-password",
	}

	suite.mockUsecase.On("ResetPassword", &resetPasswordDTO).Return(nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(resetPasswordDTO)
	c.Request = httptest.NewRequest("POST", "/reset-password", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.ResetPassword(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Password reset successfully")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerSuite) TestResetPassword_InvalidJson() {
	resetPasswordDTO := dto.ResetPasswordRequestDTO{
		Token: "",
		NewPassword: "",
	}

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(resetPasswordDTO)
	c.Request = httptest.NewRequest("POST", "/reset-password", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.authController.ResetPassword(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "'Token' failed on the 'required'")
	assert.Contains(suite.T(), w.Body.String(), "'NewPassword' failed on the 'required'")

	suite.mockUsecase.AssertExpectations(suite.T())
}

func TestAuthControllerSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerSuite))
}
