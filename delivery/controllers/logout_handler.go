package controllers

import (
	// "fmt"
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive" // should be updated
)

type LogoutController struct {
	tokenUsecase domain.RefreshTokenUsecaseInterface
}

func NewLogoutController(tk domain.RefreshTokenUsecaseInterface) *LogoutController {
	return &LogoutController{tokenUsecase: tk}
}

func (lc *LogoutController) Logout(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	
	err := lc.tokenUsecase.DeleteRefreshToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}