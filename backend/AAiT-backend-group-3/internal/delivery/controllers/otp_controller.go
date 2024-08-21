package controllers

import (
	"AAIT-backend-group-3/internal/domain/dtos"
	"AAIT-backend-group-3/internal/usecases"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type IOTPController interface {
	ForgotPassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

type OTPController struct {
	useCase usecases.IOtpUsecase
}

func NewOTPController(useCase usecases.IOtpUsecase) IOTPController {
	return &OTPController{
		useCase: useCase,
	}
}

func (c *OTPController) ForgotPassword(ctx *gin.Context) {
	var req dtos.ForgotPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.useCase.GenerateAndSendOtp(context.Background(), req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP and reset link sent to email"})
}

func (c *OTPController) ResetPassword(ctx *gin.Context) {
	otp := ctx.Query("otp")
	if otp == "" {
		ctx.JSON(400, gin.H{"error": "Missing token"})
		return
	}
	var req dtos.ResetPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	otpEntry, err := c.useCase.ValidateOtp(context.Background(), otp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("NewPassword", req.NewPassword)
	err = c.useCase.ResetPassword(context.Background(), otpEntry.UserID, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password has been successfully reset"})
}