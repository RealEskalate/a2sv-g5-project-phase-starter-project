package controller

import (
	"net/http"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshTokenController struct {
	RefreshTokenUsecase entities.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request entities.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(err))
		return
	}

	if valid, err := tokenutil.IsAuthorized(string(request.RefreshToken), rtc.Env.RefreshTokenSecret); !valid || err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrInvalidToken))
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUserNotFound))
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUserNotFound))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}
	refreshDataID, exists := c.Get("x-user-refresh-data-id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "x-user-refresh-data-id not found"})
        c.Abort()
        return
    }
    refreshDataIDStr, ok := refreshDataID.(string)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid x-user-refresh-data-id"})
        c.Abort()
        return
    }
	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour,refreshDataIDStr)


	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour,refreshDataIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}
	ID, err := primitive.ObjectIDFromHex(refreshDataIDStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}
	rtc.RefreshTokenUsecase.DeleteRefreshData(c,refreshDataIDStr)
	var refreshData entities.RefreshData
	refreshData.Id =ID
	refreshData.UserId = user.ID.Hex()
	
	refreshData.Revoked = false
	refreshData.Expire_date = refreshData.Expire_date
	refreshData.RefreshToken = refreshToken

	rtc.RefreshTokenUsecase.CreateRefreshData(c,refreshData)

	refreshTokenResponse := entities.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
