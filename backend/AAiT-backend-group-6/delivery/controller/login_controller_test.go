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
	"golang.org/x/crypto/bcrypt"
)

type loginControllerTestSuite struct {
	suite.Suite
	UserUsecase   *mocks.UserUsecase
	LoginUsecase *mocks.LoginUsecase
	Env           *bootstrap.Env
	Controller    *LoginController
	testingServer   *httptest.Server
}


func (suite *loginControllerTestSuite) SetupTest() {
	UserUsecase := new(mocks.UserUsecase)
	LoginUsecase := new(mocks.LoginUsecase)
	Env := &bootstrap.Env{
		AccessTokenSecret:     "access_token",
		AccessTokenExpiryHour: 24,
		RefreshTokenSecret:    "refresh_secret",
		RefreshTokenExpiryHour: 48,
	}
	Controller := NewLoginController(UserUsecase, LoginUsecase, Env)

	router := gin.Default()
	router.POST("/login", Controller.Login)

	// create and run the testing server
	testingServer := httptest.NewServer(router)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.UserUsecase = UserUsecase
	suite.LoginUsecase = LoginUsecase
	suite.Env = Env
	suite.Controller = Controller
}

func (suite *loginControllerTestSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}


func (suite *loginControllerTestSuite) TestLogin_Success() {

	request := domain.LoginRequest{
		Email:    "newuser@example.com",
		Password: "password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{Name: "new user", Password: string(hashedPassword)}, nil)
	suite.UserUsecase.On("UpdateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
	suite.LoginUsecase.On("CreateAccessToken", mock.Anything, suite.Env.AccessTokenSecret, suite.Env.AccessTokenExpiryHour).Return("access_token", nil)
	suite.LoginUsecase.On("CreateRefreshToken", mock.Anything, suite.Env.RefreshTokenSecret, suite.Env.RefreshTokenExpiryHour).Return("refresh_token", nil)
	
	// marshalling and some assertion
	requestBody, err := json.Marshal(&request)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(responseBody.Message, "login successful")
	suite.LoginUsecase.AssertExpectations(suite.T())
}

func (suite *loginControllerTestSuite) TestLogin_EmailNotfound() {

	request := domain.LoginRequest{
		Email:    "newuser@example.com",
		Password: "password",
	}

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("user not found"))
	// marshalling and some assertion
	requestBody, err := json.Marshal(&request)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusNotFound, response.StatusCode)
	suite.Equal(responseBody.Message, "user not found")
	suite.LoginUsecase.AssertExpectations(suite.T())
}
func (suite *loginControllerTestSuite) TestLogin_WrongPassword() {

	request := domain.LoginRequest{
		Email:    "newuser@example.com",
		Password: "password",
	}

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{Name: "new user", Password: "wrong password"}, nil)
	// marshalling and some assertion
	requestBody, err := json.Marshal(&request)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
	suite.Equal(responseBody.Message, "Invalid credentials")
	suite.LoginUsecase.AssertExpectations(suite.T())
}

func TestLoginControllerTestSuite(t *testing.T) {
	suite.Run(t, new(loginControllerTestSuite))
}
