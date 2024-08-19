package controllers

import (
	"errors"
	"net/http"
	"strings"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase interfaces.LogoutUsecase
	JwtService    interfaces.JwtService
}

func NewLogoutController(logoutUsecase interfaces.LogoutUsecase, jwtService interfaces.JwtService) *LogoutController {
	return &LogoutController{
		LogoutUsecase: logoutUsecase,
		JwtService:    jwtService,
	}
}

func (logoutController *LogoutController) Logout(ctx *gin.Context) {
	// get token from authorization header
	authHeader := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// validate token and get JwtCustom from the token
	JwtCustom, err := logoutController.JwtService.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	e := logoutController.LogoutUsecase.LogoutUser(ctx, JwtCustom.ID)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
