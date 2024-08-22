package infrastructure

import (
	"astu-backend-g1/config"
	"astu-backend-g1/domain"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthController struct {
}

func NewAuthController() GeneralAuthorizationController {
	return &AuthController{}
}

func (ac *AuthController) AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is the authorization middleware")
		configJwt, err := config.LoadConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		var jwtSecret = []byte(configJwt.Jwt.JwtKey)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		fmt.Println(authHeader)
		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		fmt.Println(tokenString)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
			fmt.Println("this is a valid claim", claims)
			c.Set("claims", claims)
			c.Next()
		}
		fmt.Println("invalid token")

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
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, errors.New("invalid token"))
		c.Abort()
	}
}

func (ac *AuthController) USERMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("user middleware")
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
		fmt.Println("this is token: ", token)
		if err != nil {
			c.JSON(http.StatusForbidden, err)
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
			fmt.Println("this is claims: ", claims)
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
