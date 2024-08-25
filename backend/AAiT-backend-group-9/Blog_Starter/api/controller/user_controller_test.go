package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "Blog_Starter/utils"
    "context"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
	"errors"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "bou.ke/monkey"
)

// Test Suite
type UserControllerTestSuite struct {
    suite.Suite
    controller *UserController
    useCase    *mocks.UserUsecase
}

func (suite *UserControllerTestSuite) SetupTest() {
    suite.useCase = new(mocks.UserUsecase)
    suite.controller = NewUserController(suite.useCase)
}

func (suite *UserControllerTestSuite) TestGetAllUsers() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user", Role: "superAdmin"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case
    suite.useCase.On("GetAllUser", mock.Anything).Return([]*domain.UserResponse{}, nil)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("GET", "/users", nil)

    // Call the controller method
    suite.controller.GetAllUsers(c)

    // Assertions
    suite.Equal(http.StatusOK, w.Code)
}

func (suite *UserControllerTestSuite) TestGetAllUsers_Unauthorized() {
    // Mock CheckUser function to return an error
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return nil, context.DeadlineExceeded
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("GET", "/users", nil)

    // Call the controller method
    suite.controller.GetAllUsers(c)

    // Assertions
    suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *UserControllerTestSuite) TestPromoteUser() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user", Role: "superAdmin"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case
    suite.useCase.On("PromoteUser", mock.Anything, "test_id").Return(nil)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("POST", "/promote", strings.NewReader(`"test_id"`))

    // Call the controller method
    suite.controller.PromoteUser(c)

    // Assertions
    suite.Equal(http.StatusOK, w.Code)
    suite.Contains(w.Body.String(), "User promoted to admin")
}

func (suite *UserControllerTestSuite) TestPromoteUser_Unauthorized() {
    // Mock CheckUser function to return a non-superAdmin user
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user", Role: "admin"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("POST", "/promote", strings.NewReader(`"test_id"`))

    // Call the controller method
    suite.controller.PromoteUser(c)

    // Assertions
    suite.Equal(http.StatusUnauthorized, w.Code)
    suite.Contains(w.Body.String(), "Unauthorized: Only the super-admin can promote a user")
}

func (suite *UserControllerTestSuite) TestPromoteUser_InternalServerError() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user", Role: "superAdmin"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case to return an error
    suite.useCase.On("PromoteUser", mock.Anything, "test_id").Return(errors.New("internal server error"))

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("POST", "/promote", strings.NewReader(`"test_id"`))

    // Call the controller method
    suite.controller.PromoteUser(c)

    // Assertions
    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.Contains(w.Body.String(), "internal server error")
}

func (suite *UserControllerTestSuite) TestDeleteUser() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user", Role: "user"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case
    suite.useCase.On("DeleteUser", mock.Anything, "test_user", "password").Return(nil)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("DELETE", "/user", strings.NewReader(`"password"`))

    // Call the controller method
    suite.controller.DeleteUser(c)

    // Assertions
    suite.Equal(http.StatusOK, w.Code)
    suite.Contains(w.Body.String(), "User deleted")
}

func (suite *UserControllerTestSuite) TestDeleteUser_Unauthorized() {
    // Mock CheckUser function to return an error
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return nil, context.DeadlineExceeded
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("DELETE", "/user", strings.NewReader(`"password"`))

    // Call the controller method
    suite.controller.DeleteUser(c)

    // Assertions
    suite.Equal(http.StatusUnauthorized, w.Code)
}

func (suite *UserControllerTestSuite) TestDeleteUser_InternalServerError() {
    // Mock CheckUser function
    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return &domain.AuthenticatedUser{UserID: "test_user", Role: "user"}, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    // Mock use case to return an error
    suite.useCase.On("DeleteUser", mock.Anything, "test_user", "password").Return(errors.New("internal server error"))

    // Create a request
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = httptest.NewRequest("DELETE", "/user", strings.NewReader(`"password"`))

    // Call the controller method
    suite.controller.DeleteUser(c)

    // Assertions
    suite.Equal(http.StatusInternalServerError, w.Code)
    suite.Contains(w.Body.String(), "internal server error")
}

func TestUserControllerTestSuite(t *testing.T) {
    suite.Run(t, new(UserControllerTestSuite))
}