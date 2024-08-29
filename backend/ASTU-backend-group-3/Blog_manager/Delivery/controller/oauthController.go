package controller

import (
	"net/http"

	"ASTU-backend-group-3/Blog_manager/infrastructure"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (u *UserController) LoginHandler(c *gin.Context) {
	// Generate the OAuth2 URL for Google login with the "prompt" parameter to force account selection
	url := infrastructure.OAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "select_account"))

	// Redirect the user to Google's OAuth2 login page
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
