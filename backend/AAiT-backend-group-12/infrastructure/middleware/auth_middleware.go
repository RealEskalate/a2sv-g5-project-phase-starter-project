package middleware

import (
	"blog_api/domain"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MiddlewareError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, domain.Response{"message": message})
	c.Abort()
}

/*
This is the authorization middleware used for the endpoints. It accepts a set of
roles for which the endpoint is open.

WORKFLOW:
  - Obtains the JWT from the authorization header
  - Parses the JWT and verifies the signature
  - Checks the role of the user associated with the token
  - Calls `c.Next()` if the querying user has permission to access the endpoint
*/
func AuthMiddlewareWithRoles(validRoles []string, secret string, ValidateToken func(string, string) (*jwt.Token, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// obtain token from the request header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			MiddlewareError(c, 401, "Authorization header not found")
			return
		}

		headerSegments := strings.Split(authHeader, " ")
		if len(headerSegments) != 2 || strings.ToLower(headerSegments[0]) != "bearer" {
			MiddlewareError(c, 401, "Authorization header is invalid")
			return
		}

		// parses token with the correct signing method and checks for errors and token validity
		token, validErr := ValidateToken(headerSegments[1], secret)
		if validErr != nil {
			MiddlewareError(c, 401, validErr.Error())
			return
		}

		// check the expiry date of the token
		expiresAt, ok := token.Claims.(jwt.MapClaims)["expiresAt"]
		if !ok {
			MiddlewareError(c, 401, "Invalid token: Expiry date not found")
			return
		}

		expiresAtTime, convErr := time.Parse(time.RFC3339Nano, fmt.Sprintf("%v", expiresAt))
		if convErr != nil {
			MiddlewareError(c, 401, "Error while parsing expiry date: "+convErr.Error())
			return
		}

		if expiresAtTime.Compare(time.Now()) == -1 {
			MiddlewareError(c, 401, "Token expired")
			return
		}

		// get the role from the claims of the JWT
		userRole, ok := token.Claims.(jwt.MapClaims)["role"]
		if !ok {
			MiddlewareError(c, 401, "Invalid token: Role not found")
			return
		}

		// get the username from the claims of the JWT
		username, ok := token.Claims.(jwt.MapClaims)["username"]
		if !ok {
			MiddlewareError(c, 401, "Invalid token: Username not found")
			return
		}

		valid := false
		for _, validRole := range validRoles {
			if userRole == validRole {
				valid = true
				break
			}
		}

		if !valid {
			MiddlewareError(c, 403, fmt.Sprintf("'%v' roles are not allowed to access this endpoint", userRole))
			return
		}

		// pass the username to the handler
		c.Set("username", username)
		c.Next()
	}
}
