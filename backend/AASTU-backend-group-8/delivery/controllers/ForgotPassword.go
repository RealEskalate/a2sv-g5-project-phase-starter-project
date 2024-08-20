package controllers

import (
	"net/http"
	"meleket/domain"

	"github.com/gin-gonic/gin"
)

// type ForgotPasswordController struct {
// 	userUsecase domain.UserUsecaseInterface
// 	otpUsecase  domain.OTPUsecaseInterface
// }

// func NewForgotPasswordController(uu domain.UserUsecaseInterface, ou domain.OTPUsecaseInterface) *ForgotPasswordController {
// 	return &ForgotPasswordController{
// 		userUsecase: uu,
// 		otpUsecase:  ou,
// 	}
// }

func NewForgotPasswordController() *SignupController{
	return &SignupController{}
}

func (fp *SignupController) ForgotPassword(c *gin.Context){
	var email string
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	err := fp.otpUsecase.ForgotPassword(&email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to email"})
}

func (fp *SignupController) VerifyForgotOTP(c *gin.Context) {
	var otpRequest domain.OTPRequest
	if err := c.ShouldBindJSON(&otpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Verify the OTP
	existingOtp, err := fp.otpUsecase.VerifyOTP(otpRequest.Email, otpRequest.Otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := domain.AuthUser{
		Username: existingOtp.Username,
		Password: existingOtp.Password,
	}

	token, refreshToken, err := fp.userUsecase.Login(&user); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name":existingOtp.Username,"password": existingOtp.Password,"message": "User logged in successfully", "token": token, "refresh_token": refreshToken})
}