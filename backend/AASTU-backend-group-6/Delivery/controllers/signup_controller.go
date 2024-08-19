package controllers

import (
	domain "blogs/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SignupController struct {
	SignupUsecase domain.SignupUseCase
}

func (s *SignupController) Signup(c *gin.Context) {
	var user domain.User
	var signupRequest domain.SignUpRequest
	err := c.BindJSON(&signupRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = copier.Copy(&user, &signupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map fields"})
		return
	}

	response := s.SignupUsecase.Create(c, user)
	HandleResponse(c, response)

}
func (s *SignupController) VerifyOTP(c *gin.Context) {
	var otp domain.OtpToken
	err := c.ShouldBindJSON(&otp)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := s.SignupUsecase.VerifyOTP(c, otp)
	HandleResponse(c, response)

}

func (s *SignupController) ForgotPassword(c *gin.Context) {
	var userEmail domain.ForgotPasswordRequest

	token := c.Query("token")

	if token == "" {

		err := c.ShouldBindJSON(&userEmail)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		response := s.SignupUsecase.ForgotPassword(c, userEmail)
		HandleResponse(c, response)
		return
	} else {
		var password domain.ResetPasswordRequest
		err := c.ShouldBindJSON(&password)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		response := s.SignupUsecase.ResetPassword(c, password, token)

		HandleResponse(c, response)

	}

}
