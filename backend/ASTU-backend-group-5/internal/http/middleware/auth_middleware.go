package middleware

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase"
	"blogApp/pkg/jwt"
	"blogApp/pkg/mongo"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader("Authorization"), " ")[1]
		if authHeader == "" {
			respondUnauthorized(c, "Authorization header required")
			return
		}

		tokenString := authHeader
		fmt.Println(authHeader)
		if tokenString == "" {
			respondUnauthorized(c, "Bearer token required")
			return
		}

		// Validate the token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			respondUnauthorized(c, "Invalid token")
			return
		}
		tokenCollection := mongo.GetCollection("tokens")
		mongoTokenRepo := mongodb.NewMongoTokenRepository(tokenCollection)
		tokenUsecase := usecase.NewTokenUsecase(mongoTokenRepo)

		// Check if the token is blacklisted
		isBlacklisted, err := tokenUsecase.IsTokenBlacklisted(context.Background(), tokenString, domain.AccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token blacklist"})
			c.Abort()
			return
		}
		if isBlacklisted {
			respondUnauthorized(c, "Token is blacklisted")
			return
		}

		// Set the claims in the context and proceed
		c.Set("claims", claims)
		c.Next()
	}
}

// func ReafreshTokenMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			respondUnauthorized(c, "Authorization header required")
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		if tokenString == "" {
// 			respondUnauthorized(c, "Bearer token required")
// 			return
// 		}

// 		// Validate the token
// 		claims, err := jwt.ValidateToken(tokenString)
// 		if err != nil {
// 			respondUnauthorized(c, "Invalid token")
// 			return
// 		}

// 		// Set the claims in the context and proceed
// 		c.Set("claims", claims)
// 		c.Next()
// 	}
// }

func respondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}
