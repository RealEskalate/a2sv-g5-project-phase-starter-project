package user_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) Login(c *gin.Context) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := uc.userUsecase.Login(c.Request.Context(), req)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}
