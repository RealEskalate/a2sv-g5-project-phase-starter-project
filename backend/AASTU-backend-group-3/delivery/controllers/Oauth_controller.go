package controllers

import (
	"net/http"

	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)



type OAuthController struct {
	oauthUsecase usecase.OAuthUsecase
}

func NewOAuthController(oauthUsecase usecase.OAuthUsecase) *OAuthController {
	return &OAuthController{oauthUsecase: oauthUsecase}
}

func (c *OAuthController) HandleGoogleLogin(ctx *gin.Context) {
	url := c.oauthUsecase.GetLoginURL()
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *OAuthController) HandleGoogleCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	code := ctx.Query("code")

	_, err := c.oauthUsecase.HandleCallback(state, code)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error: %v", err)
		return
	}

	ctx.String(http.StatusOK, "Response: %s", gin.H{"message": "Successfully logged in"})
}

