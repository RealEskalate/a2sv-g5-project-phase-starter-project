package controllers

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type RefreshController struct {
	RefreshUsecase interfaces.RefreshUsecase
	JwtService     interfaces.JwtService
}

func NewRefreshController(
	refreshUsecase interfaces.RefreshUsecase,
	jwtService interfaces.JwtService,
	oauthService interfaces.OAuthService,
) *RefreshController {
	return &RefreshController{
		RefreshUsecase: refreshUsecase,
		JwtService:     jwtService,
	}
}

func (refreshController *RefreshController) Refresh(ctx *gin.Context) {
	userId := ctx.GetString("id")

	authHeader := ctx.GetHeader("Authorization")
	authParts, err := refreshController.JwtService.ValidateAuthHeader(authHeader)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	refresh_token := authParts[1]
	accessToken, e := refreshController.RefreshUsecase.RefreshToken(ctx.Request.Context(), userId, refresh_token)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"access token": accessToken})
}
