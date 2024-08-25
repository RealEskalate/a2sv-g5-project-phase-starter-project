package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "Blog_Starter/utils"
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "bou.ke/monkey"
)

type LoginControllerTestSuite struct {
    suite.Suite
    loginController *LoginController
    mockLoginUsecase *mocks.LoginUsecase
    mockUserUsecase  *mocks.UserUsecase
    mockOtpUsecase   *mocks.OtpUsecase
}

func (suite *LoginControllerTestSuite) SetupTest() {
    suite.mockLoginUsecase = new(mocks.LoginUsecase)
    suite.mockUserUsecase = new(mocks.UserUsecase)
    suite.mockOtpUsecase = new(mocks.OtpUsecase)
    suite.loginController = NewLoginController(suite.mockLoginUsecase, suite.mockOtpUsecase, suite.mockUserUsecase)
}

func (suite *LoginControllerTestSuite) TestLogin_Success() {
    userLogin := domain.UserLogin{
        Email:    "test@example.com",
        Password: "password",
    }

    loginResponse := &domain.LoginResponse{
        UserID:       "test_user_id",
		AccessToken:  "access_token",
        RefreshToken: "refresh_token",
    }

    suite.mockLoginUsecase.On("Login", mock.Anything, &userLogin).Return(loginResponse, nil)

    body, _ := json.Marshal(userLogin)
    req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/login", suite.loginController.Login)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusOK, rr.Code)
}

func (suite *LoginControllerTestSuite) TestLogin_InvalidRequest() {
    req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(`{}`)))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/login", suite.loginController.Login)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
}

func (suite *LoginControllerTestSuite) TestLogin_UserNotFound() {
    userLogin := domain.UserLogin{
        Email:    "test@example.com",
        Password: "password",
    }

    suite.mockLoginUsecase.On("Login", mock.Anything, &userLogin).Return(nil, errors.New("mongo: no documents in result"))

    body, _ := json.Marshal(userLogin)
    req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/login", suite.loginController.Login)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusNotFound, rr.Code)
} 

func (suite *LoginControllerTestSuite) TestForgotPassword_Success() {
    forgotPasswordRequest := domain.ForgotPasswordRequest{
        Email: "test@example.com",
    }

    user := &domain.UserResponse{
        Email:       "test@example.com",
        IsActivated: true,
    }

    suite.mockUserUsecase.On("GetUserByEmail", mock.Anything, forgotPasswordRequest.Email).Return(user, nil)
    suite.mockOtpUsecase.On("GetOtpByEmail", mock.Anything, forgotPasswordRequest.Email).Return(domain.Otp{}, errors.New("not found"))
    suite.mockOtpUsecase.On("SaveOtp", mock.Anything, mock.Anything).Return(nil)

    // Monkey patch the SendTestEmail function
    monkey.Patch(utils.SendTestEmail, func(email, subject, body string) error {
        return nil
    })
    defer monkey.Unpatch(utils.SendTestEmail)

    body, _ := json.Marshal(forgotPasswordRequest)
    req, err := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/forgot-password", suite.loginController.ForgotPassword)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusOK, rr.Code)
}

func (suite *LoginControllerTestSuite) TestUpdatePassword_Success() {
    changePasswordRequest := domain.ChangePasswordRequest{
        OTP:      "1234",
        Email:    "test@example.com",
        Password: "new_passwordA#1528",
    }

    user := &domain.UserResponse{
        Email: "test@example.com",
    }

    otp := domain.Otp{
        Otp:        "1234",
        Expiration: time.Now().Add(5 * time.Minute),
    }

    suite.mockUserUsecase.On("GetUserByEmail", mock.Anything, changePasswordRequest.Email).Return(user, nil)
    suite.mockOtpUsecase.On("GetOtpByEmail", mock.Anything, changePasswordRequest.Email).Return(otp, nil)
    suite.mockLoginUsecase.On("UpdatePassword", mock.Anything, changePasswordRequest, user.UserID.Hex()).Return(nil)

    body, _ := json.Marshal(changePasswordRequest)
    req, err := http.NewRequest(http.MethodPost, "/update-password", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/update-password", suite.loginController.UpdatePassword)
    router.ServeHTTP(rr, req)
    assert.Equal(suite.T(), http.StatusOK, rr.Code)
}

func (suite *LoginControllerTestSuite) TestLogOut_Success() {
	userID := primitive.NewObjectID()
    user := &domain.AuthenticatedUser{
        UserID: userID.Hex(),
    }

    monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
        return user, nil
    })
    defer monkey.Unpatch(utils.CheckUser)

    suite.mockLoginUsecase.On("LogOut", mock.AnythingOfType("*gin.Context"), user.UserID).Return(nil)

    req, err := http.NewRequest(http.MethodPost, "/logout", nil)
    assert.NoError(suite.T(), err)

    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/logout", suite.loginController.LogOut)
    router.ServeHTTP(rr, req)
	
    assert.Equal(suite.T(), http.StatusOK, rr.Code)
}

func TestLoginControllerTestSuite(t *testing.T) {
    suite.Run(t, new(LoginControllerTestSuite))
}