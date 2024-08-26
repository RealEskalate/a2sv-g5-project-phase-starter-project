package controller

import (
	"blog/domain"
	"blog/internal/userutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
    LogoutUsecase domain.LogoutUsecase
}

// Logout handles user logout and invalidates tokens
func (lc *LogoutController) Logout(c *gin.Context) {
    var request domain.LogoutRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceFingerprint := userutil.GenerateDeviceFingerprint(ipAddress, userAgent)

    // Pass only the refresh token to the usecase
    err := lc.LogoutUsecase.Logout(c, request.RefreshToken, deviceFingerprint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
