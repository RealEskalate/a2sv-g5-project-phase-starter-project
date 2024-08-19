package infrastructure

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var jwtKey = []byte("BlogManagerSecretKey")

// Claims struct to include role
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// AuthMiddleware validates the JWT token and extracts claims
func AuthMiddleware(tokenCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Retrieve the token from the database
		var token Domain.Token
		err := tokenCollection.FindOne(c, bson.M{"access_token": tokenString}).Decode(&token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		// Check if the token is expired
		if token.IsExpired() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		fmt.Println("not expayered    ", token.ExpiresAt, time.Now(), "      kdnfmkng===========")

		// Parse the token claims (assuming you have a ParseToken function)
		claims, err := ParseToken(tokenString, []byte("BlogManagerSecretKey"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the username and role in the context
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// RoleMiddleware checks if the user has the required role.
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
