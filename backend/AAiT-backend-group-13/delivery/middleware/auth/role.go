package authmiddleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
)

// Constants for context keys used in Gin middleware.
const (
	// ContextUserClaims is the key used to store user claims in the Gin context.
	ContextUserClaims = "userClaims"
)

// Authoriz returns a Gin middleware handler that performs authentication and
// optional authorization based on the provided JWT service and admin status requirement.
//
// It extracts the JWT from the "accessToken" cookie, decodes it, and checks if the user
// has the required admin status. If the user is authenticated and meets the authorization
// criteria, their claims are attached to the request context; otherwise, an appropriate
// HTTP status code is returned and the request is aborted.
func Authoriz(jwtService ijwt.Service, hasToBeAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the access token from the cookie.
		cookie, err := c.Cookie("accessToken")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				c.Status(http.StatusUnauthorized) // No cookie found.
			} else {
				c.Status(http.StatusInternalServerError) // Internal server error.
			}
			c.Abort()
			return
		}

		// Decode the token using the JWT service.
		claims, err := jwtService.Decode(cookie)
		if err != nil {
			c.Status(http.StatusUnauthorized) // Invalid token.
			c.Abort()
			return
		}

		// Check if the user meets the required admin status.
		isAdmin, ok := claims["is_admin"].(bool)
		if !ok || (!isAdmin && hasToBeAdmin) {
			c.Status(http.StatusForbidden) // Forbidden if admin status does not match.
			c.Abort()
			return
		}

		// Attach user claims to the request context for further use.
		c.Set(ContextUserClaims, claims)
		c.Next()
	}
}
