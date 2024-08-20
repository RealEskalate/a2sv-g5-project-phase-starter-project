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

func (oc *OtpController) VerifyOtp(c *gin.Context) {
	var otpCode map[string]string

	err := c.ShouldBindJSON(&otpCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userRequest := domain.UserOTPRequest{
		UserID: id,
		Email:  otpCode["email"],
	}

	repsonse, err := oc.OtpUsecase.VerifyOTP(c, &userRequest, otpCode["otp"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repsonse)
}

func (oc *OtpController) ResendOtp(c *gin.Context) {
	var value map[string]string
	err := c.ShouldBindJSON(&value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userRequest := domain.UserOTPRequest{
		UserID: id,
		Email:  value["email"],
	}

	repsonse, err := oc.OtpUsecase.GenerateOTP(c, &userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repsonse)
}
