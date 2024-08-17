package controllers

import (
	"net/http"
	"meleket/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func (uc *UserController) ForgotPassword(c *gin.Context){
// 	var email domain.Email
// 	if err := c.ShouldBindJSON(&email); err != nil {
// 		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
// 		return
// 	}

// 	err := uc.userUsecase.ForgotPassword(&email.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to email"})
// }
func (uc *UserController) SendOTP(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)

	user, err := uc.userUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	if err := uc.userUsecase.SendOTPEmail(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (uc *UserController) VerifyOTP(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	var otpRequest domain.OTPRequest

	if err := c.ShouldBindJSON(&otpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := uc.userUsecase.VerifyOTP(userID, otpRequest.OTP); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
