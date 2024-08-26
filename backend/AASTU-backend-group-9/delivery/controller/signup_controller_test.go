package controller_test

import (
	"blog/config"
	"blog/delivery/controller"
	"blog/domain"
	"blog/domain/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignupControllerSuite struct {
	suite.Suite
	router           *gin.Engine
	SignupUsecase    *mocks.SignupUsecase
	SignupController *controller.SignupController
}

func (suite *SignupControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.SignupUsecase = new(mocks.SignupUsecase)
	env := &config.Env{
		SMTPUsername: "username",
		SMTPPassword: "password",
	}
	suite.SignupController = &controller.SignupController{
		SignupUsecase: suite.SignupUsecase,
		Env:           env,
	}
	suite.router = gin.Default()
	suite.router.POST("/signup", suite.SignupController.Signup)
	suite.router.POST("/verify-otp", suite.SignupController.VerifyOTP)
}

func (suite *SignupControllerSuite) TearDownTest() {
	suite.SignupUsecase.AssertExpectations(suite.T())
}

func (suite *SignupControllerSuite) TestSignup() {
	suite.Run("Otp_sent_successfully", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.AuthSignup{
			Email:    "non-existing-email@gmail.com",
			Username: "non-existing-username",
		}
		suite.SignupUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.SignupUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(nil, nil).Once()
		suite.SignupUsecase.On("SendOTP", mock.Anything, &user, "username", "password","1533777cbe5eb51a9de765ea723f093bb753862f1a1e9245124dc5ce21eee04f").Return(nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		suite.SignupController.Signup(c)
		expected, err := json.Marshal(gin.H{"message": "OTP sent"})
		suite.Nil(err)
		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Email_already_exists", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.AuthSignup{}
		user.Email = "existingemail@gmailcom"
		user.Username = "non-existing-username"
		suite.SignupUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(&domain.User{}, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		suite.SignupController.Signup(c)
		expected, err := json.Marshal(gin.H{"error": "Email already exists"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Username_already_exists", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.AuthSignup{}
		user.Email = "nonexistingemail@gmailcom"
		user.Username = "existing-username"
		suite.SignupUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.SignupUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(&domain.User{}, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		suite.SignupController.Signup(c)
		expected, err := json.Marshal(gin.H{"error": "Username already exists"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Invalid_request", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/signup", nil)
		suite.SignupController.Signup(c)
		expected, err := json.Marshal(gin.H{"error": "EOF"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Send_otp_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.AuthSignup{}
		user.Email = "nonexistingemail@gmailcom"
		user.Username = "non-existing-username"
		suite.SignupUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.SignupUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(nil, nil).Once()
		suite.SignupUsecase.On("SendOTP", mock.Anything, &user, "username", "password","1533777cbe5eb51a9de765ea723f093bb753862f1a1e9245124dc5ce21eee04f").Return(errors.New("error")).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(payload))
		suite.SignupController.Signup(c)
		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
}

func (suite *SignupControllerSuite) TestVerifyOTP() {
	suite.Run("Otp_verified_successfully", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		otp := domain.OTPRequest{
			Email: "email",
			Value: "123",
		}
		otpresponse := domain.OTP{
			Email:    "email",
			Username: "username",
			Password: "password",
		}
		user := &domain.AuthSignup{
			Username : "username",
			Email : "email",
			Password : "password",
			Role: "user",
		}
		suite.SignupUsecase.On("VerifyOTP", mock.Anything, &otp).Return(&otpresponse, nil).Once()
		suite.SignupUsecase.On("RegisterUser", mock.Anything, user).Return(&user.UserID, nil).Once()
		suite.SignupUsecase.On("CreateAccessToken", user, "", 0).Return("", nil).Once()
		suite.SignupUsecase.On("CreateRefreshToken", user, "", 0).Return("", nil).Once()
		suite.SignupUsecase.On("SaveRefreshToken", mock.Anything, "", user.UserID).Return(nil).Once()
		payload, _ := json.Marshal(otp)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", bytes.NewBuffer(payload))
		suite.SignupController.VerifyOTP(c)
		suite.Equal(http.StatusOK, w.Code)
	
	})
	suite.Run("Invalid_request", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", nil)
		suite.SignupController.VerifyOTP(c)
		expected, err := json.Marshal(gin.H{"error": "EOF"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Invalid_otp", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		otp := domain.OTPRequest{
			Email: "email",
			Value: "123",
		}
		suite.SignupUsecase.On("VerifyOTP", mock.Anything, &otp).Return(nil, errors.New("error")).Once()
		payload, _ := json.Marshal(otp)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", bytes.NewBuffer(payload))
		suite.SignupController.VerifyOTP(c)
		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Register_user_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		otp := domain.OTPRequest{
			Email: "email",
			Value: "123",
		}
		otpresponse := domain.OTP{
			Email:    "email",
			Username: "username",
			Password: "password",
		}
		user := &domain.AuthSignup{
		}
		suite.SignupUsecase.On("VerifyOTP", mock.Anything, &otp).Return(&otpresponse, nil).Once()
		suite.SignupUsecase.On("RegisterUser", mock.Anything, mock.Anything).Return(&user.UserID, errors.New("error")).Once()
		payload, _ := json.Marshal(otp)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", bytes.NewBuffer(payload))
		suite.SignupController.VerifyOTP(c)
		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

	suite.Run("Create_access_token_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		otp := domain.OTPRequest{
			Email: "email",
			Value: "123",
		}
		otpresponse := domain.OTP{
			Email:    "email",
			Username: "username",
			Password: "password",
		}
		user := &domain.AuthSignup{
		}
		suite.SignupUsecase.On("VerifyOTP", mock.Anything, &otp).Return(&otpresponse, nil).Once()
		suite.SignupUsecase.On("RegisterUser", mock.Anything, mock.Anything).Return(&user.UserID, nil).Once()
		suite.SignupUsecase.On("CreateAccessToken", mock.Anything, "", 0).Return("", errors.New("error")).Once()
		payload, _ := json.Marshal(otp)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", bytes.NewBuffer(payload))
		suite.SignupController.VerifyOTP(c)
		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Create_refresh_token_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		otp := domain.OTPRequest{
			Email: "email",
			Value: "123",
		}
		otpresponse := domain.OTP{
			Email:    "email",
			Username: "username",
			Password: "password",
		}
		user := &domain.AuthSignup{
		}
		suite.SignupUsecase.On("VerifyOTP", mock.Anything, &otp).Return(&otpresponse, nil).Once()
		suite.SignupUsecase.On("RegisterUser", mock.Anything, mock.Anything).Return(&user.UserID, nil).Once()
		suite.SignupUsecase.On("CreateAccessToken", mock.Anything, "", 0).Return("", nil).Once()
		suite.SignupUsecase.On("CreateRefreshToken", mock.Anything, "", 0).Return("", errors.New("error")).Once()
		payload, _ := json.Marshal(otp)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", bytes.NewBuffer(payload))
		suite.SignupController.VerifyOTP(c)
		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
	suite.Run("Save_refresh_token_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		otp := domain.OTPRequest{
			Email: "email",
			Value: "123",
		}
		otpresponse := domain.OTP{
			Email:    "email",
			Username: "username",
			Password: "password",
		}
		user := &domain.AuthSignup{
		}
		suite.SignupUsecase.On("VerifyOTP", mock.Anything, &otp).Return(&otpresponse, nil).Once()
		suite.SignupUsecase.On("RegisterUser", mock.Anything, mock.Anything).Return(&user.UserID, nil).Once()
		suite.SignupUsecase.On("CreateAccessToken", mock.Anything, "", 0).Return("", nil).Once()
		suite.SignupUsecase.On("CreateRefreshToken", mock.Anything, "", 0).Return("", nil).Once()
		suite.SignupUsecase.On("SaveRefreshToken", mock.Anything, "", user.UserID).Return(errors.New("error")).Once()
		payload, _ := json.Marshal(otp)
		c.Request = httptest.NewRequest(http.MethodPost, "/verify-otp", bytes.NewBuffer(payload))
		suite.SignupController.VerifyOTP(c)
		expected, err := json.Marshal(gin.H{"error": "error"})
		suite.Nil(err)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Equal(expected, w.Body.Bytes())
	})
}
func TestSignupControllerSuite(t *testing.T) {
	suite.Run(t, new(SignupControllerSuite))
}
