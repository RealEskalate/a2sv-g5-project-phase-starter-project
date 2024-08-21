package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

// JwtAuthMiddleware checks the Authorization header, validates the token, and attaches user claims to the context.
func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			ctx.Abort()
			return
		}

		claims, err := IsAuthorized(parts[1], secret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Attach claims to context
		ctx.Set("user", claims)
		ctx.Next()
	}
}

// AdminMiddleware ensures that the user has admin privileges by checking the role claim.
func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			ctx.Abort()
			return
		}

		claims, ok := userClaims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user claims"})
			ctx.Abort()
			return
		}

		if role, ok := claims["role"].(string); !ok || role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
