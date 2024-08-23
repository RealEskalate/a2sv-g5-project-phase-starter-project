package controller_test

// import (
// 	"backend-starter-project/delivery/controller"
// 	"backend-starter-project/domain/dto"
// 	"backend-starter-project/domain/entities"
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type AuthControllerTestSuite struct {
// 	suite.Suite
// 	authController *controller.AuthController
// 	authServiceMock *MockAuthService
// 	passwordResetServiceMock *MockPasswordResetService
// 	router *gin.Engine
// }

// // MockAuthService is a mock implementation of the AuthenticationService interface
// type MockAuthService struct {
// 	mock.Mock
// }

// func (m *MockAuthService) RegisterUser(user *entities.User) (*entities.User, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(*entities.User), args.Error(1)
// }

// func (m *MockAuthService) Login(emailOrUsername, password string) (*entities.RefreshToken, string, error) {
// 	args := m.Called(emailOrUsername, password)
// 	return args.Get(0).(*entities.RefreshToken), args.String(1), args.Error(2)
// }

// func (m *MockAuthService) Logout(userId string) error {
// 	args := m.Called(userId)
// 	return args.Error(0)
// }

// func (m *MockAuthService) RefreshAccessToken(token *entities.RefreshToken) (string, error) {
// 	args := m.Called(token)
// 	return args.String(0), args.Error(1)
// }

// func (m *MockAuthService) VerifyEmail(email string, code string) error {
// 	args := m.Called(email, code)
// 	return args.Error(0)
// }

// func (m *MockAuthService) ResendOtp(request entities.ResendOTPRequest) error {
// 	args := m.Called(request)
// 	return args.Error(0)
// }

// // MockPasswordResetService is a mock implementation of the PasswordResetService interface
// type MockPasswordResetService struct {
// 	mock.Mock
// }

// func (m *MockPasswordResetService) RequestPasswordReset(email string) error {
// 	args := m.Called(email)
// 	return args.Error(0)
// }

// func (m *MockPasswordResetService) ResetPassword(token, newPassword string) error {
// 	args := m.Called(token, newPassword)
// 	return args.Error(0)
// }

// func (m *MockPasswordResetService) GeneratePasswordResetToken(user *entities.User) (string, error) {
// 	args := m.Called(user)
// 	return args.String(0), args.Error(1)
// }

// // SetupTest initializes the test suite
// func (suite *AuthControllerTestSuite) SetupTest() {
// 	suite.authServiceMock = new(MockAuthService)
// 	suite.passwordResetServiceMock = new(MockPasswordResetService)
// 	suite.authController = controller.NewAuthController(suite.authServiceMock, suite.passwordResetServiceMock)

// 	gin.SetMode(gin.TestMode)
// 	suite.router = gin.Default()
// }

// // Test RegisterUser
// func (suite *AuthControllerTestSuite) TestRegisterUser() {
// 	suite.router.POST("/register", suite.authController.RegisterUser)

// 	user := &entities.User{Username: "testuser", Email: "test@example.com", Password: "password123"}
// 	suite.authServiceMock.On("RegisterUser", mock.AnythingOfType("*entities.User")).Return(user, nil)

// 	userRequest := dto.UserCreateRequestDTO{
// 		Username: "testuser",
// 		Email:    "test@example.com",
// 		Password: "password123",
// 	}

// 	jsonValue, _ := json.Marshal(userRequest)
// 	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusCreated, w.Code)
// 	suite.authServiceMock.AssertExpectations(suite.T())
// }

// // Test Login
// func (suite *AuthControllerTestSuite) TestLogin() {
// 	suite.router.POST("/login", suite.authController.Login)

// 	refreshToken := &entities.RefreshToken{
// 		UserID: "userID",
// 		Token:  "refreshToken",
// 	}

// 	suite.authServiceMock.On("Login", "test@example.com", "password123").Return(refreshToken, "accessToken", nil)

// 	loginRequest := dto.LoginDto{
// 		Email:    "test@example.com",
// 		Password: "password123",
// 	}

// 	jsonValue, _ := json.Marshal(loginRequest)
// 	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.authServiceMock.AssertExpectations(suite.T())
// }

// func (suite *AuthControllerTestSuite) TestLogout() {
// 	suite.router.POST("/logout", suite.authController.Logout)

// 	suite.authServiceMock.On("Logout", "userID").Return(nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/logout", nil)
// 	req.Header.Set("userId", "userID")

// 	suite.router.ServeHTTP(w, req)

// 	// Check if the Logout method was called
// 	suite.authServiceMock.AssertCalled(suite.T(), "Logout", "userID")

// 	// Check the response code
// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.authServiceMock.AssertExpectations(suite.T())
// }


// // Test RefreshAccessToken
// func (suite *AuthControllerTestSuite) TestRefreshAccessToken() {
// 	suite.router.POST("/refresh-token", suite.authController.RefreshAccessToken)

// 	refreshToken := &entities.RefreshToken{Token: "oldRefreshToken"}
// 	suite.authServiceMock.On("RefreshAccessToken", refreshToken).Return("newAccessToken", nil)

// 	jsonValue, _ := json.Marshal(refreshToken)
// 	req, _ := http.NewRequest("POST", "/refresh-token", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.authServiceMock.AssertExpectations(suite.T())
// }

// // Test VerifyEmail
// func (suite *AuthControllerTestSuite) TestVerifyEmail() {
// 	suite.router.POST("/verify-email", suite.authController.VerifyEmail)

// 	emailVerification := entities.EmailVerificationRequest{Email: "test@example.com", Code: "123456"}
// 	suite.authServiceMock.On("VerifyEmail", "test@example.com", "123456").Return(nil)

// 	jsonValue, _ := json.Marshal(emailVerification)
// 	req, _ := http.NewRequest("POST", "/verify-email", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.authServiceMock.AssertExpectations(suite.T())
// }

// // Test RequestPasswordReset
// func (suite *AuthControllerTestSuite) TestRequestPasswordReset() {
// 	suite.router.POST("/request-password-reset", suite.authController.RequestPasswordReset)

// 	suite.passwordResetServiceMock.On("RequestPasswordReset", "test@example.com").Return(nil)

// 	forgetPasswordRequest := entities.ForgetPasswordRequest{Email: "test@example.com"}
// 	jsonValue, _ := json.Marshal(forgetPasswordRequest)
// 	req, _ := http.NewRequest("POST", "/request-password-reset", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.passwordResetServiceMock.AssertExpectations(suite.T())
// }

// // Test ResetPassword
// func (suite *AuthControllerTestSuite) TestResetPassword() {
// 	suite.router.POST("/reset-password", suite.authController.ResetPassword)

// 	passwordReset := entities.PasswordReset{Token: "resetToken", NewPassword: "newPassword"}
// 	suite.passwordResetServiceMock.On("ResetPassword", "resetToken", "newPassword").Return(nil)

// 	jsonValue, _ := json.Marshal(passwordReset)
// 	req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.passwordResetServiceMock.AssertExpectations(suite.T())
// }

// // Test ResendOtp
// func (suite *AuthControllerTestSuite) TestResendOtp() {
// 	suite.router.POST("/resend-otp", suite.authController.ResendOtp)

// 	otpRequest := entities.ResendOTPRequest{Email: "test@example.com"}
// 	suite.authServiceMock.On("ResendOtp", otpRequest).Return(nil)

// 	jsonValue, _ := json.Marshal(otpRequest)
// 	req, _ := http.NewRequest("POST", "/resend-otp", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	suite.router.ServeHTTP(w, req)

// 	suite.Equal(http.StatusOK, w.Code)
// 	suite.authServiceMock.AssertExpectations(suite.T())
// }

// // Test suite runner
// func TestAuthControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(AuthControllerTestSuite))
// }
