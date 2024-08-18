package controllers

import (
	"errors"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	utils "github.com/aait.backend.g5.main/backend/Utils"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase interfaces.LoginUsecase
	Env          *utils.Env
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

	accessTokenExp := loginController.Env.AccessTokenExpiryHour
	refreshTokenExp := loginController.Env.RefreshTokenExpiryHour

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
