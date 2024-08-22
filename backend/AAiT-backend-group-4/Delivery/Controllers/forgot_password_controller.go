package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	OtpService            domain.OtpInfrastructure
	ForgotPasswordUsecase domain.ForgotPasswordUsecase
	Env                   *bootstrap.Env
}

func (fc *ForgotPasswordController) ForgotPassword(c *gin.Context) {
	var request map[string]string

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := fc.ForgotPasswordUsecase.ForgotPassword(c, request["email"], fc.Env.EmailApiKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (fc *ForgotPasswordController) ServePage(c *gin.Context) {
	c.HTML(http.StatusOK, "password_reset.html", nil)
}

func (fc *ForgotPasswordController) VerifyForgotPassowrd(c *gin.Context) {
	var passwordChangeForm domain.ForgotPasswordRequest

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found in context"})
		return
	}

	if err := c.ShouldBind(&passwordChangeForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailStr, ok := email.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email type assertion failed"})
		return
	}

	response, err := fc.ForgotPasswordUsecase.VerifyChangePassword(c, emailStr, passwordChangeForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
