package controller

import (
	"net/http"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase entities.LogOutUsecase
	Env           *bootstrap.Env
}

func (lc *LogoutController) Logout(c *gin.Context) {
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
	lc.LogoutUsecase.LogOut(c,refreshDataIDStr)
}
