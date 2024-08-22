package controllers_test

import (
	"blogapp/Domain"
	"blogapp/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "blogapp/Delivery/controllers"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerTestSuite struct {
	suite.Suite
	mockUserUseCase *mocks.UserUseCases
	userController  *controllers.UserController
	patch           *monkey.PatchGuard
	userID          primitive.ObjectID
	Role            string
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUserUseCase = new(mocks.UserUseCases)
	var err error
	suite.userController = controllers.NewUserController(suite.mockUserUseCase)
	assert.NoError(suite.T(), err)
	suite.userID = user_id
	suite.Role = "admin"
	suite.patch = monkey.Patch(controllers.Getclaim, mockExtractUser)
}

func (suite *UserControllerTestSuite) TestGetUsers() {
	// Arrange
	users := []*Domain.OmitedUser{
		{
			ID:    primitive.NewObjectID(),
			Email: "user1@example.com",
		},
		{
			ID:    primitive.NewObjectID(),
			Email: "user2@example.com",
		},
	}

	suite.mockUserUseCase.On("GetUsers", mock.Anything).Return(users, nil, 200).Once()

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.userController.GetUsers(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// assert.Equal(suite.T(), "{\"users\":[{\"ID\":\"", w.Body.String())
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestGetUsersError() {
	// Arrange
	suite.mockUserUseCase.On("GetUsers", mock.Anything).Return(nil, errors.New("get users error"), 500).Once()

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	user_id := primitive.NewObjectID().Hex()
	c.Set("user_id", user_id)
	// Act
	suite.userController.GetUsers(c)

	// Assert
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	// assert.Equal(suite.T(), "{\"error\":\"get users error\"}", w.Body.String())
	suite.mockUserUseCase.AssertExpectations(suite.T())
}
func (suite *UserControllerTestSuite) TestGetUser() {
	// Arrange
	loged_user := Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	}
	// Create a mock user with type Domain.OmitedUser
	mockUser := Domain.OmitedUser{}

	suite.mockUserUseCase.On("GetUsersById", mock.Anything, loged_user.ID, loged_user).Return(mockUser, nil, 200).Once()

	req, _ := http.NewRequest(http.MethodGet, "/users/"+loged_user.ID.Hex(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: loged_user.ID.Hex()},
	}
	c.Set("claim", &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	})

	// Act
	suite.userController.GetUser(c)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

// In user_controller_test.go
func (suite *UserControllerTestSuite) TestCreateUser() {
	// Arrange
	user := Domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	createdUser := Domain.OmitedUser{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	suite.mockUserUseCase.On("CreateUser", mock.Anything, &user).Return(createdUser, nil, http.StatusCreated).Once()

	// Act
	suite.userController.CreateUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestUpdateUser() {
	// Arrange
	userID := suite.userID
	user := Domain.User{
		ID:    userID,
		Email: "updated@example.com",
	}

	currentUser := Domain.AccessClaims{
		ID:   userID,
		Role: "admin",
	}

	updatedUser := Domain.OmitedUser{
		ID:    userID,
		Email: "updated@example.com",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/users/"+userID.Hex(), bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: userID.Hex()},
	}
	c.Set("claim", &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	})
	suite.mockUserUseCase.On("UpdateUsersById", c, userID, user, currentUser).Return(updatedUser, nil, 200).Once()

	// Act
	suite.userController.UpdateUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDeleteUser() {
	// Arrange
	currentUser := Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	}

	req, _ := http.NewRequest(http.MethodDelete, "/users/"+currentUser.ID.Hex(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: currentUser.ID.Hex()},
	}
	c.Set("claim", &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	})

	suite.mockUserUseCase.On("DeleteUsersById", c, currentUser.ID, currentUser).Return(nil, 200).Once()

	// Act
	suite.userController.DeleteUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// assert.Equal(suite.T(), "{\"message\":\"User deleted successfully\"}", w.Body.String())
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteUser() {
	// Arrange
	user := &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	}
	expected_User := Domain.OmitedUser{
		ID:    suite.userID,
		Email: "",
	}
	id := primitive.NewObjectID()
	objectID := id.Hex()

	// Mock the PromoteUser method
	suite.mockUserUseCase.On("PromoteUser", mock.Anything, mock.Anything, mock.Anything).Return(expected_User, nil, http.StatusOK).Once()

	req, _ := http.NewRequest(http.MethodPost, "/users/"+objectID+"/promote", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: objectID},
	}
	c.Set("claim", user)

	// Act
	suite.userController.PromoteUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDemoteUser() {
	// Arrange
	user := &Domain.AccessClaims{
		ID:   suite.userID,
		Role: suite.Role,
	}

	expected_User := Domain.OmitedUser{
		ID:    suite.userID,
		Email: "",
	}
	id := primitive.NewObjectID()
	objectID := id.Hex()

	suite.mockUserUseCase.On("DemoteUser", mock.Anything, mock.Anything, mock.Anything).Return(expected_User, nil, http.StatusOK).Once()

	req, _ := http.NewRequest(http.MethodPost, "/users/"+objectID+"/demote", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: objectID},
	}
	c.Set("claim", user)

	// Act
	suite.userController.DemoteUser(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// assert.Equal(suite.T(), "{\"message\":\"User demoted successfully\", \"user\":null}", w.Body.String())
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
