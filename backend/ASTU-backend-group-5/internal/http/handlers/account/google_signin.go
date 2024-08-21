package account

import (
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
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
