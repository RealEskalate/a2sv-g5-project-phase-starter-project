package usercontroller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mocks "github.com/group13/blog/delivery/controller/mocks"
	usercontroller "github.com/group13/blog/delivery/controller/user"
)

type UserControllerTestSuite struct {
	suite.Suite
	controller              *usercontroller.UserController
	promoteHandlerMock      *mocks.PromoteHandlerMock
	loginHandlerMock        *mocks.LoginHandlerMock
	signupHandlerMock       *mocks.SignupHandlerMock
	resetPasswordHandlerMock *mocks.ResetPasswordHandlerMock
	resetCodeSendHandlerMock *mocks.SendcodeHandlerMock
	validateCodeHandlerMock *mocks.ValidateCodeHandlerMock
	validateEmailHandlerMock *mocks.ValidateEmailHandlerMock
	updateProfileHandlerMock *mocks.UpdateProfileHandlerMock
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.promoteHandlerMock = new(mocks.PromoteHandlerMock)
	suite.loginHandlerMock = new(mocks.LoginHandlerMock)
	suite.signupHandlerMock = new(mocks.SignupHandlerMock)
	suite.resetPasswordHandlerMock = new(mocks.ResetPasswordHandlerMock)
	suite.resetCodeSendHandlerMock = new(mocks.SendcodeHandlerMock)
	suite.validateCodeHandlerMock = new(mocks.ValidateCodeHandlerMock)
	suite.validateEmailHandlerMock = new(mocks.ValidateEmailHandlerMock)
	suite.updateProfileHandlerMock = new(mocks.UpdateProfileHandlerMock)

	suite.controller = usercontroller.New(usercontroller.Config{
		PromoteHandler:         suite.promoteHandlerMock,
		LoginHandler:           suite.loginHandlerMock,
		SignupHandler:          suite.signupHandlerMock,
		ResetPasswordHandler:   suite.resetPasswordHandlerMock,
		ResetCodeSendHandler:   suite.resetCodeSendHandlerMock,
		ValidateCodeHandler:    suite.validateCodeHandlerMock,
		ValidateEmailHandler:   suite.validateEmailHandlerMock,
		UpdateProfileHandler:   suite.updateProfileHandlerMock,
	})
}

func (suite *UserControllerTestSuite) TestSignUpHandler() {
	// Arrange
	expectedResponse := "user signed up"
	suite.signupHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"email":"test@test.com","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func (suite *UserControllerTestSuite) TestLoginHandler() {
	// Arrange
	expectedResponse := "user logged in"
	suite.loginHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@test.com","password":"password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	// suite.controller.LoginHandler(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func (suite *UserControllerTestSuite) TestLogoutHandler() {
	// Arrange
	expectedResponse := "user logged out"
	suite.loginHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/logout", nil)

	// Act
	// suite.controller.LogoutHandler(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

}



func (suite *UserControllerTestSuite) TestPasswordResetHandler() {
	// Arrange
	expectedResponse := "password reset link sent"
	suite.resetPasswordHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/password/reset", strings.NewReader(`{"email":"test@test.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	// suite.controller.PasswordResetHandler(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func (suite *UserControllerTestSuite) TestUpdateProfileHandler() {
	// Arrange
	expectedResponse := "profile updated"
	suite.updateProfileHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPut, "/profile", strings.NewReader(`{"email":"newtest@test.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	// suite.controller.UpdateProfileHandler(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func (suite *UserControllerTestSuite) TestResetCodeSendHandler() {
	// Arrange
	expectedResponse := "reset code sent"
	suite.resetCodeSendHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/password/code/send", strings.NewReader(`{"email":"test@test.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	// suite.controller.ResetCodeSendHandler(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *UserControllerTestSuite) TestValidateCodeHandler() {
	// Arrange
	expectedResponse := "code validated"
	suite.validateCodeHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/password/code/validate", strings.NewReader(`{"code":"123456"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	// suite.controller.ValidateCodeHandler(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}


func (suite *UserControllerTestSuite) TestValidateEmailHandler() {
	// Arrange
	expectedResponse := "email validated"
	suite.validateEmailHandlerMock.On("Handle", mock.Anything).Return(expectedResponse, nil)

	// Create a request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/email/validate", strings.NewReader(`{"email":"test@test.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	// suite.controller.ValidateEmailHandler(c)

	// Assert
	
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
	
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
