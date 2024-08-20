package controllers

import (
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
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	e := logoutController.LogoutUsecase.LogoutUser(ctx, JwtCustom.ID)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
