package controllers

// import (
// 	"blog_g2/domain"
// 	"blog_g2/infrastructure"
// 	"context"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/oauth2"
// )

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OAuthController struct {
	Userusecase domain.UserUsecase
}

func NewOAuthController(Usermgr domain.UserUsecase) *OAuthController {
	return &OAuthController{
		Userusecase: Usermgr,
	}
}

func (oc *OAuthController) HandleGoogleLogin(c *gin.Context) {
	if infrastructure.GoogleOAuthConfig == nil {
		c.JSON(http.StatusInternalServerError, "Google OAuth configuration is not initialized")
		return
	}
	url := infrastructure.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (oc *OAuthController) HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := infrastructure.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	userInfo, err := infrastructure.GetGoogleUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	// Print userInfo for debugging
	c.JSON(http.StatusOK, gin.H{"user_info": userInfo})

	var user domain.User

	if email, ok := userInfo["email"].(string); ok {
		user.Email = email
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found in Google user info"})
		return
	}

	if name, ok := userInfo["name"].(string); ok {
		user.UserName = name
	} else {
		user.UserName = user.Email // Fallback to email if no name is provided
	}

	if picture, ok := userInfo["picture"].(string); ok {
		user.Imageuri = picture
	}

	user.Oauth = true

	_ = oc.Userusecase.RegisterUser(c, &user)

	logintoken, err := oc.Userusecase.LoginUser(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in", "token": logintoken})
}
