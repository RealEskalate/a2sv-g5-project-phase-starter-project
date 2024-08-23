package middleware

import (
	"net/http"
	"strings"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				claims, err := tokenutil.ExtractUserClaimsFromToken(authToken, secret)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					return
				}
				c.Set("x-user-id", claims["id"])
				c.Set("x-user-role", claims["role"])
				c.Set("x-user-owner", claims["is_owner"])
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
	}
}
