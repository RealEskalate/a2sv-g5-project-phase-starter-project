package controller

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mocks"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type promoteControllerTestSuite struct {
	suite.Suite
	UserUsecase    *mocks.UserUsecase
	PromoteUsecase *mocks.PromoteUsecase
	Env            *bootstrap.Env
	Controller     *PromoteController
	testingServer  *httptest.Server
}

func (suite *promoteControllerTestSuite) SetupTest() {
	UserUsecase := new(mocks.UserUsecase)
	PromoteUsecase := new(mocks.PromoteUsecase)
	Env := &bootstrap.Env{
		ServerAddress: "http://localhost:8080",
	}
	Controller := NewPromoteController(UserUsecase, PromoteUsecase, Env)

	router := gin.Default()
	router.POST("/promote/:id", Controller.PromoteUser)
	router.POST("/demote/:id", Controller.DemoteUser)


	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer
	suite.UserUsecase = UserUsecase
	suite.PromoteUsecase = PromoteUsecase
	suite.Env = Env
	suite.Controller = Controller
}

func (suite *promoteControllerTestSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *promoteControllerTestSuite) TestPromoteUser_Success() {
	userID := "123"

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{User_type: "ADMIN"}, nil)
	suite.PromoteUsecase.On("PromoteUser", mock.Anything, userID).Return(nil)

	req := httptest.NewRequest(http.MethodPost, suite.testingServer.URL+"/promote/"+userID, nil)

	req.Header.Set("email", "admin@example.com")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	log.Println(c.Request.Header)
	suite.Controller.PromoteUser(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "user promoted to ADMIN")
	suite.UserUsecase.AssertExpectations(suite.T())
	suite.PromoteUsecase.AssertExpectations(suite.T())
}

func (suite *promoteControllerTestSuite) TestPromoteUser_Unauthorized_UserNotAdmin() {
	userID := "123"

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{User_type: "USER"}, nil)

	req, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/promote/"+userID, nil)
	req.Header.Set("email", "user@example.com")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	suite.Controller.PromoteUser(c)

	suite.Equal(http.StatusUnauthorized, w.Code)
	suite.Contains(w.Body.String(), "Unauthorized")
	suite.UserUsecase.AssertExpectations(suite.T())
}

func (suite *promoteControllerTestSuite) TestPromoteUser_UserNotFound() {
	userID := "123"

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{User_type: "ADMIN"}, nil)
	suite.PromoteUsecase.On("PromoteUser", mock.Anything, userID).Return(errors.New("user with the given userID is not found"))

	req, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/promote/"+userID, nil)
	req.Header.Set("email", "admin@example.com")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	suite.Controller.PromoteUser(c)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.Contains(w.Body.String(), "user with the given userID is not found")
	suite.UserUsecase.AssertExpectations(suite.T())
	suite.PromoteUsecase.AssertExpectations(suite.T())
}

func (suite *promoteControllerTestSuite) TestDemoteUser_Success() {
	userID := "123"

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{User_type: "ADMIN"}, nil)
	suite.PromoteUsecase.On("DemoteUser", mock.Anything, userID).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/demote/"+userID, nil)
	req.Header.Set("email", "admin@example.com")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	suite.Controller.DemoteUser(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "user demoted to USER")
	suite.UserUsecase.AssertExpectations(suite.T())
	suite.PromoteUsecase.AssertExpectations(suite.T())
}

func (suite *promoteControllerTestSuite) TestDemoteUser_Unauthorized_UserNotAdmin() {
	userID := "123"

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{User_type: "USER"}, nil)

	req, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/demote/"+userID, nil)
	req.Header.Set("email", "user@example.com")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	suite.Controller.DemoteUser(c)

	suite.Equal(http.StatusUnauthorized, w.Code)
	suite.Contains(w.Body.String(), "Unauthorized")
	suite.UserUsecase.AssertExpectations(suite.T())
}

func (suite *promoteControllerTestSuite) TestDemoteUser_UserNotFound() {
	userID := "123"

	suite.UserUsecase.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&domain.User{User_type: "ADMIN"}, nil)
	suite.PromoteUsecase.On("DemoteUser", mock.Anything, userID).Return(errors.New("user with the given userID is not found"))

	req, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/demote/"+userID, nil)
	req.Header.Set("email", "admin@example.com")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	suite.Controller.DemoteUser(c)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.Contains(w.Body.String(), "user with the given userID is not found")
	suite.UserUsecase.AssertExpectations(suite.T())
	suite.PromoteUsecase.AssertExpectations(suite.T())
}

func TestPromoteControllerTestSuite(t *testing.T) {
	suite.Run(t, new(promoteControllerTestSuite))
}
