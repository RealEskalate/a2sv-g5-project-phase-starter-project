package controller_test

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	userUsecase    *mocks.UserUsecase
	userController *controller.UserController
	router         *gin.Engine
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.userUsecase = new(mocks.UserUsecase)
	suite.userController = &controller.UserController{}

	suite.router = gin.Default()
	suite.router.POST("/register", suite.userController.Register)
	suite.router.PUT("/update/:username", suite.userController.UpdateUser)
	suite.router.POST("/login", suite.userController.Login)
	suite.router.POST("/refresh", suite.userController.RefreshToken)
	suite.router.POST("/forgot-password", suite.userController.ForgotPassword)
	suite.router.GET("/reset/:token", suite.userController.ResetPassword)
	suite.router.GET("/verify/:token", suite.userController.Verify)
	suite.router.DELETE("/delete/:username", suite.userController.DeleteUser)
	suite.router.PUT("/promote/:username", suite.userController.PromoteToAdmin)
	suite.router.PUT("/change_password", suite.userController.ChangePassword)
	suite.router.POST("/logout", suite.userController.Logout)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestRegister() {

	input := Domain.RegisterInput{
		Name:     "Hamza",
		Email:    "hajihamza172@gmail.com",
		Username: "hajihamza172",
		Password: "password123",
	}
	suite.userUsecase.On("Register", input).Return(&Domain.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}, nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
	suite.userUsecase.AssertExpectations(suite.T())

}

func (suite *UserControllerTestSuite) TestUpdateUser() {
	Username := "hamza"
	input := Domain.UpdateUserInput{
		Username: Username,
		Password: "1234",
	}

	// Mock the user usecase
	suite.userUsecase.On("UpdateUser", Username, &input).Return(nil)

	// Create a new recorder and context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Manually set the username  in the Gin context
	c.Set("username", "hamza")

	// Set the request params and body
	c.Params = []gin.Param{{Key: "username", Value: Username}}
	c.Request = httptest.NewRequest("PUT", "/update/"+Username, bytes.NewBuffer([]byte(`{
        "Username": "hamza",
        "Password": "1234"
    }`)))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the update user handler
	suite.userController.UpdateUser(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDeleteUser() {
	username := "hamza"
	suite.userUsecase.On("DeleteUser", username).Return(nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "username", Value: username}}
	c.Set("username", "hamza")
	c.Set("role", "admin")

	suite.userController.DeleteUser(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "{\"message\":\"User deleted successfully\"}")

}

func (suite *UserControllerTestSuite) TestLogin() {
	username := "hamza"
	input := &Domain.LoginInput{
		Username: username,
		Password: "1234",
	}
	dummyToken := "token"
	suite.userUsecase.On("Login", mock.AnythingOfType("*gin.Context"), input).Return(dummyToken, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{
        "username": "hamza",
        "password": "1234"
    }`)))
	c.Request.Header.Set("Content-Type", "application/json")
	suite.userController.Login(c)
	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `{"access_token":"`+dummyToken+`"}`)

}

func (suite *UserControllerTestSuite) TestForgotPassword() {
	username := "H@mza101112"
	email := "hajihamza172@gmail.com"
	input := Domain.ForgetPasswordInput{
		Username: username,
		Email:    email,
	}

	dummyToken := "token"
	suite.userUsecase.On("ForgotPassword", input.Username).Return(dummyToken, nil)

	// Marshal the input struct into JSON
	inputJSON, err := json.Marshal(input)
	if err != nil {
		suite.T().Fatalf("Failed to marshal input: %v", err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/forgot-password", bytes.NewReader(inputJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.userController.ForgotPassword(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"token":"`+dummyToken+`"}`, w.Body.String())
}

func (suite *UserControllerTestSuite) TestResetPassword() {
	resetToken := "dummyResetToken"
	newAccessToken := "newAccessToken"

	suite.userUsecase.On("Reset", resetToken).Return(newAccessToken, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "token", Value: resetToken},
	}
	c.Request = httptest.NewRequest("POST", "/reset-password/"+resetToken, nil)
	suite.userController.ResetPassword(c)

	// Assertions
	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"access_token":"`+newAccessToken+`"}`, w.Body.String())

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestChangePassword() {
	username := "Hamza10@"
	newPassword := "Hamza123$"

	input := Domain.ChangePasswordInput{
		NewPassword: newPassword,
	}

	// Set up the mock to expect the UpdatePassword method
	suite.userUsecase.On("UpdatePassword", username, newPassword).Return(nil)

	inputJSON, err := json.Marshal(input)
	if err != nil {
		suite.T().Fatalf("Failed to marshal input: %v", err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("PUT", "/change-password", bytes.NewReader(inputJSON))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.userController.ChangePassword(c)

	// Check the response
	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "{\"message\":\"Password changed successfully\"}")

	// Assert that the mock expectations were met
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogout() {
	token := "validTokenString"

	suite.userUsecase.On("Logout", token).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/logout", nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)

	suite.userController.Logout(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "{\"message\":\"User logged out successfully\"}")

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteToAdmin() {
	username := "hamza10@"

	suite.userUsecase.On("PromoteTOAdmin", username).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "username", Value: username}}
	c.Request = httptest.NewRequest("POST", "/promote-to-admin/"+username, nil)

	suite.userController.PromoteToAdmin(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "{\"message\":\"User promoted to admin successfully\"}")

	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestVerify() {
	token := "validTokenString"

	// Set up the mock
	suite.userUsecase.On("Verify", token).Return(nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/verify", nil)
	c.Params = append(c.Params, gin.Param{Key: "token", Value: token})

	suite.userController.Verify(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "{\"message\":\"Email verified successfully\"}")

	suite.userUsecase.AssertExpectations(suite.T())
}
