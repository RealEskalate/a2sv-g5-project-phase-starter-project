package user_controller

import (
	"blog-api/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) ForgotPassword(c *gin.Context) {
	var req domain.ForgotPasswordRequest

	// Parse the JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ForgotPasswordResponse{Message: "Invalid request payload"})
		return
	}

	// Call the usecase to request a password reset
	err := uc.userUsecase.RequestPasswordReset(context.Background(), req.Email, uc.Env.FrontendBaseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ForgotPasswordResponse{Message: "Failed to process password reset"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, domain.ForgotPasswordResponse{Message: "Password reset link sent to your email"})
}
