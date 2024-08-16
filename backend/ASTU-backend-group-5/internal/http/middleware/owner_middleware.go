package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "blogApp/internal/domain"
)

func OwnerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, exists := c.Get("user")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
            c.Abort()
            return
        }

        userClaims, ok := claims.(*domain.Claims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
            c.Abort()
            return
        }

        if userClaims.Role != "owner" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Owner access required"})
            c.Abort()
            return
        }

        c.Next()
    }
}