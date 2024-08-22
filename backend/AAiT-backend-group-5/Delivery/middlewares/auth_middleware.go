package middlewares

import (
	"net/http"
	// "log"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JwtService interfaces.JwtService
	repo       interfaces.SessionRepository
}

func NewJwtAuthMiddleware(jwtService interfaces.JwtService,
	repo interfaces.SessionRepository,
) AuthMiddleware {
	return AuthMiddleware{
		JwtService: jwtService,
		repo:       repo,
	}
}

func (j *AuthMiddleware) JWTAuthMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		auth_parts, err := j.JwtService.ValidateAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := auth_parts[1]
		authorizedToken, err := j.JwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		session, nErr := j.repo.GetToken(c, authorizedToken.ID)

		if nErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorizedddd"})
			c.Abort()
			return
		}

		if session.AccessToken != tokenString {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("id", authorizedToken.ID)
		c.Set("email", authorizedToken.Email)
		c.Set("role", authorizedToken.Role)

		c.Next()
	}
}

func (j *AuthMiddleware) JWTRefreshAuthMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		auth_parts, err := j.JwtService.ValidateAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := auth_parts[1]
		authorizedToken, err := j.JwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		_, nErr := j.repo.GetToken(c, authorizedToken.ID)

		if nErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("id", authorizedToken.ID)
		c.Set("email", authorizedToken.Email)
		c.Set("role", authorizedToken.Role)

		c.Next()
	}
}
