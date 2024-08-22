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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	router         *gin.Engine
	UserUsecase    *mocks.UserUsecase
	UserController *controller.UserController
}

func (suite *UserControllerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.UserUsecase = new(mocks.UserUsecase)
	env := &config.Env{}
	suite.UserController = &controller.UserController{
		UserUsecase: suite.UserUsecase,
		Env:         env,
	}
	suite.router = gin.Default()
	suite.router.POST("/create_user", suite.UserController.CreateUser)
	suite.router.PATCH("/update_user/:id", suite.UserController.UpdateUser)
	suite.router.DELETE("/delete_user/:id", suite.UserController.DeleteUser)
	suite.router.GET("/get_user/:id", suite.UserController.GetUser)
	suite.router.GET("/get_all_users", suite.UserController.GetUsers)
	suite.router.PATCH("/promote_user/:id", suite.UserController.PromoteUser)
	suite.router.PATCH("/demote_user/:id", suite.UserController.DemoteUser)
}

func (suite *UserControllerSuite) TearDownTest() {
	suite.UserUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestCreateUser() {
	suite.Run("create_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.CreateUser{
			Email:    "non-existing-email@gmail.com",
			Username: "non-existing-username",
			Password: "password",
			Role:     "user",
		}
		claims := &domain.JwtCustomClaims{
			UserID:   primitive.NewObjectID(),
			Username: "test",
		}
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.UserUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(nil, nil).Once()
		suite.UserUsecase.On("CreateUser", mock.Anything, &user, claims).Return(nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/create_user", bytes.NewBuffer(payload))
		c.Set("claim", *claims) // Set claims before calling the handler
		suite.UserController.CreateUser(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(gin.H{"message": "User created successfully"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("Email_already_exists", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.CreateUser{
			Email:    "existing-email@gmail.com",
			Username: "non-existing-username",
			Password: "password",
			Role:     "user",
		}
		existingUser := &domain.User{}
		claims := &domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(existingUser, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/create_user", bytes.NewBuffer(payload))
		c.Set("claim", *claims) // Set claims before calling the handler
		suite.UserController.CreateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "Email already exists"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("Username_already_exists", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.CreateUser{
			Email:    "non-existing-email@gmail.com",
			Username: "existing-username",
			Password: "password",
			Role:     "user",
		}
		existingUser := &domain.User{}
		claims := &domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.UserUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(existingUser, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/create_user", bytes.NewBuffer(payload))
		c.Set("claim", *claims) // Set claims before calling the handler
		suite.UserController.CreateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "Username already exists"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("create_user_failure", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.CreateUser{
			Email:    "non-existing-email@gmail.com",
			Username: "non-existing-username",
			Password: "password",
			Role:     "user",
		}
		claims := &domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.UserUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(nil, nil).Once()
		suite.UserUsecase.On("CreateUser", mock.Anything, &user, claims).Return(errors.New("Internal Server Error")).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPost, "/create_user", bytes.NewBuffer(payload))
		c.Set("claim", *claims) // Set claims before calling the handler
		suite.UserController.CreateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "Internal server error"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func (suite *UserControllerSuite) TestUpdateUser() {
	suite.Run("update_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.User{
			Email:    "non-existing-email@gmail.com",
			Username: "non-existing-username",
			Password: "password",
			Role:     "user",
		}
		existinguser := domain.User{}
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existinguser, nil).Once()
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.UserUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(nil, nil).Once()
		suite.UserUsecase.On("UpdateUser", mock.Anything, &user, &claims, &existinguser).Return(&user, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/update_user/"+string(id.Hex()), bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.UserController.UpdateUser(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(gin.H{"message": "User updated successfully", "data": user})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
	suite.Run("user_not_found", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.User{
			Email:    "test@gmail.com",
			Username: "test",
			Password: "password",
			Role:     "user",
		}
		claims := domain.JwtCustomClaims{}
		id := primitive.NewObjectID()
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/update_user/"+string(id.Hex()), bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.UserController.UpdateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "User not found"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
	suite.Run("email_already_exists", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.User{
			Email:    "existing-email@gmail.com",
			Username: "non-existing-username",
			Password: "password",
			Role:     "user",
		}
		existingUser := domain.User{}
		claims := domain.JwtCustomClaims{}
		id := primitive.NewObjectID()
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existingUser, nil).Once()
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(&existingUser, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/update_user/"+string(id.Hex()), bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.UserController.UpdateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "Email already exists"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
	suite.Run("username_already_exists", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.User{
			Email:    "non-existing-email@gmail.com",
			Username: "existing-username",
			Password: "password",
			Role:     "user",
		}
		existingUser := domain.User{}
		claims := domain.JwtCustomClaims{}
		id := primitive.NewObjectID()
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existingUser, nil).Once()
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.UserUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(&existingUser, nil).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/update_user/"+string(id.Hex()), bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.UserController.UpdateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "Username already exists"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
	suite.Run("update_user_failure", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.User{
			Email:    "non-existing-email@gmail.com",
			Username: "non-existing-username",
			Password: "password",
			Role:     "user",
		}
		existingUser := domain.User{}
		claims := domain.JwtCustomClaims{}
		id := primitive.NewObjectID()
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existingUser, nil).Once()
		suite.UserUsecase.On("GetUserByEmail", mock.Anything, user.Email).Return(nil, nil).Once()
		suite.UserUsecase.On("GetUserByUsername", mock.Anything, user.Username).Return(nil, nil).Once()
		suite.UserUsecase.On("UpdateUser", mock.Anything, &user, &claims, &existingUser).Return(nil, errors.New("Internal Server Error")).Once()
		payload, _ := json.Marshal(user)
		c.Request = httptest.NewRequest(http.MethodPatch, "/update_user/"+string(id.Hex()), bytes.NewBuffer(payload))
		c.Set("claim", claims)
		suite.UserController.UpdateUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": errors.New("Internal Server Error").Error()})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func (suite *UserControllerSuite) TestDeleteUser() {
	suite.Run("delete_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		existinguser := domain.User{}
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existinguser, nil).Once()
		suite.UserUsecase.On("DeleteUser", mock.Anything, mock.Anything, &claims).Return(nil).Once()
		c.Request = httptest.NewRequest(http.MethodDelete, "/delete_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.DeleteUser(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(gin.H{"message": "User deleted successfully"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("user_not_found", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
		c.Request = httptest.NewRequest(http.MethodDelete, "/delete_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.DeleteUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "User not found"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("delete_user_failure", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		existingUser := domain.User{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existingUser, nil).Once()
		suite.UserUsecase.On("DeleteUser", mock.Anything, mock.Anything, &claims).Return(errors.New("Internal Server Error")).Once()
		c.Request = httptest.NewRequest(http.MethodDelete, "/delete_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.DeleteUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "Internal Server Error"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func (suite *UserControllerSuite) TestGetUser() {
	suite.Run("get_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		existingUser := domain.User{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&existingUser, nil).Once()
		c.Request = httptest.NewRequest(http.MethodGet, "/get_user/"+string(id.Hex()), nil)
		suite.UserController.GetUser(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(existingUser)
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("user_not_found", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, errors.New("Internal Server Error")).Once()
		c.Request = httptest.NewRequest(http.MethodGet, "/get_user/"+string(id.Hex()), nil)
		suite.UserController.GetUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "User not found"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func (suite *UserControllerSuite) TestGetUsers() {
	suite.Run("get_all_users_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		users := []*domain.User{
			{
				Email: "test@gmail.com",
			},
		}
		suite.UserUsecase.On("GetAllUsers", mock.Anything).Return(users, nil).Once()
		c.Request = httptest.NewRequest(http.MethodGet, "/get_all_users", nil)
		suite.UserController.GetUsers(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(users)
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("get_all_users_failure", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		suite.UserUsecase.On("GetAllUsers", mock.Anything).Return(nil, errors.New("Internal Server Error")).Once()
		c.Request = httptest.NewRequest(http.MethodGet, "/get_all_users", nil)
		suite.UserController.GetUsers(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "No users found"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func (suite *UserControllerSuite) TestPromoteUser() {
	suite.Run("promote_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		user := domain.User{
			Email:    "test@gmail.com",
			Username: "test",
			Password: "password",
			Role:     "user",
		}
		
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(&user, nil).Once()
		suite.UserUsecase.On("PromoteUser", mock.Anything, user.ID, &claims).Return(nil).Once()
		c.Request = httptest.NewRequest(http.MethodPatch, "/promote_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.PromoteUser(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(gin.H{"message": "User promoted successfully"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("promote_user_failure", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, errors.New("Internal Server Error")).Once()
		c.Request = httptest.NewRequest(http.MethodPatch, "/promote_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.PromoteUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "User not found"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}

func (suite *UserControllerSuite) TestDemoteUser() {
	suite.Run("demote_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.User{
			Email:    "test@gmail.com",
			Username: "test",
			Password: "password",
			Role:     "admin",
		}
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, user.ID).Return(&user, nil).Once()
		suite.UserUsecase.On("DemoteUser", mock.Anything, user.ID, &claims).Return(nil).Once()
		c.Request = httptest.NewRequest(http.MethodPatch, "/demote_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.DemoteUser(c)
		suite.Equal(200, w.Code)
		expect, err := json.Marshal(gin.H{"message": "User demoted successfully"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})

	suite.Run("demote_user_failure", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := primitive.NewObjectID()
		claims := domain.JwtCustomClaims{}
		suite.UserUsecase.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, errors.New("Internal Server Error")).Once()
		c.Request = httptest.NewRequest(http.MethodPatch, "/demote_user/"+string(id.Hex()), nil)
		c.Set("claim", claims)
		suite.UserController.DemoteUser(c)
		suite.Equal(400, w.Code)
		expect, err := json.Marshal(gin.H{"error": "User not found"})
		suite.Nil(err)
		suite.Equal(expect, w.Body.Bytes())
	})
}
func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
