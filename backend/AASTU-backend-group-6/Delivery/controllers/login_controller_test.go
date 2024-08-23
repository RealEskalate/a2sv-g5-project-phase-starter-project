package controllers

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"blogs/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type LoginControllerTestSuite struct {
	suite.Suite
	loginUsecase *mocks.LoginUsecase
	env          *infrastructure.Config
	controller   *LoginController
}

func (suite *LoginControllerTestSuite) SetupTest() {
	suite.loginUsecase = new(mocks.LoginUsecase)
	suite.env = &infrastructure.Config{
		AccessTokenSecret:      "access_token_secret",
		AccessTokenExpiryHour:  1,
		RefreshTokenSecret:     "refresh_token_secret",
		RefreshTokenExpiryHour: 24,
	}
	suite.controller = &LoginController{
		LoginUsecase: suite.loginUsecase,
		Env:          suite.env,
	}
}

func (suite *LoginControllerTestSuite) TestLogin_Success() {
	// Setup
	request := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: string(hash),
	}
	accessToken := "access_token"
	refreshToken := "refresh_token"
	activeUser := domain.ActiveUser{
		ID:           user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "UserAgent",
	}

	suite.loginUsecase.On("GetUserByEmail", mock.Anything, request.Email).Return(user, nil).Once()
	suite.loginUsecase.On("CreateAccessToken", &user, suite.env.AccessTokenSecret, suite.env.AccessTokenExpiryHour).Return(accessToken, nil).Once()
	suite.loginUsecase.On("CreateRefreshToken", &user, suite.env.RefreshTokenSecret, suite.env.RefreshTokenExpiryHour).Return(refreshToken, nil).Once()
	suite.loginUsecase.On("SaveAsActiveUser", activeUser, refreshToken, mock.Anything).Return(nil).Once()

	// Encode request to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		suite.T().Fatal(err)
	}

	// Execute
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("User-Agent", activeUser.UserAgent)
	c.Request.Header.Set("Content-Type", "application/json")
	suite.controller.Login(c)

	// Assertions
	suite.Equal(200, w.Code)
	expectedResponse := gin.H{"response": map[string]interface{}{"accessToken": "access_token", "message": "", "refreshToken": "refresh_token"}, "user agent": "UserAgent"}
	var actualResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(expectedResponse, actualResponse)
}

func (suite *LoginControllerTestSuite) TestLogin_InvalidRequest() {
	// Setup
	request := domain.LoginRequest{
		Email:    "",
		Password: "password",
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		suite.T().Fatal(err)
	}
	// Execute
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
	c.Set("email", request.Email)
	c.Set("password", request.Password)
	suite.controller.Login(c)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Equal(suite.T(), string("{\"message\":\"Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag\\nKey: 'LoginRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag\",\"status\":0}"), w.Body.String())
}

func (suite *LoginControllerTestSuite) TestLogin_UserNotFound() {
    // Setup
    request := domain.LoginRequest{
        Email:    "test@example.com",
        Password: "password",
    }

    suite.loginUsecase.On("GetUserByEmail", mock.Anything, request.Email).Return(domain.User{}, errors.New("user not found")).Once()

    requestBody, err := json.Marshal(request)
    if err != nil {
        suite.T().Fatal(err)
    }

    // Execute
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
    c.Request.Header.Set("Content-Type", "application/json")
    suite.controller.Login(c)

    // Assert
    assert.Equal(suite.T(), http.StatusNotFound, w.Code)
    expectedResponse := gin.H{"message": "User not found with the given email", "status": float64(http.StatusNotFound)}
    var actualResponse gin.H
    err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
    if err != nil {
        suite.T().Fatal(err)
    }
    assert.Equal(suite.T(), expectedResponse, actualResponse)
}

func (suite *LoginControllerTestSuite) TestLogin_InvalidCredentials() {
	// Setup
	request := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "$2a$10$XV8WY4vZQ9L0X1tZ6Z1K7e8X6Z1K7e8X6Z1K7e8X6Z1K7e8X6Z1K7e",
		
	}

	suite.loginUsecase.On("GetUserByEmail", mock.Anything, request.Email).Return(user, nil)

	// Execute
	requestBody, err := json.Marshal(request)
    if err != nil {
        suite.T().Fatal(err)
    }

    // Execute
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
    c.Request.Header.Set("Content-Type", "application/json")
    suite.controller.Login(c)


	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.Equal(suite.T(),  string("{\"message\":\"Invalid credentials\",\"status\":0}"), w.Body.String())

	suite.loginUsecase.AssertExpectations(suite.T())
}

func (suite *LoginControllerTestSuite) TestLogin_UserNotVerified() {
	// Setup
	request := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: string(hash),
		
	}

	suite.loginUsecase.On("GetUserByEmail", mock.Anything, request.Email).Return(user, nil)

	// Execute
	requestBody, err := json.Marshal(request)
    if err != nil {
        suite.T().Fatal(err)
    }

    // Execute
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
    c.Request.Header.Set("Content-Type", "application/json")
    suite.controller.Login(c)

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.Equal(suite.T(), string("{\"message\":\"User not verified\",\"status\":0}"), w.Body.String())

	suite.loginUsecase.AssertExpectations(suite.T())
}

func TestLoginControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LoginControllerTestSuite))
}
