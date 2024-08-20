package controller

import (
	"AAiT-backend-group-8/Domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) RegisterUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.UserUsecase.RegisterUser(&user)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful. Check your email for verification link."})
}

func (h *Controller) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	err := h.UserUsecase.VerifyEmail(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Email verified successfully"})
}

func (h *Controller) Login(c *gin.Context) {
	// Corrected struct with exported fields
	type EmailPass struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var ep EmailPass

	if err := c.ShouldBindJSON(&ep); err != nil {
		c.JSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, refresher, err := h.UserUsecase.Login(ep.Email, ep.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresher": refresher})
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var cred Domain.Credential

	bind_err := c.BindJSON(&cred)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, err := h.UserUsecase.RefreshToken(cred.Email, cred.Refresher)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid refresh token "})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func (h *UserHandler) ForgotPassword(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}

	err := h.UserUsecase.GenerateResetPasswordToken(email)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Password reset email sent"})
}

func (h *UserHandler) StoreToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	err := h.UserUsecase.StoreToken(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Token stored successfully. You can now reset your password."})
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	var payload struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.UserUsecase.ResetPassword(payload.Token, payload.NewPassword)
	if err != nil {
		if err.Error() == "invalid or expired token" || err.Error() == "invalid token payload" || err.Error() == "invalid or mismatched token" {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successful"})
}
