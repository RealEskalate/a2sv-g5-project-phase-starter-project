package infrastructure

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configJwt,err := config.LoadConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}
	var jwtSecret = []byte(configJwt.Jwt.JwtKey)
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, err)
			return
		}
		if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
			// Check expiration
			if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				c.Abort()
				return
			}

			// Extract the claims
			if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {

				
				c.Set("id", claims.ID)
				c.Set("username", claims.Username)
				c.Set("isAdmin", claims.IsAdmin)
				c.Set("email", claims.Email)
				c.Next()
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
	}
}
