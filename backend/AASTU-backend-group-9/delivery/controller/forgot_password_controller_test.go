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

type ForgotPasswordControllerSuite struct {
	suite.Suite
	router          *gin.Engine
	ForgotPasswordUsecase    *mocks.ForgotPasswordUsecase
	ForgotPasswordController *controller.ForgotPasswordController
}

func (suite *ForgotPasswordControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.ForgotPasswordUsecase = new(mocks.ForgotPasswordUsecase)
	env := &config.Env{}
	suite.ForgotPasswordController = &controller.ForgotPasswordController{
		ForgotPasswordUsecase: suite.ForgotPasswordUsecase,
		Env:          env,
	}
	suite.router = gin.Default()
	suite.router.POST("/forgot_password", suite.ForgotPasswordController.ForgotPassword)
	suite.router.POST("/reset_password", suite.ForgotPasswordController.ResetPassword)
}
func (suite *ForgotPasswordControllerSuite) TearDownTest() {
	suite.ForgotPasswordUsecase.AssertExpectations(suite.T())
}

func (suite *ForgotPasswordControllerSuite) TestForgotPassword() {
	suite.Run("forgot_password_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.ForgotPasswordRequest{
			Email:    "test@gmail.com",
		}
		stmpusername := ""
		stmpassword := ""
		suite.ForgotPasswordUsecase.On("SendResetOTP", mock.Anything, user.Email,stmpusername,stmpassword).Return(nil).Once()
		payload, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/forgot_password", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("forgot_password_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		user := domain.ForgotPasswordRequest{
			Email:    "test@gmail.com",
		}
		stmpusername := ""
		stmpassword := ""
		suite.ForgotPasswordUsecase.On("SendResetOTP", mock.Anything, user.Email,stmpusername,stmpassword).Return(errors.New("error")).Once()
		payload, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/forgot_password", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func (suite *ForgotPasswordControllerSuite) TestResetPassword() {
	suite.Run("reset_password_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		var user struct {
			Email       string `json:"email"`
			OTPValue    string `json:"otp_value"`
			NewPassword string `json:"new_password"`
		}
		suite.ForgotPasswordUsecase.On("ResetPassword", mock.Anything, user.Email, user.OTPValue, user.NewPassword).Return(nil).Once()
		payload, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/reset_password", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("reset_password_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		var user struct {
			Email       string `json:"email"`
			OTPValue    string `json:"otp_value"`
			NewPassword string `json:"new_password"`
		}
		suite.ForgotPasswordUsecase.On("ResetPassword", mock.Anything, user.Email, user.OTPValue, user.NewPassword).Return(errors.New("error")).Once()
		payload, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/reset_password", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func TestForgotPasswordControllerSuite(t *testing.T) {
	suite.Run(t, new(ForgotPasswordControllerSuite))
}

