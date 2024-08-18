package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) LogoutController(c *gin.Context) {
	userID := c.GetString("userID")

	err := uc.usecase.DeleteRefreshTokenByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User logged out successfully"})
}
