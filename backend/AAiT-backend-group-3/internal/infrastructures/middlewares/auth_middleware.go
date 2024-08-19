package middlewares

import (
	"AAIT-backend-group-3/internal/infrastructures/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(jwtservice services.IJWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authSlice := strings.Split(authHeader, " ")
		if len(authSlice) != 2 || strings.ToLower(authSlice[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token, err := jwtservice.ValidateAccessToken(authSlice[1])
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

