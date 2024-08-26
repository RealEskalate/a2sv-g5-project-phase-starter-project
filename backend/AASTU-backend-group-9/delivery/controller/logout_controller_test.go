package controller_test

import (
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

type LogoutControllerSuite struct {
	suite.Suite
	router           *gin.Engine
	LogoutUsecase    *mocks.LogoutUsecase
	LogoutController *controller.LogoutController
}

func (suite *LogoutControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.LogoutUsecase = new(mocks.LogoutUsecase)
	suite.LogoutController = &controller.LogoutController{
		LogoutUsecase: suite.LogoutUsecase,
	}
	suite.router = gin.Default()
	suite.router.POST("/logout", suite.LogoutController.Logout)
}
func (suite *LogoutControllerSuite) TearDownTest() {
	suite.LogoutUsecase.AssertExpectations(suite.T())
}

func (suite *LogoutControllerSuite) TestLogout() {
	suite.Run("logout_user_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.LogoutRequest{}
		suite.LogoutUsecase.On("Logout", mock.Anything, user.RefreshToken,"cbe5cfdf7c2118a9c3d78ef1d684f3afa089201352886449a06a6511cfef74a7").Return(nil).Once()
		payload, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/logout", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("logout_user_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.LogoutRequest{}
		suite.LogoutUsecase.On("Logout", mock.Anything, user.RefreshToken,"cbe5cfdf7c2118a9c3d78ef1d684f3afa089201352886449a06a6511cfef74a7").Return(errors.New("error")).Once()
		payload, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/logout", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func TestLogoutControllerSuite(t *testing.T) {
	suite.Run(t, new(LogoutControllerSuite))
}
