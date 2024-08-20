package controllers

import (
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	userUsecase domain.UserUsecaseInterface
	otpUsecase  domain.OTPUsecaseInterface
}

func NewSignupController(usercase domain.UserUsecaseInterface, otpcase domain.OTPUsecaseInterface) *SignupController {
	return &SignupController{
		userUsecase: usercase,
		otpUsecase:  otpcase,
	}
}

// Sends an OTP before registering
func (sc *SignupController) Signup(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if email already exists
	existingUser, _ := sc.userUsecase.GetUserByEmail(&user.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Check if username already exists
	existingUser, _ = sc.userUsecase.GetUserByUsername(&user.Name)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Generate and send OTP
	err := sc.otpUsecase.GenerateAndSendOTP(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err, "error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// VerifyOTP verifies the OTP and completes the signup process
func (sc *SignupController) VerifyOTP(c *gin.Context) {
	var otpRequest domain.OTPRequest
	if err := c.ShouldBindJSON(&otpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Verify the OTP
	existingOtp, err := sc.otpUsecase.VerifyOTP(otpRequest.Email, otpRequest.Otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error1"})
		return
	}

	// Get the user
	user := domain.User{
		Name:     existingOtp.Username,
		Email:    existingOtp.Email,
		Password: existingOtp.Password,
		Role:     existingOtp.Role,
	}
	err = sc.userUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
