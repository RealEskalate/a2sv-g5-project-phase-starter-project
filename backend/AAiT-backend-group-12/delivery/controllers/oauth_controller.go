package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func BeginGoogleAuth(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Add("provider", "google")
	c.Request.URL.RawQuery = query.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func OAuthCallback(c *gin.Context) {
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
