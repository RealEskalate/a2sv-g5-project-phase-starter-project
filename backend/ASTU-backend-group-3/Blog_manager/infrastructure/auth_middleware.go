package infrastructure

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func RoleMiddleware(requiredRole string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		role, exists := c.Get("role")
// 		if !exists || role != requiredRole {
// 			c.AbortWithStatus(http.StatusForbidden)
// 			return
// 		}
// 		c.Next()
// 	}
// }

// // AuthMiddleware validates the JWT token and extracts claims
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
// 			c.Abort()
// 			return
// 		}

// 		if strings.HasPrefix(tokenString, "Bearer ") {
// 			tokenString = tokenString[7:]
// 		}

// 		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, jwt.ErrSignatureInvalid
// 			}
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
// 			c.Abort()
// 			return
// 		}

// 		claims, ok := token.Claims.(*Claims)
// 		if !ok || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("username", claims.Username)
// 		c.Set("role", claims.Role)
// 		c.Next()
// 	}
// }
import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// AuthMiddleware is the authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Fatal("SECRET_KEY is not set in .env file")
		}

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		// Check if the token is valid
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token is expired
		if !token.Valid && claims.ExpiresAt < time.Now().Unix() {
			refreshToken := c.GetHeader("Refresh-Token")

			// Validate the refresh token
			if refreshToken == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not provided"})
				c.Abort()
				return
			}

			// Parse the refresh token
			refreshClaims := &Claims{}
			refreshTokenObj, err := jwt.ParseWithClaims(refreshToken, refreshClaims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil || !refreshTokenObj.Valid || refreshClaims.Username != claims.Username {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
				c.Abort()
				return
			}

			// Check if refresh token is in the blacklist (pseudo code, implement your own blacklist check)
			if IsTokenBlacklisted(refreshToken) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is blacklisted"})
				c.Abort()
				return
			}

			// Generate a new access token
			newTokenString, err := GenerateToken(claims.Username, claims.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new access token"})
				c.Abort()
				return
			}

			// Send the new access token in the response header
			c.Header("New-Access-Token", newTokenString)
		}

		// Proceed to the next middleware or request handler
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Load the secret key from the environment
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Secret key not set in environment"})
			c.Abort()
			return
		}

		// Initialize a new Claims object
		claims := &Claims{}

		// Parse the JWT token and validate it
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		// Check if there was an error in parsing or the token is invalid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the role in the claims is "admin"
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: admin role required"})
			c.Abort()
			return
		}

		// If the user is an admin, proceed to the next handler
		c.Next()
	}
}
