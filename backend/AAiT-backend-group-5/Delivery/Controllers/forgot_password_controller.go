package controllers

import (
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

	// attempt to bind IndentedJSON payload
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// generate URL to be sent via email
	resetURL, e := forgotPasswordController.PasswordUsecase.GenerateResetURL(ctx, request.Email)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	// send confirmation email
	e = forgotPasswordController.PasswordUsecase.SendResetEmail(ctx, request.Email, resetURL)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "confirmation email sent"})
}

func (forgotPasswordController *ForgotPasswordController) SetNewPassword(ctx *gin.Context) {
	var setUpPasswordRequest *dtos.SetUpPasswordRequest

	// attempt to bind the payload carrying the new password
	err := ctx.ShouldBind(&setUpPasswordRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// get short code from the URL
	shortURLCode := ctx.Param("id")

	e := forgotPasswordController.PasswordUsecase.SetUpdateUserPassword(ctx, shortURLCode, setUpPasswordRequest.Password)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "password reset, login again"})
}
