package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not allowed"})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok || strings.ToUpper(roleStr) != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not allowed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
