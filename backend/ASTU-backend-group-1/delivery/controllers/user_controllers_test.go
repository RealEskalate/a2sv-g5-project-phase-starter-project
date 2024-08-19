package controllers

import (
	"astu-backend-g1/domain"
	"astu-backend-g1/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	userController *userController
	userUsecase    *mocks.UserUsecase
	data           []domain.User
}

func (suite *UserControllerTestSuite) SetupSuite() {
	suite.userUsecase = mocks.NewUserUsecase(suite.T())
	suite.userController = NewUserController(suite.userUsecase)
	suite.data = []domain.User{
		{
			ID:        "1",
			Username:  "john_doe",
			Email:     "john.doe@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "hashed_password_1",
			IsAdmin:   false,
			IsActive:  true,
		},
		{
			ID:        "2",
			Username:  "jane_smith",
			Email:     "jane.smith@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
			Password:  "hashed_password_2",
			IsAdmin:   true,
			IsActive:  true,
		},
		{
			ID:        "3",
			Username:  "mike_jones",
			Email:     "mike.jones@example.com",
			FirstName: "Mike",
			LastName:  "Jones",
			Password:  "hashed_password_3",
			IsAdmin:   false,
			IsActive:  false,
		},
	}
}

func (suite *UserControllerTestSuite) TestGetUsers() {
	assert := assert.New(suite.T())
	suite.T().Parallel()
	suite.Run("getting all users", func() {
		suite.userUsecase.On("Get").Return(suite.data, nil).Once()
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/users", nil)
		ctx.Request = req
		suite.userController.GetUsers(ctx)
		assert.Equal(http.StatusOK, w.Code)
		data, err := json.Marshal(suite.data)
		assert.NoError(err)
		assert.JSONEq(string(data), w.Body.String())
	})
	suite.Run("getting user by username", func() {
		for _, user := range suite.data {
			suite.userUsecase.On("GetByUsername", user.Username).Return(user, nil).Once()
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest("GET", "/users?username="+user.Username, nil)
			ctx.Request = req
			suite.userController.GetUsers(ctx)
			assert.Equal(http.StatusOK, w.Code)
			data, err := json.Marshal(user)
			assert.NoError(err)
			assert.JSONEq(string(data), w.Body.String())
		}
	})
	suite.Run("getting user by email", func() {
		for _, user := range suite.data {
			suite.userUsecase.On("GetByEmail", user.Email).Return(user, nil).Once()
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest("GET", "/users?email="+user.Email, nil)
			ctx.Request = req
			suite.userController.GetUsers(ctx)
			assert.Equal(http.StatusOK, w.Code)
			data, err := json.Marshal(user)
			assert.NoError(err)
			assert.JSONEq(string(data), w.Body.String())
		}
	})
	suite.Run("error", func() {
		expectedError := fmt.Errorf("there is no users in the database")
		suite.userUsecase.On("Get").Return([]domain.User{}, expectedError).Once()
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/users", nil)
		ctx.Request = req
		suite.userController.GetUsers(ctx)
		assert.Equal(http.StatusNotFound, w.Code)
		assert.Contains(w.Body.String(), expectedError.Error())
	})
}

func (suite *UserControllerTestSuite) TestGetUserByID() {
	assert := assert.New(suite.T())
	suite.T().Parallel()
	suite.Run("success", func() {
		for _, user := range suite.data {
			suite.userUsecase.On("GetByID", user.ID).Return(user, nil).Once()
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest("GET", "/users/:id", nil)
			ctx.Request = req
			ctx.Params = gin.Params{{Key: "id", Value: user.ID}}
			suite.userController.GetUserByID(ctx)
			assert.Equal(http.StatusOK, w.Code)
			data, err := json.Marshal(user)
			assert.NoError(err)
			assert.JSONEq(string(data), w.Body.String())
		}
	})
	suite.Run("not found", func() {
		expectedError := fmt.Errorf("there is no user with the given id")
		suite.userUsecase.On("GetByID", "4").Return(domain.User{}, expectedError).Once()
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/users/:id", nil)
		ctx.Request = req
		ctx.Params = gin.Params{{Key: "id", Value: "4"}}
		suite.userController.GetUserByID(ctx)
		assert.Contains(w.Body.String(), expectedError.Error())
	})
}

func (suite *UserControllerTestSuite) TestDeleteUser() {
	assert := assert.New(suite.T())
	suite.T().Parallel()
	suite.Run("success", func() {
		for _, user := range suite.data {
			suite.userUsecase.On("Delete", user.ID).Return(nil).Once()
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest("DELETE", "/users/:id", nil)
			ctx.Request = req
			ctx.Params = gin.Params{{Key: "id", Value: user.ID}}
			suite.userController.DeleteUser(ctx)
			assert.Equal(http.StatusNoContent, w.Code)
		}
	})
}

func (suite *UserControllerTestSuite) TestUpdateUser() {
	assert := assert.New(suite.T())
	suite.Run("success", func() {
		user := suite.data[0]
		suite.userUsecase.On("Update", user.ID, mock.AnythingOfType("User")).Return(user, nil).Once()
		w := httptest.NewRecorder()
		data, err := json.Marshal(user)
		ctx, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("PUT", "/users/:id", bytes.NewBuffer(data))
		assert.Nil(err)
		ctx.Request = req
		ctx.Params = gin.Params{{Key: "id", Value: user.ID}}
		suite.userController.UpdateUser(ctx)
		suite.T().Log(w.Body.String())
		assert.Equal(http.StatusOK, w.Code)
		assert.JSONEq(string(data), w.Body.String())
	})
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
