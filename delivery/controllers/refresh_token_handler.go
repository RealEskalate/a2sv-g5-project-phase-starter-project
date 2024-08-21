package controllers

import (
	"meleket/domain"
	"meleket/infrastructure"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
    UserUsecase domain.UserUsecaseInterface
    RefreshTokenUsecase domain.RefreshTokenUsecaseInterface
    jwtsrv infrastructure.JWTService
}

func NewRefreshTokenController(uu domain.UserUsecaseInterface, rt domain.RefreshTokenUsecaseInterface, js infrastructure.JWTService) *RefreshTokenController {
    return &RefreshTokenController{
        UserUsecase: uu,
        RefreshTokenUsecase: rt,
        jwtsrv: js,
    }
}

func (c *RefreshTokenController) RefreshToken(ctx *gin.Context) {
    var refreshT domain.RefreshTokenRequest
    if err := ctx.ShouldBindJSON(&refreshT); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "err":"bind"})
        return
    }

    // secret_key := c.jwtsrv.GetKey("refresh")
    // fmt.Println(secret_key)
    // fmt.Println(refreshT)
    claim, err := c.jwtsrv.ValidateRefreshToken(refreshT.RefreshToken)
    if err != nil {
        // Debugging: Print the token expiry and current time
        // fmt.Printf("Current Time: %v, Token Expiry Time: %v\n", time.Now().Unix(), claim.ExpiresAt)
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "expired refresh token", 
            "err": err})
        return
    }


    newAccessToken, err := c.RefreshTokenUsecase.RefreshToken(claim.ID,claim.Role)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"token": newAccessToken})
}
