package user_controller

import (
	"blog-api/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) ResetPasswordController(c *gin.Context) {
	var request user.ResetPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := uc.usecase.ResetPassword(c, request.ResetToken, request.NewPassword, uc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Password reset successfully"})
}
