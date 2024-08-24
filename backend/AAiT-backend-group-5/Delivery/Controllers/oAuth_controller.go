package controllers

import (
	"encoding/json"
	"net/http"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"golang.org/x/oauth2"
)

type oAuthController struct {
	config  oauth2.Config
	env     config.Env
	usecase interfaces.OAuthUseCase
}

func NewOAuthController(
	config oauth2.Config,
	env config.Env,
	usecase interfaces.OAuthUseCase,

) interfaces.OAuthController {
	return &oAuthController{
		config:  config,
		env:     env,
		usecase: usecase,
	}
}

func (c *oAuthController) LoginHandlerController(ctx *gin.Context) {
	user_ag := ctx.Request.UserAgent()

	agent := user_agent.New(user_ag)
	if agent.Mobile() {
		ctx.IndentedJSON(http.StatusTemporaryRedirect, gin.H{"error": "Google login is not supported on mobile devices, yet"})
		return
	}

	ctx.HTML(http.StatusOK, "login.html", nil)
}

func (c *oAuthController) OAuthHanderController(ctx *gin.Context) {
	url := c.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *oAuthController) OAuthCallbackHandler(ctx *gin.Context) {
	code := ctx.Query("code")

	t, err := a.config.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := a.config.Client(ctx, t)
	resp, err := client.Get(a.env.GOOGEL_CLIENT_ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var v map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRequest := dtos.OAuthRequest{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Email:        v["email"].(string),
		Name:         v["name"].(string),
	}

	if err := a.usecase.LoginHandlerUseCase(ctx, userRequest); err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  t.AccessToken,
		"refresh_token": t.RefreshToken,
	})
}
