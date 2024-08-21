package controller

import (
	"Blog_Starter/domain"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
}

func (lc *LogoutController) LogOut(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	userID := c.Param("user_id") // notparameter from middleware find the user_id
	err := lc.LogoutUsecase.LogOut(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot logout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}