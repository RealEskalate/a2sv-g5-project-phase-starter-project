package middleware

import (
	"backend-starter-project/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	AuthMiddleware(role string) gin.HandlerFunc
}

type authMiddleware struct {
	TokenService interfaces.TokenService
}

func NewAuthMiddleware(token_service interfaces.TokenService) AuthMiddleware {
	return &authMiddleware{TokenService: token_service}
}

func (middleware *authMiddleware) AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		refresh := c.GetHeader("refresh_token")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		accessToken, err := middleware.TokenService.VerifyAccessToken(refresh, header)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims := middleware.TokenService.GetClaimsFromToken(accessToken)
		if claims["role"] != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}
		c.Next()
	}

}
