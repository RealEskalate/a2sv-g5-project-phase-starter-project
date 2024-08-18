package middleware

import (
	"blog/internal/tokenutil"
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMidd(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	claims, err := tokenutil.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("claim", *claims)
	c.Next()
}

