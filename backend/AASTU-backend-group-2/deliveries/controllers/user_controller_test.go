package controllers

import (
	"blog_g2/domain"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"blog_g2/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerTestSuite struct {
	suite.Suite
	controller  *UserController
	mockUsecase *mocks.UserUsecase
}

func (suite *UserControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockUsecase = new(mocks.UserUsecase)
	suite.controller = NewUserController(suite.mockUsecase)
}

func (suite *UserControllerTestSuite) TestUpdateUserDetails_Success() {
	userID := primitive.NewObjectID()
	user := domain.User{ID: userID, UserName: "testuser", Email: "test@example.com"}
	suite.mockUsecase.On("UpdateUserDetails", mock.Anything, &user).Return(nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Set("userid", userID.Hex())
	context.Request = httptest.NewRequest(http.MethodPut, "/user/update", bytes.NewReader([]byte(`{"username":"testuser","email":"test@example.com"}`)))

	suite.controller.UpdateUserDetails(context)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	expectedResponse := `{"message":"User details updated successfully"}`
	assert.JSONEq(suite.T(), expectedResponse, recorder.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestUpdateUserDetails_BadRequest() {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Set("userid", "invalid")
	context.Request = httptest.NewRequest(http.MethodPut, "/user/update", bytes.NewReader([]byte(`{"username":"testuser"}`)))

	suite.controller.UpdateUserDetails(context)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
}

func (suite *UserControllerTestSuite) TestRegisterUser_Success() {
	// Define the expected user data
	expectedUser := domain.User{
		UserName: "testuser",
		Email:    "test@example.com",
		Password: "passwoRd123!",
	}

	// Set up the mock expectation
	suite.mockUsecase.On("RegisterUser", mock.Anything, mock.MatchedBy(func(user *domain.User) bool {
		return user.UserName == expectedUser.UserName &&
			user.Email == expectedUser.Email &&
			user.Password == expectedUser.Password
	})).Return(nil).Once()

	// Prepare the recorder and context
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader([]byte(`{"username":"testuser","email":"test@example.com","password":"passwoRd123!"}`)))

	// Call the RegisterUser method
	suite.controller.RegisterUser(context)

	// Assert the response status
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	assert.Contains(suite.T(), recorder.Body.String(), `"message":"User registered successfully"`)

	// Verify that mock expectations were met
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRegisterUser_BadRequest() {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader([]byte(`{"username":"testuser"}`)))

	suite.controller.RegisterUser(context)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
}

func (suite *UserControllerTestSuite) TestLoginUser_Success() {
	user := domain.User{Email: "test@example.com", Password: "password123"}
	suite.mockUsecase.On("LoginUser", mock.Anything, user).Return("mocked-token", nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader([]byte(`{"email":"test@example.com","password":"password123"}`)))

	suite.controller.LoginUser(context)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	expectedResponse := `{"message":"user logged in","token":"mocked-token"}`
	assert.JSONEq(suite.T(), expectedResponse, recorder.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLoginUser_BadRequest() {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader([]byte(`{"email":"test@example.com"}`)))

	suite.controller.LoginUser(context)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
}

func (suite *UserControllerTestSuite) TestForgotPassword_Success() {
	suite.mockUsecase.On("ForgotPassword", mock.Anything, "test@example.com").Return(nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/forgot-password", bytes.NewReader([]byte(`{"email":"test@example.com"}`)))

	suite.controller.ForgotPassword(context)

	assert.Equal(suite.T(), http.StatusAccepted, recorder.Code)
	expectedResponse := `{"message":"email succefully sent to the email provided"}`
	assert.JSONEq(suite.T(), expectedResponse, recorder.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

// func (suite *UserControllerTestSuite) TestForgotPassword_BadRequest() {
// 	recorder := httptest.NewRecorder()
// 	context, _ := gin.CreateTestContext(recorder)
// 	context.Request = httptest.NewRequest(http.MethodPost, "/user/forgot-password", bytes.NewReader([]byte(`{"email":"invalid"}`)))

// 	suite.controller.ForgotPassword(context)

// 	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
// }

func (suite *UserControllerTestSuite) TestResetPassword_Success() {
	suite.mockUsecase.On("ResetPassword", mock.Anything, "mock-token", "Abem@12hhshsd").Return(nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/reset-password?token=mock-token", bytes.NewReader([]byte(`{"password":"Abem@12hhshsd"}`)))

	suite.controller.ResetPassword(context)

	assert.Equal(suite.T(), http.StatusAccepted, recorder.Code)
	expectedResponse := `{"message":"Password has been reset successfully"}`
	assert.JSONEq(suite.T(), expectedResponse, recorder.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestResetPassword_BadRequest() {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/reset-password?token=mock-token", bytes.NewReader([]byte(`{"newPassword":""}`)))

	suite.controller.ResetPassword(context)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
}

func (suite *UserControllerTestSuite) TestLogoutUser_Success() {
	suite.mockUsecase.On("LogoutUser", mock.Anything, "mock-userid").Return(nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Set("userid", "mock-userid")
	context.Request = httptest.NewRequest(http.MethodPost, "/user/logout", nil)

	suite.controller.LogoutUser(context)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	expectedResponse := `{"message":"User logged out successfully"}`
	assert.JSONEq(suite.T(), expectedResponse, recorder.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteDemoteUser_Success() {
	suite.mockUsecase.On("PromoteDemoteUser", mock.Anything, "mock-id", true).Return(nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/promote-demote?id=mock-id&isadmin=true", nil)

	suite.controller.PromoteDemoteUser(context)

	assert.Equal(suite.T(), http.StatusAccepted, recorder.Code)
	expectedResponse := `{"message":"user admin privilege succesfully updated"}`
	assert.JSONEq(suite.T(), expectedResponse, recorder.Body.String())
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromoteDemoteUser_BadRequest() {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/promote-demote", nil)

	suite.controller.PromoteDemoteUser(context)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
