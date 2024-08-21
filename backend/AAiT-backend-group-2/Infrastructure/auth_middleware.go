package infrastructure

import (
	"AAiT-backend-group-2/Infrastructure/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)



func AuthMiddleWare(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		jwtService := services.NewJWTService([]byte(jwtSecret))

		jwtToken := authParts[1]
		claims, err := jwtService.ValidateToken(jwtToken, "access_token")
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", claims["userID"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

func RoleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("role")
        if !exists || userRole != "admin" {
            c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            c.Abort()
            return
        }

        c.Next()
    }
}