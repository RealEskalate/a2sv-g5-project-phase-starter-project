package middleware

import (
	"net/http"
	"strings"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.ErrorResponse{Message: "Invalid Token"})
			return
		}
		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, secret)

		if err != nil || !authorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.ErrorResponse{Message: err.Error()})
			return
		}

		claims, err := tokenutil.ExtractUserClaimsFromToken(authToken, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entities.ErrorResponse{Message: err.Error()})
			return
		}
		c.Set("x-user-id", claims["id"])
		c.Set("x-user-role", claims["role"])
		c.Set("x-user-owner", claims["is_owner"])
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.MustGet("x-user-role")
		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, entities.ErrorResponse{Message: "unauthorized"})
			ctx.Abort()
		}
		ctx.Next()
	}
}
