package controllers

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *infrastructure.Config
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	valid, _ := infrastructure.IsAuthorized(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}
	actvuser, err := rtc.RefreshTokenUsecase.CheckActiveUser(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found in active user"})
		return
	}
	if id == actvuser.ID.Hex() && !valid {
		user_agent := c.Request.UserAgent()
		rtc.RefreshTokenUsecase.RemoveActiveUser(c, id, user_agent)
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "session expired"})
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}
	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
