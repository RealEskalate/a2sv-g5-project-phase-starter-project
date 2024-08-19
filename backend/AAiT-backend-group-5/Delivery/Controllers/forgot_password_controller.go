package controllers

import (
	"errors"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	PasswordUsecase interfaces.PasswordUsecase
}

func NewForgotPasswordController(PasswordUsecase interfaces.PasswordUsecase) *ForgotPasswordController {
	return &ForgotPasswordController{
		PasswordUsecase: PasswordUsecase,
	}
}

func (forgotPasswordController *ForgotPasswordController) ForgotPasswordRequest(ctx *gin.Context) {
	var request dtos.PasswordResetRequest

	// attempt to bind json payload
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// generate URL to be sent via email
	resetURL, e := forgotPasswordController.PasswordUsecase.GenerateResetURL(ctx, request.Email)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	// send confirmation email
	e = forgotPasswordController.PasswordUsecase.SendResetEmail(ctx, request.Email, resetURL)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "confirmation email sent"})
}

func (forgotPasswordController *ForgotPasswordController) SetNewPassword(ctx *gin.Context) {
	var setUpPasswordRequest *dtos.SetUpPasswordRequest

	// attempt to bind the payload carrying the new password
	err := ctx.ShouldBind(&setUpPasswordRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	// get short code from the URL
	shortURLCode := ctx.Param("id")

	e := forgotPasswordController.PasswordUsecase.UpdateUserPassword(ctx, setUpPasswordRequest.Password, shortURLCode)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "password reset, login again"})
}
