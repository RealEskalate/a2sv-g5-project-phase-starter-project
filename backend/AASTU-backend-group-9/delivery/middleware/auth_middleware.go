package middleware

import (
    "net/http"
	"blog/internal/tokenutil"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *tokenutil.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
            return
        }

        _, err := authService.ValidateAccessToken(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired access token"})
            return
        }

        c.Next()
    }
}
