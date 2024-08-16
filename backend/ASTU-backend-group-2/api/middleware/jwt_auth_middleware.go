package middleware

import (
	"net/http"
	"strings"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUserFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userRole := "user.role" + user.(string)

		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

func getUserRoleFromContext(c *gin.Context) (string, error) {
	userRole, exists := c.Get("user")
	if !exists {
		return "", nil
	}
	return userRole.(string), nil
}

func getUserFromHeader(c *gin.Context) (string, error) {
	// get the user from the header authorization field
	jwt := GetTokenFromHeader(c)
	if jwt == "" {
		return "", nil
	}
	// get the user from the jwt token
	user, err := tokenutil.GetUserFromToken(jwt)
	if err != nil {
		return "", err
	}
	return user, nil
}

func GetTokenFromHeader(c *gin.Context) string {
	token := ""
	if strings.HasPrefix(c.GetHeader("Authrorization"), "Bearer ") {
		token = strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	}
	return token
}
