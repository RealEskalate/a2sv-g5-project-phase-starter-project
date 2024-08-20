package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type OAuthController struct{}

// NewOAuthController initializes the OAuth controller
func NewOAuthController() *OAuthController {
	return &OAuthController{}
}

// Handler for /auth/google/start
func (controller *OAuthController) GoogleAuthInit(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Add("provider", "google")
	c.Request.URL.RawQuery = query.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// Handler for /auth/google/callback
func (controller *OAuthController) OAuthCallback(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Add("provider", "google")
	c.Request.URL.RawQuery = query.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jsonString := string(res)
	c.JSON(http.StatusAccepted, jsonString)
}
