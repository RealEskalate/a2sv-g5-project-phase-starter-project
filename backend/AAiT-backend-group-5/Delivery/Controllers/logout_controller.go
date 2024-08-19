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

func (logoutController *LogoutController) Logout(ctx *gin.Context) {
	// get claims from authorization header
	authHeader := ctx.GetHeader("Authorization")

	claims, err := logoutController.JwtService.GetClaims(authHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid token"))
		return
	}

	// get userId from claims
	userId, err := primitive.ObjectIDFromHex(claims.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	e := logoutController.LogoutUsecase.LogoutUser(ctx, userId.Hex())
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
