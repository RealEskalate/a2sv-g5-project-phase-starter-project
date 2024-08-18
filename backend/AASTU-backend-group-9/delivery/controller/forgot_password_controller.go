package controller

import (
    "blog/domain"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
    ForgotPasswordUsecase domain.ForgotPasswordUsecase
}

func (fpc *ForgotPasswordController) ForgotPassword(c *gin.Context) {
    var request struct {
        Email string `json:"email"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Pass SMTP credentials directly to the usecase
    smtpUsername := "your_smtp_username" // Replace with actual SMTP username
    smtpPassword := "your_smtp_password" // Replace with actual SMTP password
    err := fpc.ForgotPasswordUsecase.SendResetOTP(context.Background(), request.Email, smtpUsername, smtpPassword)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}


func (fpc *ForgotPasswordController) ResetPassword(c *gin.Context) {
    var request struct {
        Email       string `json:"email"`
        OTPValue    string `json:"otp_value"`
        NewPassword string `json:"new_password"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    err := fpc.ForgotPasswordUsecase.ResetPassword(context.Background(), request.Email, request.OTPValue, request.NewPassword)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
