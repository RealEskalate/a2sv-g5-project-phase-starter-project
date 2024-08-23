package controllers

import (
	"net/http"

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
	userId := ctx.GetString("id")

	e := logoutController.LogoutUsecase.LogoutUser(ctx, userId)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
