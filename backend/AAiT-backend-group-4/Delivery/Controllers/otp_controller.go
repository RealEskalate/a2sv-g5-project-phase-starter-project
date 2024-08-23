package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OtpController struct {
	OtpUsecase domain.OTPUsecase
	Env        bootstrap.Env
}

// VerifyOtp is a method of the OtpController struct that handles the verification of an OTP code.
// It takes a gin.Context object as a parameter and expects a JSON payload containing the OTP code and email.
// If the JSON payload is invalid, it returns a JSON response with a 400 Bad Request status and an error message.
// If the OTP verification fails, it returns a JSON response with a 400 Bad Request status and an error message.
// If the OTP verification is successful, it returns a JSON response with a 200 OK status and the verification response.
func (oc *OtpController) VerifyOtp(c *gin.Context) {
	var otpCode map[string]string

	err := c.ShouldBindJSON(&otpCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repsonse, err := oc.OtpUsecase.VerifyOTP(c, otpCode["email"], otpCode["otp"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repsonse)
}

// ResendOtp is a method of the OtpController struct that handles the request to resend an OTP.
// It takes a gin.Context object as a parameter and expects a JSON payload with a "email" field.
// If the JSON payload is valid, it calls the ResendOTP method of the OtpUsecase to resend the OTP to the specified email.
// If there is an error in binding the JSON payload or in resending the OTP, it returns a JSON response with the corresponding error message.
// If the OTP is successfully resent, it returns a JSON response with the success message.
func (oc *OtpController) ResendOtp(c *gin.Context) {
	var value map[string]string
	err := c.ShouldBindJSON(&value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repsonse, err := oc.OtpUsecase.ResendOTP(c, value["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repsonse)
}
