package infrastructure

import (
	"net/http"
	"strings"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type middlewareService struct {
	jwtService   domain.JwtService
	redisService domain.CacheService
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewMiddlewareService(jwtService domain.JwtService, cacheService domain.CacheService) domain.MiddlewareService {
	return &middlewareService{jwtService: jwtService, redisService: cacheService}
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

		blacklisted, errRedis := m.redisService.Get(authParts[1]) // check if token is blacklisted
		if errRedis != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errRedis.Error()})
			ctx.Abort()
			return
		}
		if blacklisted != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is blacklisted"})
			ctx.Abort()
			return
		}

		token, err := m.jwtService.ValidateAccessToken(authParts[1])
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

func (m *middlewareService) Authorize(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := m.jwtService.ValidateAccessToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["role"] == nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized for this action"})
			ctx.Abort()
			return
		}

		userRole := claims["role"].(string)
		validRole := false
		for _, r := range role {
			if r == userRole {
				validRole = true
				break
			}
		}

		if !validRole {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized for this action"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
