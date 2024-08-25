package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type forgetPWControllerTestSuite struct {
	suite.Suite
	UserUsecase     *mocks.UserUsecase
	ForgetPWUsecase *mocks.ForgetPWUsecase
	Env             *bootstrap.Env
	Controller      *ForgetPWController
	testingServer   *httptest.Server
}

func (suite *forgetPWControllerTestSuite) SetupTest() {
	UserUsecase := new(mocks.UserUsecase)
	ForgetPWUsecase := new(mocks.ForgetPWUsecase)
	Env := &bootstrap.Env{
		ServerAddress: "http://localhost:8080",
	}
	Controller := NewForgetPWController(ForgetPWUsecase, UserUsecase, Env)

	router := gin.Default()
	router.POST("/forget-password", Controller.ForgetPW)
	router.POST("/reset-password", Controller.ResetPW)

	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer
	suite.UserUsecase = UserUsecase
	suite.ForgetPWUsecase = ForgetPWUsecase
	suite.Env = Env
	suite.Controller = Controller
}

func (suite *forgetPWControllerTestSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *forgetPWControllerTestSuite) TestForgetPW_Success() {
	request := domain.ForgetPWRequest{
		Email: "newuser@example.com",
	}

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{Email: "newuser@example.com"}, nil)
	suite.ForgetPWUsecase.On("ForgetPW", mock.Anything, "newuser@example.com", suite.Env.ServerAddress).Return(nil)

	requestBody, err := json.Marshal(&request)
	suite.NoError(err)

	response, err := http.Post(fmt.Sprintf("%s/forget-password", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err)
	defer response.Body.Close()

	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("reset password request accepted", responseBody.Message)
	suite.UserUsecase.AssertExpectations(suite.T())
	suite.ForgetPWUsecase.AssertExpectations(suite.T())
}

func (suite *forgetPWControllerTestSuite) TestForgetPW_UserNotFound() {
	request := domain.ForgetPWRequest{
		Email: "nonexistentuser@example.com",
	}

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("user not found"))

	requestBody, err := json.Marshal(&request)
	suite.NoError(err)

	response, err := http.Post(fmt.Sprintf("%s/forget-password", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err)
	defer response.Body.Close()

	responseBody := domain.ErrorResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusNotFound, response.StatusCode)
	suite.Equal("user not found", responseBody.Message)
	suite.UserUsecase.AssertExpectations(suite.T())
}

func (suite *forgetPWControllerTestSuite) TestResetPW_Success() {
	request := domain.ResetPWRequest{
		Email:    "newuser@example.com",
		Password: "newpassword",
	}

	suite.ForgetPWUsecase.On("VerifyForgetPWRequest", mock.Anything, "newuser", "validtoken").Return(nil)
	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{Username: "newuser"}, nil)
	suite.UserUsecase.On("UpdateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	requestBody, err := json.Marshal(&request)
	suite.NoError(err)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/reset-password?user=newuser&token=validtoken", suite.testingServer.URL), bytes.NewBuffer(requestBody))
	response, err := http.DefaultClient.Do(req)
	suite.NoError(err)
	defer response.Body.Close()

	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("password reset successful", responseBody.Message)
	suite.UserUsecase.AssertExpectations(suite.T())
	suite.ForgetPWUsecase.AssertExpectations(suite.T())
}

func (suite *forgetPWControllerTestSuite) TestResetPW_InvalidToken() {
	request := domain.ResetPWRequest{
		Email:    "newuser@example.com",
		Password: "newpassword",
	}

	suite.ForgetPWUsecase.On("VerifyForgetPWRequest", mock.Anything, "newuser", "invalidtoken").Return(errors.New("invalid token"))

	requestBody, err := json.Marshal(&request)
	suite.NoError(err)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/reset-password?user=newuser&token=invalidtoken", suite.testingServer.URL), bytes.NewBuffer(requestBody))
	response, err := http.DefaultClient.Do(req)
	suite.NoError(err)
	defer response.Body.Close()

	responseBody := domain.ErrorResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid token", responseBody.Message)
	suite.ForgetPWUsecase.AssertExpectations(suite.T())
}

func TestForgetPWControllerTestSuite(t *testing.T) {
	suite.Run(t, new(forgetPWControllerTestSuite))
}
