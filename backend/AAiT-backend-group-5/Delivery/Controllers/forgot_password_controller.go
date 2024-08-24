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

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := request.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "one or more fields are missing"})
		return
	}

	agent := ctx.Request.UserAgent()
	resetURL, e := forgotPasswordController.PasswordUsecase.GenerateResetURL(ctx, request.Email, agent)
	if e != nil {
		ctx.IndentedJSON(e.Code, e.Error())
		return
	}

	e = forgotPasswordController.PasswordUsecase.SendResetEmail(ctx, request.Email, resetURL)
	if e != nil {
		ctx.IndentedJSON(e.Code, e.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "confirmation email sent"})
}

func (forgotPasswordController *ForgotPasswordController) SetNewPassword(ctx *gin.Context) {
	var setUpPasswordRequest dtos.SetUpPasswordRequest

	err := ctx.ShouldBind(&setUpPasswordRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := setUpPasswordRequest.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "one or more fields are missing"})
		return
	}

	shortURLCode := ctx.Param("id")

	e := forgotPasswordController.PasswordUsecase.SetUpdateUserPassword(ctx, shortURLCode, setUpPasswordRequest.Password)
	if e != nil {
		ctx.IndentedJSON(e.Code, e.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "password reset, login again"})
}
