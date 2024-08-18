package controllers

import (
	"errors"
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogoutController struct {
	LogoutUsecase interfaces.LogoutUsecase
	JwtService    interfaces.JwtService
}

func (logoutController *LogoutController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	claims, err := forgotPasswordController.JwtService.GetClaims(authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	userId, err := primitive.ObjectIDFromHex(claims.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid user id"))
		return
	}
}
