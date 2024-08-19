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
	var request dtos.LoginRequest

	// attempt to bind the json payload
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	// envode Login_Usecase
	user, e := loginController.LoginUsecase.LoginUser(c, request.UsernameOrEmail, request.Password)
	if e != nil {
		c.JSON(e.Code, e.Error())
		return
	}

	accessTokenExp := loginController.Env.ACCESS_TOKEN_EXPIRY_HOUR
	refreshTokenExp := loginController.Env.REFRESH_TOKEN_EXPIRY_HOUR

	// generate access token
	accessToken, e := loginController.LoginUsecase.GenerateAccessToken(user, accessTokenExp)
	if e != nil {
		c.JSON(e.Code, e.Error())
		return
	}

	// generate refresh token
	refreshToken, e := loginController.LoginUsecase.GenerateRefreshToken(user, refreshTokenExp)
	if e != nil {
		c.JSON(e.Code, e.Error())
		return
	}

	loginResponse := dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
