package infrastructure

import (
	"net/http"
	"strings"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type middlewareService struct {
	jwtService domain.JwtService
}

func NewMiddlewareService(jwtService domain.JwtService) domain.MiddlewareService {
	return &middlewareService{jwtService: jwtService}
}

func (m *middlewareService) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			ctx.Abort()
			return
		}

		tokenString := authParts[1]
		token, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			ctx.Abort()
			return
		}
		
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("username", claims["username"])
		ctx.Set("role", claims["role"])
		ctx.Set("exp", claims["exp"])
		ctx.Set("iss", claims["iss"])

		ctx.Next()
	}
}

func (m *middlewareService) Authorize(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, _ := m.jwtService.ValidateToken(tokenString)
		claims, _ := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)

		if userRole != role {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized for this action"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}