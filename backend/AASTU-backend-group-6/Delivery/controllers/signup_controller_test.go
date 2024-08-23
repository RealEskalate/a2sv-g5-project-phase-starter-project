package controllers_test

import (
	controllers "blogs/Delivery/controllers"
	domain "blogs/Domain"
	"blogs/mocks"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignupControllerTestSuite struct {
	suite.Suite
	mockSignupUsecase *mocks.SignupUseCase
	signupController  *controllers.SignupController
}

func (suite *SignupControllerTestSuite) SetupTest() {
	suite.mockSignupUsecase = new(mocks.SignupUseCase)
	suite.signupController = &controllers.SignupController{
		SignupUsecase: suite.mockSignupUsecase,
	}
}

func (suite *SignupControllerTestSuite) TestSignup_Success() {
	// Prepare the request and context
	signupRequest := domain.SignUpRequest{
		Email:    "test@example.com",
		Username: "testuser",
		Password: "TestPassword123",
	}
	body, _ := json.Marshal(signupRequest)
	req, err := http.NewRequest(http.MethodPost, "/signup", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Mock the Create method to return a successful response
	suite.mockSignupUsecase.On("Create", mock.Anything, mock.Anything).Return(&domain.SuccessResponse{
		Message: "Registered Successfully, Verify your account",
		Status:  http.StatusOK,
	}).Once()

	// Call the Signup method
	suite.signupController.Signup(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	var response domain.SuccessResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("Registered Successfully, Verify your account", response.Message)

	// Assert that the mock was called with the correct parameters
	suite.mockSignupUsecase.AssertCalled(suite.T(), "Create", mock.Anything, mock.AnythingOfType("domain.User"))
}

func (suite *SignupControllerTestSuite) TestSignup_InvalidRequest() {
	// Prepare an invalid request with missing fields
	signupRequest := domain.SignUpRequest{
		Email:    "",
		Username: "",
		Password: "",
	}
	body, _ := json.Marshal(signupRequest)
	req, err := http.NewRequest(http.MethodPost, "/signup", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Mock the Create method to return an error response
	suite.mockSignupUsecase.On("Create", mock.Anything, mock.AnythingOfType("domain.User")).Return(&domain.ErrorResponse{
		Message: "All fields are required",
		Status:  http.StatusBadRequest,
	}).Once()

	// Call the Signup method
	suite.signupController.Signup(c)

	// Verify the response status code
	suite.Equal(http.StatusBadRequest, w.Code)

	// Decode the response and verify the error message
	var response domain.ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("All fields are required", response.Message)
	suite.Equal(http.StatusBadRequest, response.Status)

	// Assert that the mock was called with the correct parameters
	suite.mockSignupUsecase.AssertCalled(suite.T(), "Create", mock.Anything, mock.AnythingOfType("domain.User"))
}

func (suite *SignupControllerTestSuite) TestVerifyOTP_Success() {
	// Prepare a valid OTP token request
	otpRequest := domain.OtpToken{
		Email: "test@example.com",
		OTP:   "valid-otp",
	}
	body, _ := json.Marshal(otpRequest)
	req, err := http.NewRequest(http.MethodPost, "/signup/verify", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Mock the VerifyOTP method to return a successful response
	verifiedUser := domain.User{Email: otpRequest.Email}
	suite.mockSignupUsecase.On("VerifyOTP", mock.Anything, otpRequest).Return(&domain.SuccessResponse{
		Message: "Account verified successfully",
		Data:    verifiedUser,
		Status:  http.StatusOK,
	}).Once()

	// Call the VerifyOTP method
	suite.signupController.VerifyOTP(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	var response domain.SuccessResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("Account verified successfully", response.Message)
	suite.Equal(http.StatusOK, response.Status)

	// Check if response.Data is nil
	if response.Data == nil {
		suite.T().Fatal("expected Data to be non-nil")
	}

	// Convert response.Data to map[string]interface{}
	dataMap, ok := response.Data.(map[string]interface{})
	if !ok {
		suite.T().Fatal("expected Data to be of type map[string]interface{}")
	}

	// Ensure that "Email" and "Verified" are not nil before asserting
	email, emailOk := dataMap["email"].(string)
	if !emailOk || email == "" {
		suite.T().Fatal("expected Email to be a non-empty string")
	}
	
	// Convert map to domain.User
	decodedUser := domain.User{
		Email:    email,
		
	}

	// Assert that the user was verified successfully
	suite.Equal(verifiedUser.Email, decodedUser.Email)


	// Assert that the mock was called with the correct parameters
	suite.mockSignupUsecase.AssertCalled(suite.T(), "VerifyOTP", mock.Anything, otpRequest)
}

func (suite *SignupControllerTestSuite) TestForgotPassword_Success() {
	// Prepare a valid ForgotPasswordRequest
	userEmail := domain.ForgotPasswordRequest{
		Email: "test@example.com",
	}
	body, _ := json.Marshal(userEmail)
	req, err := http.NewRequest(http.MethodPost, "/signup/forgot-password", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Mock the ForgotPassword method to return a successful response
	suite.mockSignupUsecase.On("ForgotPassword", mock.Anything, userEmail).Return(&domain.SuccessResponse{
		Message: "Reset email sent",
		Data:    "",
		Status:  http.StatusOK,
	}).Once()

	// Call the ForgotPassword method
	suite.signupController.ForgotPassword(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	var response domain.SuccessResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("Reset email sent", response.Message)
	suite.Equal(http.StatusOK, response.Status)

	// Assert that the mock was called with the correct parameters
	suite.mockSignupUsecase.AssertCalled(suite.T(), "ForgotPassword", mock.Anything, userEmail)
}

func (suite *SignupControllerTestSuite) TestResetPassword_Success() {
	// Prepare a ResetPasswordRequest
	resetPasswordRequest := domain.ResetPasswordRequest{
		Password: "new-strong-password",
	}
	token := "valid-reset-token"
	body, _ := json.Marshal(resetPasswordRequest)
	req, err := http.NewRequest(http.MethodPost, "/signup/reset-password?token="+token, nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Mock the ResetPassword method to return a successful response
	suite.mockSignupUsecase.On("ResetPassword", mock.Anything, resetPasswordRequest, token).Return(&domain.SuccessResponse{
		Message: "Password Reset Successfully",
		Status:  http.StatusOK,
	}).Once()

	// Call the ResetPassword method
	suite.signupController.ForgotPassword(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	var response domain.SuccessResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal("Password Reset Successfully", response.Message)
	suite.Equal(http.StatusOK, response.Status)

	// Assert that the mock was called with the correct parameters
	suite.mockSignupUsecase.AssertCalled(suite.T(), "ResetPassword", mock.Anything, resetPasswordRequest, token)
}

func TestSignupControllerTestSuite(t *testing.T) {
	suite.Run(t, new(SignupControllerTestSuite))
}
