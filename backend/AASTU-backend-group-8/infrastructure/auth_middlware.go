package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdminMiddleware checks if the user is an admin
func AdminMiddleware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// secret_key := jwtService.GetKey("access")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.Role != "admin" && claims.Role != "root" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// secret_key := jwtService.GetKey("access")
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// You can use the claims, e.g., setting them in the context for later use
		c.Set("userID", claims.ID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
