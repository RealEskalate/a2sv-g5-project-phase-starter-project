package infrastructures

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errss := godotenv.Load()
		jwtSecret := os.Getenv("jwtSecret")
		fmt.Println("jwtSecret", jwtSecret)
		if errss != nil {
			ctx.JSON(500, gin.H{
				"message": "server error"})
			ctx.Abort()
			return
		}
		// check if user is authenticated
		// if not, return 401
		// else, continue
		tokens := ctx.GetHeader("Authorization")
		if tokens == "" {
			ctx.JSON(400, gin.H{"message": "Invalid Token!"})
			ctx.Abort()
			return
		}
		authtokens := strings.Split(tokens, " ")
		fmt.Println(authtokens)
		if len(authtokens) != 2 || strings.ToLower(authtokens[0]) != "bearer" {
			ctx.JSON(401, gin.H{"error": "Invalid authorization header"})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(authtokens[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

			}
			return []byte(jwtSecret), nil
		})
		fmt.Println(err)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "Invalid JWT", "valid": token.Valid})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("claims", claims)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}

// Isadmin middleware
func Isadmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "No token claims found"})
			c.Abort()
			return
		}

		jwtClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if role, ok := jwtClaims["role"].(string); !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have the required role"})
			c.Abort()
			return
		}

		c.Next()
	}
}
