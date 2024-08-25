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

type signupControllerTestSuite struct {
	suite.Suite
	UserUsecase   *mocks.UserUsecase
	SignupUsecase *mocks.SignupUsecase
	Env           *bootstrap.Env
	Controller    *SignupController
	testingServer   *httptest.Server
}


func (suite *signupControllerTestSuite) SetupTest() {
	UserUsecase := new(mocks.UserUsecase)
	SignupUsecase := new(mocks.SignupUsecase)
	Env := &bootstrap.Env{
		AccessTokenSecret:     "access_token",
		AccessTokenExpiryHour: 24,
		RefreshTokenSecret:    "refresh_secret",
		RefreshTokenExpiryHour: 48,
	}
	Controller := NewSignupController(UserUsecase, SignupUsecase, Env)

	router := gin.Default()
	router.POST("/signup", Controller.Signup)

	// create and run the testing server
	testingServer := httptest.NewServer(router)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.UserUsecase = UserUsecase
	suite.SignupUsecase = SignupUsecase
	suite.Env = Env
	suite.Controller = Controller
}

func (suite *signupControllerTestSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}


func (suite *signupControllerTestSuite) TestSignup_Success() {

	request := domain.SignupRequest{
		Name:     "Test User",
		Email:    "newuser@example.com",
		Username: "newuser",
		Password: "password",
	}

	suite.SignupUsecase.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("user not found"))
	suite.UserUsecase.On("GetUserByUsername", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("user not found"))
	suite.SignupUsecase.On("CreateAccessToken", mock.Anything, suite.Env.AccessTokenSecret, suite.Env.AccessTokenExpiryHour).Return("access_token", nil)
	suite.SignupUsecase.On("CreateRefreshToken", mock.Anything, suite.Env.RefreshTokenSecret, suite.Env.RefreshTokenExpiryHour).Return("refresh_token", nil)


	// marshalling and some assertion
	requestBody, err := json.Marshal(&request)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/signup", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	fmt.Println(responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal(responseBody.Message, "user created successfuly")
	suite.SignupUsecase.AssertExpectations(suite.T())
}

func (suite *signupControllerTestSuite) TestSignup_UserAlreadyExistsByEmail() {
	request := domain.SignupRequest{
		Name:     "Test User",
		Email:    "newuser@example.com",
		Username: "newuser",
		Password: "password",
	}

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, nil)
	suite.UserUsecase.On("GetUserByUsername", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, nil)

	// marshalling and some assertion
	requestBody, err := json.Marshal(&request)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/signup", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	fmt.Println(responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusConflict, response.StatusCode)
	suite.Equal(responseBody.Message, "User already exists with the given email")
	suite.SignupUsecase.AssertExpectations(suite.T())
}

func (suite *signupControllerTestSuite) TestSignup_UserAlreadyExistsByUsername() {
	request := domain.SignupRequest{
		Name:     "Test User",
		Email:    "newuser@example.com",
		Username: "newuser",
		Password: "password",
	}

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("user not found"))
	suite.UserUsecase.On("GetUserByUsername", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{}, nil)

	// marshalling and some assertion
	requestBody, err := json.Marshal(&request)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/signup", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	fmt.Println(responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusConflict, response.StatusCode)
	suite.Equal(responseBody.Message, "User already exists with the given username")
	suite.SignupUsecase.AssertExpectations(suite.T())
}


func TestSignupControllerTestSuite(t *testing.T) {
	suite.Run(t, new(signupControllerTestSuite))
}
