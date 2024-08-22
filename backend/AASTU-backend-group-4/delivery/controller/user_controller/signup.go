package user_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) SignUp(c *gin.Context) {
	var req domain.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := uc.userUsecase.SignUp(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"refresh_token",        // Cookie name
		resp.RefreshToken,      // Cookie value (refresh token)
		uc.Env.RefreshTokenExpiryHour,            // Cookie max age (7 days)
		"/",                    // Cookie path (available across entire site)
		"localhost",            // Domain for cookie (use "localhost" for local development)
		false,                  // Secure flag (false for HTTP, true for HTTPS)
		true,                    // HttpOnly flag (true to prevent JavaScript access)
	)

	c.JSON(http.StatusCreated, resp)
}
