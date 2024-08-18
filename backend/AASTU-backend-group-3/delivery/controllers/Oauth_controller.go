package controllers

import (
	"context"
	"encoding/json"
	"group3-blogApi/domain"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"group3-blogApi/infrastracture"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     "33361312477-ddpanahl6fj6sk82mav0c2kijcpcgvts.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
)


func (uc *UserController) HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (uc *UserController) HandleGoogleCallback(c *gin.Context) {
    if c.Query("state") != oauthStateString {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth state"})
        return
    }

    token, err := googleOauthConfig.Exchange(context.Background(), c.Query("code"))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code"})
        return
    }

    client := googleOauthConfig.Client(context.Background(), token)
    response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
        return
    }
    defer response.Body.Close()

    data, _ := ioutil.ReadAll(response.Body)

    var googleUser domain.OAuthUserInfo
    json.Unmarshal(data, &googleUser)

    googleUser.Provider = domain.Google

    ipAddress := c.ClientIP()
    userAgent := c.Request.UserAgent()
    deviceFingerprint := infrastracture.GenerateDeviceFingerprint(ipAddress, userAgent)

    loginResponse, err := uc.UserUsecase.OAuthLogin(googleUser, deviceFingerprint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login with Google"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tokens": loginResponse})
}