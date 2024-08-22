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
)

type LoginControllerSuite struct {
	suite.Suite
	router          *gin.Engine
	LoginUsecase    *mocks.LoginUsecase
	LoginController *controller.LoginController
}

func (suite *LoginControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.LoginUsecase = new(mocks.LoginUsecase)
	env := &config.Env{}
	suite.LoginController = &controller.LoginController{
		LoginUsecase: suite.LoginUsecase,
		Env:          env,
	}
	suite.router = gin.Default()
	suite.router.POST("/login", suite.LoginController.Login)
	suite.router.POST("/refresh", suite.LoginController.RefreshTokenHandler)
}
func (suite *LoginControllerSuite) TearDownTest() {
	suite.LoginUsecase.AssertExpectations(suite.T())
}

func (suite *LoginControllerSuite) TestLogin() {
	suite.Run("login_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		login := domain.AuthLogin{
			Email:    "Test@gmail.com",
			Username: "test",
		}
		resp := domain.User{}
		accessToken := "token"
		suite.LoginUsecase.On("AuthenticateUser", mock.Anything, &login).Return(&resp, nil).Once()
		suite.LoginUsecase.On("CreateAccessToken", &resp, "", 0).Return(accessToken, nil).Once()
		suite.LoginUsecase.On("CreateRefreshToken", &resp, "", 0).Return("token", nil).Once()
		suite.LoginUsecase.On("SaveRefreshToken", mock.Anything, mock.Anything).Return(nil).Once()
		payload, _ := json.Marshal(login)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("login_user_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		login := domain.AuthLogin{
			Email:    "test@gmail.com",
			Username: "test",
		}
		resp := domain.User{}
		accessToken := "token"
		suite.LoginUsecase.On("AuthenticateUser", mock.Anything, &login).Return(&resp, nil).Once()
		suite.LoginUsecase.On("CreateAccessToken", &resp, "", 0).Return(accessToken, nil).Once()
		suite.LoginUsecase.On("CreateRefreshToken", &resp, "", 0).Return("token", nil).Once()
		suite.LoginUsecase.On("SaveRefreshToken", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		payload, _ := json.Marshal(login)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})

}

func TestLoginControllerSuite(t *testing.T) {
	suite.Run(t, new(LoginControllerSuite))
}
