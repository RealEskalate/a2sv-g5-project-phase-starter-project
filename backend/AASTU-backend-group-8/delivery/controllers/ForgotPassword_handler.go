package controllers

import (
	"net/http"
	"meleket/domain"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	userUsecase domain.UserUsecaseInterface
	otpUsecase  domain.OTPUsecaseInterface
}

func NewForgotPasswordController(uu domain.UserUsecaseInterface, ou domain.OTPUsecaseInterface) *ForgotPasswordController {
	return &ForgotPasswordController{
		userUsecase: uu,
		otpUsecase:  ou,
	}
}

// func NewForgotPasswordController() *ForgotPasswordController{
// 	return &SignupController{}
// }

func (fp *ForgotPasswordController) ForgotPassword(c *gin.Context){
	email:= domain.Email{}
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	// panic(email)
	err := fp.otpUsecase.ForgotPassword(email.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error(),"err":"controller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to email"})
}

func (fp *ForgotPasswordController) VerifyForgotOTP(c *gin.Context) {
	var otpRequest domain.OTPResetPassword
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

	err = fp.userUsecase.UpdateUser(existingOtp.Username, otpRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error(), "err":"could not update the password"})
		return
	}

	authUser := domain.AuthUser{
		Username: existingOtp.Username,
		Password: otpRequest.Password,
	}

	token, refreshToken, err := fp.userUsecase.Login(&authUser); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name":existingOtp.Username,"password": existingOtp.Password,"message": "User logged in successfully", "token": token, "refresh_token": refreshToken})
}