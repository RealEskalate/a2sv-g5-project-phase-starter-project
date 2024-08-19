package controllers

import (
	"errors"
	"net/http"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase interfaces.LoginUsecase
	Env          *config.Env
}

func NewLoginController(loginUsecase interfaces.LoginUsecase, env *config.Env) *LoginController {
	return &LoginController{
		LoginUsecase: loginUsecase,
		Env:          env,
	}
}

func (loginController *LoginController) Login(c *gin.Context) {
	var loginRequest dtos.LoginRequest

	// attempt to bind the json payload
	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	// envode Login_Usecase
	loginResponse, e := loginController.LoginUsecase.LoginUser(c, loginRequest)
	if e != nil {
		c.JSON(e.Code, e.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}
