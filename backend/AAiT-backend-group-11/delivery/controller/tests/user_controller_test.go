package controller_test

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/domain/entities"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	userController  *controller.UserController
	mockUserService *MockUserService
}


type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) FindUserByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) FindUserById(userId string) (*entities.User, error) {
	args := m.Called(userId)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(userId string) error {
	args := m.Called(userId)
	return args.Error(0)
}

func (m *MockUserService) PromoteUserToAdmin(userId string) error {
	args := m.Called(userId)
	return args.Error(0)
}

func (m *MockUserService) DemoteUserToRegular(userId string) error {
	args := m.Called(userId)
	return args.Error(0)
}

func (m *MockUserService) MarkUserAsVerified(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUserService = new(MockUserService)
	suite.userController = controller.NewUserController(suite.mockUserService)
}

func (suite *UserControllerTestSuite) TestPromoteUser_Success() {
	suite.mockUserService.On("PromoteUserToAdmin", "123").Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "123"}}

	suite.userController.PromoteUser(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "User promoted successfully")
	suite.mockUserService.AssertCalled(suite.T(), "PromoteUserToAdmin", "123")
}

func (suite *UserControllerTestSuite) TestPromoteUser_Failure() {
	// Simulate an error in the service layer
	suite.mockUserService.On("PromoteUserToAdmin", "123").Return(errors.New("mocked error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "123"}}

	suite.userController.PromoteUser(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "mocked error")
	suite.mockUserService.AssertCalled(suite.T(), "PromoteUserToAdmin", "123")
}

func (suite *UserControllerTestSuite) TestDemoteUser_Success() {
	suite.mockUserService.On("DemoteUserToRegular", "123").Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "123"}}

	suite.userController.DemoteUser(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "User demoted successfully")
	suite.mockUserService.AssertCalled(suite.T(), "DemoteUserToRegular", "123")
}

func (suite *UserControllerTestSuite) TestDemoteUser_Failure() {
	// Simulate an error in the service layer
	suite.mockUserService.On("DemoteUserToRegular", "123").Return(errors.New("mocked error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "123"}}

	suite.userController.DemoteUser(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "mocked error")
	suite.mockUserService.AssertCalled(suite.T(), "DemoteUserToRegular", "123")
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
