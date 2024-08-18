package controllers

import (
    // "net/http"
    "github.com/gin-gonic/gin"
    "meleket/domain"
)

type RefreshTokenController struct {
    UserUsecase domain.UserUsecaseInterface
}

func NewRefreshTokenController(uu domain.UserUsecaseInterface) *RefreshTokenController {
    return &RefreshTokenController{
        UserUsecase: uu,
    }
}

type RefreshTokenRequest struct {
    RefreshToken string `json:"refresh_token" binding:"required"`
}

func (c *RefreshTokenController) RefreshToken(ctx *gin.Context) {
    // var req RefreshTokenRequest
    // if err := ctx.ShouldBindJSON(&req); err != nil {
    //     ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    //     return
    // }

    // newToken, err := c.UserUsecase.RefreshToken(req.RefreshToken)
    // if err != nil {
    //     ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
    //     return
    // }

    // ctx.JSON(http.StatusOK, gin.H{"token": newToken})
}
