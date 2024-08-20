package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "meleket/domain"
)

type RefreshTokenController struct {
    UserUsecase domain.UserUsecaseInterface
    RefreshTokenUsecase domain.RefreshTokenUsecaseInterface
}

func NewRefreshTokenController(uu domain.UserUsecaseInterface, rt domain.RefreshTokenUsecaseInterface) *RefreshTokenController {
    return &RefreshTokenController{
        UserUsecase: uu,
        RefreshTokenUsecase: rt,
    }
}

func (c *RefreshTokenController) RefreshToken(ctx *gin.Context) {
    var refreshT domain.RefreshToken
    if err := ctx.ShouldBindJSON(&refreshT); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newAccessToken, err := c.RefreshTokenUsecase.RefreshToken(&refreshT)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": newAccessToken})
}
