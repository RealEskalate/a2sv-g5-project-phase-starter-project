package controllers

import (
	"github.com/gin-gonic/gin"
	"AAIT-backend-group-3/internal/domain/dtos"
	"AAIT-backend-group-3/internal/usecases"
)


type UserController struct {
	user_usecase *usecases.UserUsecase
}

func NewController(u *usecases.UserUsecase) *UserController{
	return &UserController{
		user_usecase: u,
	}
}

func (usecases *UserController) ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := usecases.user_usecase.SendPasswordResetLink(req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send reset link"})
		return
	}

	c.JSON(200, gin.H{"message": "Reset link sent"})
}

func (usecases *UserController) ResetPassword(c *gin.Context) {
	var req dtos.ResetPassword

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := usecases.user_usecase.ResetPassword(req.OTP, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to reset password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successful"})
}

func (usecases *UserController) Register(c *gin.Context) {

}

func (usecases *UserController) Login(c *gin.Context) {

}

func (usecase *UserController) RefreshToken(c *gin.Context) {
	
}