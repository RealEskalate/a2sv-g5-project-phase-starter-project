package Infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserMiddleware(jwt_signer string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwt_signer), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		fmt.Println(claims)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")

		if !exists {
			// fmt.Printf("Claims: %v\n", claims)
			c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
			c.Abort()
			return
		}

		userClaims := claims.(jwt.MapClaims)

		role := strings.ToLower(userClaims["role"].(string))

		if role != "admin" && role != "super-admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func SuperAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")

		if !exists {
			// fmt.Printf("Claims: %v\n", claims)
			c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
			c.Abort()
			return
		}

		userClaims := claims.(jwt.MapClaims)

		role := strings.ToLower(userClaims["role"].(string))

		if role != "super-admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
