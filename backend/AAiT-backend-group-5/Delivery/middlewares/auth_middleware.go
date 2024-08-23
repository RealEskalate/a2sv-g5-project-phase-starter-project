package middlewares

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JwtService   interfaces.JwtService
	oauthService interfaces.OAuthService
	repo         interfaces.SessionRepository
}

func NewJwtAuthMiddleware(jwtService interfaces.JwtService,
	repo interfaces.SessionRepository,
	oauthService interfaces.OAuthService,
) AuthMiddleware {
	return AuthMiddleware{
		JwtService:   jwtService,
		repo:         repo,
		oauthService: oauthService,
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
		authorizedTokenJWT, err := j.JwtService.ValidateToken(tokenString)
		authorizedOauthToken, oErr := j.oauthService.OAuthTokenValidator(tokenString, c)

		if err != nil && oErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		authorizedToken := models.JWTCustome{}
		if authorizedTokenJWT != nil {
			authorizedToken = *authorizedTokenJWT
		} else {
			authorizedToken = *authorizedOauthToken
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
		authorizedTokenJWT, err := j.JwtService.ValidateToken(tokenString)
		authorizedOauthToken, oErr := j.oauthService.RefreshTokenValidator(tokenString, c)

		if err != nil && oErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		authorizedToken := models.JWTCustome{}
		if authorizedTokenJWT != nil {
			authorizedToken = *authorizedTokenJWT
		} else {
			authorizedToken = *authorizedOauthToken
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
