package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user role from the context
		user_role, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "role not found in context"})
			c.Abort()
			return
		}

		if user_role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
			c.Abort()
			return
		}
	}
}
