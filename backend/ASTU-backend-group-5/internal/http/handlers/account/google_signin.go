package account

import (
	"blogApp/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GoogleCallback is a handler that handles the google oauth callback
func (h *UserHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	user, token, err := h.UserUsecase.GoogleCallback(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getUSer := &domain.GetUserDTO{
		Email:    user.Email,
		UserName: user.UserName,
		Role:     user.Role,
		Profile:  user.Profile,
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  getUSer,
		"token": token,
	})
}
