package controllers

import (
	"errors"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ForgotPasswordController struct {
	ForgotPasswordUsecase interfaces.ForgotPasswordUsecase
	JwtService            interfaces.JwtService
}

func (forgotPasswordController *ForgotPasswordController) Login(c *gin.Context) {
	var request dtos.PasswordResetRequest

	resetURL, e := forgotPasswordController.ForgotPasswordUsecase.GenerateResetURL(c, request.Email)
	if e != nil {
		c.JSON(e.Code, e.Error())
	}

	e = forgotPasswordController.ForgotPasswordUsecase.SendResetEmail(c, request.Email, resetURL)
	if e != nil {
		c.JSON(e.Code, e.Error())
	}

	setUpPasswordRequest := &dtos.SetUpPasswordRequest{}

	err := c.ShouldBind(setUpPasswordRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid request"))
		return
	}

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

	setUpPasswordRequest.UserID = userId
	e = forgotPasswordController.ForgotPasswordUsecase.UpdateUserPassword(c, setUpPasswordRequest.Password, setUpPasswordRequest.Password)
	if e != nil {
		c.JSON(e.Code, e.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password reset, login again"})
}
