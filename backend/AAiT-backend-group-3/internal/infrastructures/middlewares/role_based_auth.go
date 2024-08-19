package middlewares 


import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"strings"
)

func RoleAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(401, gin.H{
				"error": "Claims not found in context",
			})
			c.Abort()
			return
		}

		claimsMap, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT claims format"})
			c.Abort()
			return
		}

		role, ok := claimsMap["role"].(string)
		if !ok || role == "" {
			c.JSON(401, gin.H{"error": "Role not found in JWT claims"})
			c.Abort()
			return
		}

		roleAuthorized := false
		for _, elem := range roles {
			if strings.EqualFold(elem, role) {
				roleAuthorized = true
				break
			}
		}

		if !roleAuthorized {
			c.JSON(403, gin.H{"error": "Your role does not have access to this resource"})
			c.Abort()
			return
		}

		c.Set("userId", claimsMap["userId"])
		c.Next()
	}
}
