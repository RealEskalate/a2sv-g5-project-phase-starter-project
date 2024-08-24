package user_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) ResetPassword(c *gin.Context) {
	var req domain.ResetPasswordRequest

	// Bind JSON request to the ResetPasswordRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the ResetPassword use case
	if err := uc.userUsecase.ResetPassword(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, domain.ResetPasswordResponse{Message: "Password has been successfully reset"})
}
