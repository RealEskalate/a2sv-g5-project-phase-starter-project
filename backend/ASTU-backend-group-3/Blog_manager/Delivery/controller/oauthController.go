package controller

import (
	"net/http"

	"ASTU-backend-group-3/Blog_manager/infrastructure"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (u *UserController) LoginHandler(c *gin.Context) {
	// Generate a URL for the Google OAuth 2.0 consent screen
	url := infrastructure.OAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (u *UserController) CallbackHandler(c *gin.Context) {
	// Retrieve the authorization code from the callback URL
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code parameter"})
		return
	}

	// Use the OAuthLogin use case to handle the OAuth login process
	_, token, err := u.UserUsecase.OAuthLogin(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the user information and token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
