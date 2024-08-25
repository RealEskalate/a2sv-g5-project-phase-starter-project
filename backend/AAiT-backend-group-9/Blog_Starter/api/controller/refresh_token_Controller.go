package controller

import (
	"Blog_Starter/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
}

func NewRefreshTokenController(refreshTokenUsecase domain.RefreshTokenUsecase) *RefreshTokenController {
	return &RefreshTokenController{
		RefreshTokenUsecase: refreshTokenUsecase,
	}
}

func (rc *RefreshTokenController) RefreshToken(c *gin.Context) {
	userID := c.GetString("userID")

	var req domain.RefreshTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := rc.RefreshTokenUsecase.CheckRefreshToken(c, userID, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := rc.RefreshTokenUsecase.UpdateTokens(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
