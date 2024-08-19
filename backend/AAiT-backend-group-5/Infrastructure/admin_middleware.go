package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthenticateAdmin is a middleware function that checks if the user is an admin.
// It retrieves the user role from the context and verifies if it is "ADMIN".
// If the user is not an admin, it returns an unauthorized error.
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
