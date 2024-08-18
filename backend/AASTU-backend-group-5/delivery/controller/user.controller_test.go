package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	controller *UserController
	usecase    *mocks.User_Usecase_interface
}

func (suite *UserControllerSuite) SetupTest() {
	usecase := new(mocks.User_Usecase_interface)
	suite.controller = &UserController{UserUsecase: usecase}
	suite.usecase = usecase
}

func (suite *UserControllerSuite) TestGetOneUser() {
	user := domain.ResponseUser{
		ID: primitive.NewObjectID().Hex(),
		UserName: "username",
		Email: "testuser@gmail.com",
		Is_Admin: false,
		Bio: "test bio",
		ProfilePicture: "pp",
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: user.ID}}

	suite.usecase.On("GetOneUser", user.ID).Return(user, nil)

	handler := suite.controller.GetOneUser()
	handler(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *UserControllerSuite) TestGetUsers() {
	users := []domain.ResponseUser{
		 {
			ID: primitive.NewObjectID().Hex(),
			UserName: "username",
			Email: "testuser@gmail.com",
			Is_Admin: false,
			Bio: "test bio",
			ProfilePicture: "pp",
		},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	suite.usecase.On("GetUsers").Return(users, nil)

	handler := suite.controller.GetUsers()
	handler(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *UserControllerSuite) TestUpdateUser() {
	id := primitive.NewObjectID()
	updateUser := domain.UpdateUser{
		UserName: "Updated username",
	}
	updatedUser := domain.ResponseUser{
		ID: id.Hex(),
		UserName: "Updated username",
		Email: "testuser@gmail.com",
		Is_Admin: false,
		Bio: "test bio",
		ProfilePicture: "pp",
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"username": "Updated username"
	}`)

	c.Request = httptest.NewRequest(http.MethodPut, "/user/"+id.Hex(), body)
	c.Params = gin.Params{gin.Param{Key: "_id", Value: id.Hex()}}
	c.Request.Header.Set("Content-Type", "application/json")


	suite.usecase.On("UpdateUser", "", updateUser).Return(updatedUser, nil)

	handler := suite.controller.UpdateUser()
	handler(c)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *UserControllerSuite) TestDeleteUser() {
	id := primitive.NewObjectID().Hex()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: id}}

	suite.usecase.On("DeleteUser", id).Return(nil)

	handler := suite.controller.DeleteUser()
	handler(c)

	suite.Equal(http.StatusAccepted, w.Code)
	suite.JSONEq(`{"message": "accepted!"}`, w.Body.String())
}

func (suite *UserControllerSuite) TestLogIn() {
	loginModel := domain.LogINUser{
		UserName: "username",
		Email: "test@gmail.com",
		Password: "password"}

	user := domain.ResponseUser{
		ID: primitive.NewObjectID().Hex(),
		UserName: "username",
		Email: "testuser@gmail.com",
		Is_Admin: false,
		Bio: "test bio",
		ProfilePicture: "pp",
	}

	body := bytes.NewBufferString(`{
		"username": "username",
		"email":"test@gmail.com",
		"password" : "password"
	}`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/login", body)
	c.Request.Header.Set("Content-Type", "application/json")

	suite.usecase.On("LogIn", loginModel).Return(user, nil)

	handler := suite.controller.LogIn()
	handler(c)

	suite.Equal(http.StatusOK, w.Code , w.Body.String())
}

func (suite *UserControllerSuite) TestRegister() {
	registerUser := domain.RegisterUser{
		UserName: "username",
		Bio: "test bio",
		ProfilePicture: "pp",
		Email: "test@gmail.com",	
		Password: "password",
	}

	user := domain.ResponseUser{
		ID: primitive.NewObjectID().Hex(),
		UserName: "username",
		Email: "testuser@gmail.com",
		Is_Admin: false,
		Bio: "test bio",
		ProfilePicture: "pp",
	}

	body := bytes.NewBufferString(`{
		"username": "username",
		"email": "test@gmail.com",
		"password": "password",
		"bio": "test bio",
		"profile_picture": "pp"
	}`)
	

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/register/", body)
	c.Request.Header.Set("Content-Type", "application/json")

	suite.usecase.On("Register", registerUser).Return(user, nil)

	handler := suite.controller.Register()
	handler(c)

	suite.Equal(http.StatusOK, w.Code , w.Body.String())
}

func (suite *UserControllerSuite) TestFilterUser() {
	filter := map[string]string{"email": "test@gmail.com"}
	users := []domain.ResponseUser{
		{
			ID: primitive.NewObjectID().Hex(),
			UserName: "username",
			Email: "testuser@gmail.com",
			Is_Admin: false,
			Bio: "test bio",
			ProfilePicture: "pp",
		},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/users?email=test@gmail.com", nil)

	suite.usecase.On("FilterUser", filter).Return(users, nil)

	handler := suite.controller.FilterUser()
	handler(c)

	suite.Equal(http.StatusOK, w.Code)
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
