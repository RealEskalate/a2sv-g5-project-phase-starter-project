package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdminMiddleware is a middleware function that checks if the user has the "ADMIN" role.
// If the user does not have the "ADMIN" role, it returns a JSON response with an error message and aborts the request.
// Otherwise, it allows the request to proceed to the next middleware or handler.
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
