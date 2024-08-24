package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
)

type OAuthController struct {
	CompleteUserAuth func(res http.ResponseWriter, req *http.Request) (goth.User, error)
	BeginAuthHandler func(res http.ResponseWriter, req *http.Request)
}

// NewOAuthController initializes the OAuth controller
func NewOAuthController(
	CompleteUserAuth func(res http.ResponseWriter, req *http.Request) (goth.User, error),
	BeginAuthHandler func(res http.ResponseWriter, req *http.Request),
) *OAuthController {
	return &OAuthController{
		CompleteUserAuth: CompleteUserAuth,
		BeginAuthHandler: BeginAuthHandler,
	}
}

// Handler for /auth/google/start
func (controller *OAuthController) GoogleAuthInit(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Add("provider", "google")
	c.Request.URL.RawQuery = query.Encode()
	controller.BeginAuthHandler(c.Writer, c.Request)
}

// Handler for /auth/google/callback
func (controller *OAuthController) OAuthCallback(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Add("provider", "google")
	c.Request.URL.RawQuery = query.Encode()
	user, err := controller.CompleteUserAuth(c.Writer, c.Request)
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
