package infrastructure

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware is a middleware function that performs JWT authentication.
// It checks the Authorization header for a valid JWT token and sets the claims to the context.
// If the token is invalid or missing, it returns an error response.
// The secret parameter is used to validate the token's signature.

func JWTAuthMiddelware(service interfaces.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		auth_parts, err := service.ValidateAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// check if token is authorized
		tokenString := auth_parts[1]
		authorizedToken, err := service.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// set the claims to the context
		c.Set("id", authorizedToken.ID)
		c.Set("email", authorizedToken.Email)
		c.Set("role", authorizedToken.Role)

		c.Next()
	}
}
