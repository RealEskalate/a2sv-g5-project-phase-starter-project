package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		authToken := t[1]

		claims, err := VerifyToken(authToken, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email claims"})
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username claims"})
			c.Abort()
			return
		}
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin role claims"})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Set("email", email)
		c.Set("username", username)
		c.Set("isAdmin", isAdmin)
		c.Next()
	}
}
