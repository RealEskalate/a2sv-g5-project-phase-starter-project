package tests

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/controllers"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	NewPasswor      string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string
	ResetToken      int `json:"reset_token"`
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserControllerTestSuite struct {
	suite.Suite
	controller  domain.UserController
	userUsecase *mocks.UserUseCase
	router      *gin.Engine
	middleware  *mocks.MiddlewareService
}

func (suite *UserControllerTestSuite) SetupTest() {
	userUC := new(mocks.UserUseCase)
	suite.middleware = new(mocks.MiddlewareService)
	suite.userUsecase = userUC
	suite.controller = controllers.NewUserController(userUC)
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()

	suite.router.POST("/user/register", suite.controller.Register)
	suite.router.GET("/user/verify/:token", suite.controller.VerifyEmail)
	suite.router.POST("/user/login", suite.controller.Login)
	suite.router.POST("/user/forgot_password", suite.controller.ForgotPassword)
	suite.router.POST("/user/reset/:token", suite.controller.ResetPassword)
	suite.router.POST("/user/refresh_token", suite.controller.RefreshToken)

	suite.router.POST("/user/logout", suite.controller.Logout)
	suite.router.POST("/user/update/:id", suite.controller.UpdateProfile)
	suite.router.POST("/user/upload_profile_picture/:id", suite.controller.ImageUpload)

	suite.router.POST("/user/promote", suite.controller.PromoteUser)
	suite.router.POST("/user/demote", suite.controller.DemoteUser)
}
func (suite *UserControllerTestSuite) TearDownTest() {
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRegister_Success() {
	user := domain.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password123",
	}

	suite.userUsecase.On("RegisterStart", mock.Anything, &user).Return(nil)

	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusAccepted, rr.Code)
	suite.JSONEq(`{"Message": "User verification email sent"}`, rr.Body.String())

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRegister_UnmarshalError() {
	invalidReqBody := `{"Username": "testuser", "Email": "testuser@example.com", "Password": 123}` // Invalid password type

	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(invalidReqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	suite.Contains(rr.Body.String(), "Error")

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRegister_RegisterStartError() {
	user := domain.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password123",
	}

	mockError := &domain.CustomError{Message: "User already exists", Code: http.StatusConflict}
	suite.userUsecase.On("RegisterStart", mock.Anything, &user).Return(mockError)

	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusConflict, rr.Code)
	suite.JSONEq(`{"Error": "User already exists"}`, rr.Body.String())

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestVerifyEmail_Success() {
	token := "validToken"
	suite.userUsecase.On("RegisterEnd", mock.Anything, token).Return(nil)

	req := httptest.NewRequest(http.MethodGet, "/user/verify/"+token, nil)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusAccepted, rr.Code)
	suite.JSONEq(`{"Message": "User email verified successfully "}`, rr.Body.String())

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestVerifyEmail_RegisterEndError() {
	token := "invalidToken"
	mockError := &domain.CustomError{Message: "Invalid or expired token", Code: http.StatusUnauthorized}
	suite.userUsecase.On("RegisterEnd", mock.Anything, token).Return(mockError)

	req := httptest.NewRequest(http.MethodGet, "/user/verify/"+token, nil)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusUnauthorized, rr.Code)
	suite.JSONEq(`{"Error": "Invalid or expired token"}`, rr.Body.String())

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin_Success() {
	loginInfo := LoginRequest{Username: "user", Password: "pass"}
	loginResult := map[string]string{"access_token": "access_token", "refresh_token": "refresh_token"}
	suite.userUsecase.On("Login", mock.Anything, loginInfo.Username, loginInfo.Password).Return(loginResult, nil)

	body, _ := json.Marshal(loginInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse, _ := json.Marshal(gin.H{"data": loginResult})
	suite.JSONEq(string(expectedResponse), rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin_InvalidRequest() {
	req := httptest.NewRequest(http.MethodPost, "/user/login", nil)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
}

func (suite *UserControllerTestSuite) TestLogin_LoginError() {
	loginInfo := LoginRequest{Username: "user", Password: "wrongpass"}
	mockError := &domain.CustomError{Message: "Unauthorized", Code: http.StatusUnauthorized}
	suite.userUsecase.On("Login", mock.Anything, loginInfo.Username, loginInfo.Password).Return(nil, mockError)

	body, _ := json.Marshal(loginInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusUnauthorized, rr.Code)
	suite.JSONEq(`{"Error": "Unauthorized"}`, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRefreshToken_Success() {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{
		RefreshToken: "valid_refresh_token",
	}
	refreshResult := map[string]string{"access_token": "new_access_token"}
	suite.userUsecase.On("RefreshToken", mock.Anything, refreshToken.RefreshToken).Return(refreshResult, nil)

	body, _ := json.Marshal(refreshToken)
	req := httptest.NewRequest(http.MethodPost, "/user/refresh_token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"data":{"access_token":"new_access_token"}}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRefreshToken_InvalidTokenFormat() {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{
		RefreshToken: "",
	}
	suite.userUsecase.On("RefreshToken", mock.Anything, refreshToken.RefreshToken).Return(nil, &domain.CustomError{Message: "Invalid token format", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(refreshToken)
	req := httptest.NewRequest(http.MethodPost, "/user/refresh_token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Invalid token format"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRefreshToken_Error() {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{
		RefreshToken: "valid_refresh_token",
	}
	suite.userUsecase.On("RefreshToken", mock.Anything, refreshToken.RefreshToken).Return(nil, &domain.CustomError{Message: "Token refresh error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(refreshToken)
	req := httptest.NewRequest(http.MethodPost, "/user/refresh_token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Token refresh error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRefreshToken_MissingToken() {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{
		RefreshToken: "",
	}
	suite.userUsecase.On("RefreshToken", mock.Anything, refreshToken.RefreshToken).Return(nil, &domain.CustomError{Message: "Refresh token required", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(refreshToken)
	req := httptest.NewRequest(http.MethodPost, "/user/refresh_token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Refresh token required"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestForgotPassword_Success() {
	email := struct {
		Email string `json:"email"`
	}{
		Email: "user@example.com",
	}
	suite.userUsecase.On("ForgotPassword", mock.Anything, email.Email).Return(nil)

	body, _ := json.Marshal(email)
	req := httptest.NewRequest(http.MethodPost, "/user/forgot_password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"Reset link have been sent to the email user@example.com"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestForgotPassword_InvalidEmailFormat() {
	email := struct {
		Email string `json:"email"`
	}{
		Email: "invalid-email",
	}
	suite.userUsecase.On("ForgotPassword", mock.Anything, email.Email).Return(&domain.CustomError{Message: "Invalid email format", Code: http.StatusBadRequest})

	body, _ := json.Marshal(email)
	req := httptest.NewRequest(http.MethodPost, "/user/forgot_password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Invalid email format"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestForgotPassword_EmailNotFound() {
	email := struct {
		Email string `json:"email"`
	}{
		Email: "nonexistent@example.com",
	}
	suite.userUsecase.On("ForgotPassword", mock.Anything, email.Email).Return(&domain.CustomError{Message: "Email not found", Code: http.StatusNotFound})

	body, _ := json.Marshal(email)
	req := httptest.NewRequest(http.MethodPost, "/user/forgot_password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusNotFound, rr.Code)
	expectedResponse := `{"Error":"Email not found"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestForgotPassword_Error() {
	email := struct {
		Email string `json:"email"`
	}{
		Email: "user@example.com",
	}
	suite.userUsecase.On("ForgotPassword", mock.Anything, email.Email).Return(&domain.CustomError{Message: "Unexpected error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(email)
	req := httptest.NewRequest(http.MethodPost, "/user/forgot_password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Unexpected error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestResetPassword_Success() {
	resetInfo := ResetPasswordRequest{
		NewPasswor:      "newpassword",
		ConfirmPassword: "newpassword",
		Token:           "validToken",
		ResetToken:      1234,
	}
	suite.userUsecase.On("ResetPassword", mock.Anything, resetInfo.NewPasswor, resetInfo.ConfirmPassword, resetInfo.Token, resetInfo.ResetToken).Return(nil)

	body, _ := json.Marshal(resetInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/reset/validToken", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"Password reset successfully"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestResetPassword_PasswordMismatch() {
	resetInfo := ResetPasswordRequest{
		NewPasswor:      "newpassword",
		ConfirmPassword: "differentpassword",
		Token:           "validToken",
		ResetToken:      1234,
	}
	suite.userUsecase.On("ResetPassword", mock.Anything, resetInfo.NewPasswor, resetInfo.ConfirmPassword, resetInfo.Token, resetInfo.ResetToken).Return(&domain.CustomError{Message: "Passwords do not match", Code: http.StatusBadRequest})

	body, _ := json.Marshal(resetInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/reset/validToken", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Passwords do not match"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestResetPassword_InvalidToken() {
	resetInfo := ResetPasswordRequest{
		NewPasswor:      "newpassword",
		ConfirmPassword: "newpassword",
		Token:           "invalidToken",
		ResetToken:      1234,
	}
	suite.userUsecase.On("ResetPassword", mock.Anything, resetInfo.NewPasswor, resetInfo.ConfirmPassword, resetInfo.Token, resetInfo.ResetToken).Return(&domain.CustomError{Message: "Invalid reset token", Code: http.StatusBadRequest})

	body, _ := json.Marshal(resetInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/reset/invalidToken", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Invalid reset token"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestResetPassword_Error() {
	resetInfo := ResetPasswordRequest{
		NewPasswor:      "newpassword",
		ConfirmPassword: "newpassword",
		Token:           "validToken",
		ResetToken:      1234,
	}
	suite.userUsecase.On("ResetPassword", mock.Anything, resetInfo.NewPasswor, resetInfo.ConfirmPassword, resetInfo.Token, resetInfo.ResetToken).Return(&domain.CustomError{Message: "Unexpected error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(resetInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/reset/validToken", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Unexpected error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogout_Success() {
	logoutInfo := LogoutRequest{
		RefreshToken: "refreshToken",
	}
	accessToken := "Bearer accessToken"
	suite.userUsecase.On("Logout", mock.Anything, map[string]string{
		"access_token":  "accessToken",
		"refresh_token": logoutInfo.RefreshToken,
	}).Return(nil)

	body, _ := json.Marshal(logoutInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/logout", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"User logged out successfully"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogout_InvalidRequest() {
	logoutInfo := LogoutRequest{
		RefreshToken: "refreshToken",
	}
	suite.userUsecase.On("Logout", mock.Anything, mock.Anything).Return(&domain.CustomError{Message: "Invalid request", Code: http.StatusBadRequest})

	body, _ := json.Marshal(logoutInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/logout", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Invalid request"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogout_Error() {
	logoutInfo := LogoutRequest{
		RefreshToken: "refreshToken",
	}
	suite.userUsecase.On("Logout", mock.Anything, mock.Anything).Return(&domain.CustomError{Message: "Unexpected error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(logoutInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/logout", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Unexpected error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteUser_Success() {
	updateID := map[string]string{"id": "userID"}
	suite.userUsecase.On("PromoteUser", mock.Anything, updateID["id"]).Return(nil)

	body, _ := json.Marshal(updateID)
	req := httptest.NewRequest(http.MethodPost, "/user/promote", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"User promoted successfully"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteUser_BadRequest() {
	updateID := map[string]string{"id": "userID"}
	suite.userUsecase.On("PromoteUser", mock.Anything, updateID["id"]).Return(&domain.CustomError{Message: "Invalid ID", Code: http.StatusBadRequest})

	body, _ := json.Marshal(updateID)
	req := httptest.NewRequest(http.MethodPost, "/user/promote", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Invalid ID"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteUser_Error() {
	updateID := map[string]string{"id": "userID"}
	suite.userUsecase.On("PromoteUser", mock.Anything, updateID["id"]).Return(&domain.CustomError{Message: "Unexpected error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(updateID)
	req := httptest.NewRequest(http.MethodPost, "/user/promote", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Unexpected error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDemoteUser_Success() {
	updateID := map[string]string{"id": "userID"}
	suite.userUsecase.On("DemoteUser", mock.Anything, updateID["id"]).Return(nil)

	body, _ := json.Marshal(updateID)
	req := httptest.NewRequest(http.MethodPost, "/user/demote", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"User demoted successfully"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDemoteUser_BadRequest() {
	updateID := map[string]string{"id": "userID"}
	suite.userUsecase.On("DemoteUser", mock.Anything, updateID["id"]).Return(&domain.CustomError{Message: "Invalid ID", Code: http.StatusBadRequest})

	body, _ := json.Marshal(updateID)
	req := httptest.NewRequest(http.MethodPost, "/user/demote", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Invalid ID"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDemoteUser_Error() {
	updateID := map[string]string{"id": "userID"}
	suite.userUsecase.On("DemoteUser", mock.Anything, updateID["id"]).Return(&domain.CustomError{Message: "Unexpected error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(updateID)
	req := httptest.NewRequest(http.MethodPost, "/user/demote", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Unexpected error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestUpdateProfile_Success() {
	updateInfo := map[string]interface{}{"name": "John Doe", "email": "john@example.com"}
	userID := "userID"
	suite.userUsecase.On("UpdateProfile", mock.Anything, userID, updateInfo).Return(nil)

	body, _ := json.Marshal(updateInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/update/"+userID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"User profile updated successfully"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestUpdateProfile_BadRequest() {
	updateInfo := map[string]interface{}{"name": "John Doe", "email": "john@example.com"}
	userID := "userID"
	suite.userUsecase.On("UpdateProfile", mock.Anything, userID, updateInfo).Return(&domain.CustomError{Message: "Invalid data", Code: http.StatusBadRequest})

	body, _ := json.Marshal(updateInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/update/"+userID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Invalid data"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestImageUpload_Success() {
	fileContent := []byte("file content")
	header := &multipart.FileHeader{Filename: "test_image.png"}
	userID := "userID"

	suite.userUsecase.On("ImageUpload", mock.Anything, mock.Anything, mock.Anything, userID).Return(nil)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("profile_picture", header.Filename)
	part.Write(fileContent)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/user/upload_profile_picture/"+userID, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code)
	expectedResponse := `{"Message":"Image uploaded successfully"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestUpdateProfile_Error() {
	updateInfo := map[string]interface{}{"name": "John Doe", "email": "john@example.com"}
	userID := "userID"
	suite.userUsecase.On("UpdateProfile", mock.Anything, userID, updateInfo).Return(&domain.CustomError{Message: "Unexpected error", Code: http.StatusInternalServerError})

	body, _ := json.Marshal(updateInfo)
	req := httptest.NewRequest(http.MethodPost, "/user/update/"+userID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"Error":"Unexpected error"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestImageUpload_BadRequest() {
	fileContent := []byte("file content")
	header := &multipart.FileHeader{Filename: "test_image.png"}
	userID := "userID"

	suite.userUsecase.On("ImageUpload", mock.Anything, mock.Anything, mock.Anything, userID).Return(&domain.CustomError{Message: "Upload failed", Code: http.StatusBadRequest})

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("profile_picture", header.Filename)
	part.Write(fileContent)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/user/upload_profile_picture/"+userID, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer accessToken")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := `{"Error":"Upload failed"}`
	suite.JSONEq(expectedResponse, rr.Body.String())
	suite.userUsecase.AssertExpectations(suite.T())
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
