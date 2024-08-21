package infrastructure

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	userCollection *mongo.Collection
}

func NewAuthController(userCollection *mongo.Collection) GeneralAuthorizationController {
	return &AuthController{
		userCollection: userCollection,
	}
}

func (ac *AuthController) AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configJwt, err := config.LoadConfig()
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
		token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, err)
			c.Abort()

		}
		claims, ok := token.Claims.(*domain.Claims)
		if claims.ExpiresAt < time.Now().Unix() {
			var TheUser domain.User
			err := ac.userCollection.FindOne(context.TODO(), bson.M{"_id": claims.ID}).Decode(&TheUser)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
				c.Abort()
				return

			}
			refreshToken, err := jwt.ParseWithClaims(TheUser.RefreshToken, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err == nil {
				refreshClaims, ok := refreshToken.Claims.(*domain.Claims)
				if refreshClaims.ExpiresAt < time.Now().Unix() {
					ac.userCollection.UpdateOne(context.TODO(), bson.M{"_id": claims.ID}, domain.User{RefreshToken: ""})
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
					c.Abort()
					return

				}
				if ok && refreshToken.Valid {
					claims.ExpiresAt = time.Now().Add(1 * time.Minute).Unix()
					newTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
					signedToken, err := newTokenString.SignedString(jwtSecret)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign new token"})
						c.Abort()
						return
					}
					c.IndentedJSON(200, gin.H{"refreshed access token": signedToken})
					return
				}
			}
		}
		if ok && token.Valid {
			if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
				c.Set("claims", claims)
				c.Set("id", claims.ID)
				c.Set("username", claims.Username)
				c.Set("isAdmin", claims.IsAdmin)
				c.Set("email", claims.Email)
				c.Next()
				return
			}
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
		c.Abort()
	}
}

func (ac *AuthController) ADMINMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configJwt, err := config.LoadConfig()
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

		token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, err)
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
			if claims.IsAdmin {
				c.Next()
			}
			c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
			c.Abort()
			return

		}
		c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
		c.Abort()
	}
}

func (ac *AuthController) USERMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configJwt, err := config.LoadConfig()
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

		token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			c.JSON(http.StatusForbidden, err)
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
			if claims.IsActive {
				c.Next()
				return
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
				c.Abort()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
		c.Abort()
	}
}

func GetClaims(c *gin.Context) (domain.Claims, error) {
	configJwt, err := config.LoadConfig()
	if err != nil {
		return domain.Claims{}, err
	}
	var jwtSecret = []byte(configJwt.Jwt.JwtKey)
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return domain.Claims{}, errors.New("missing authorization header")
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return domain.Claims{}, err
	}
	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		return *claims, err
	}
	return domain.Claims{}, errors.New("invalid token")
}
