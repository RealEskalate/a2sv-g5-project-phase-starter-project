package middleware

import (
	"net/http"
	"strings"

	"github.com/RealEskalate/blogpost/config"
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	tokenservice "github.com/RealEskalate/blogpost/infrastructure/token_service"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/gin-gonic/gin"
)

func LoggedIn(TS tokenservice.TokenService_imp) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for the Authorization header
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization header missing"})
			c.Abort()
			return
		}

		// Extract the Bearer token
		authSplit := strings.Split(auth, " ")
		if len(authSplit) != 2 || strings.ToLower(authSplit[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization format is incorrect"})
			c.Abort()
			return
		}

		accessToken := authSplit[1]
		user, err := TS.ValidateAccessToken(accessToken)
		if err != nil {
			// Check for token expiration or invalid token
			if err.Error() == "token has expired" {
				// Try to get the refresh token from cookies
				refreshToken, err := c.Cookie("refresh_token")
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"msg": "Refresh token missing or expired"})
					c.Abort()
					return
				}

				// Validate the refresh token
				user, err = TS.ValidateRefreshToken(refreshToken)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid refresh token"})
					c.Abort()
					return
				}

				// Generate a new access token
				newAccessToken, err := TS.GenerateAccessToken(*user)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to generate new access token"})
					c.Abort()
					return
				}

				// Send the new access token in response headers (optional)
				c.Header("Authorization", "Bearer "+newAccessToken)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid access token"})
				c.Abort()
				return
			}
		}

		// Token is valid, store the user in the context
		SC := config.ServerConnection{}
		SC.Connect_could()
		coll := &database.MongoCollection{Collection: SC.Client.Database("BlogPost").Collection("Users")}
		UR := repository.NewUserRepository(coll)

		userr, err := UR.GetUserDocumentByID(user.ID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal server error"})
			c.Abort()
		}
		usr := domain.CreateResponseUser(userr)
		c.Set("user", usr)

		// Proceed to the next handler
		c.Next()
	}
}
