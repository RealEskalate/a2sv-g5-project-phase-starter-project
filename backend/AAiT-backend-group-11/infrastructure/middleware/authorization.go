package middleware

import (
	"backend-starter-project/domain/interfaces"
	"errors"
	"log"
	"net/http"
	"strings"

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
		refresh, err := c.Cookie("refresh_token")
		if err != nil {
			err := middleware.TokenService.VerifyRefreshToken(refresh)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}
		}

		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		authParts := strings.Split(header, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusBadRequest, "Authorization header format must be Bearer {token}")
		}

		err = middleware.TokenService.VerifyAccessToken(authParts[1])

		if err != nil {
			if errors.Is(err, &interfaces.ErrTokenExpired{}) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
				c.Abort()
			}

			if errors.Is(err, &interfaces.ErrTokenInvalid{}) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		claims := middleware.TokenService.GetClaimsFromToken(authParts[1])
		if role != "" {
			if claims["role"] != role {
				c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				c.Abort()
				return
			}
		}

		log.Println(claims)

		c.Set("userId", claims["userId"])
		c.Set("role", claims["role"])
		c.Set("userName", claims["useName"])
		c.Next()
	}

}
