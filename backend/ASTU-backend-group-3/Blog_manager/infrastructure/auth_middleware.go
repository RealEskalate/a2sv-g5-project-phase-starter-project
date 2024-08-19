package infrastructure

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("BlogManagerSecretKey")

// Claims struct to include role
type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

// AuthMiddleware validates the JWT token and extracts claims
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        // Remove "Bearer " prefix if present
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

        claims, err := ParseToken(tokenString, []byte("BlogManagerSecretKey"))

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }



        if claims.ExpiresAt < time.Now().Unix() {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
            c.Abort()
            return
        }
       
    
 
        c.Set("username", claims.Username)
        c.Set("role", claims.Role)
        c.Next()
    }
}

// RoleMiddleware checks if the user has the required role.
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != requiredRole {
            c.AbortWithStatus(http.StatusForbidden)
            return
        }
        c.Next()
    }
}