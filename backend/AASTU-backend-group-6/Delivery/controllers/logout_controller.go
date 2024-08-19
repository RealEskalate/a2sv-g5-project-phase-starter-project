package controllers

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
	Env           *infrastructure.Config
}

func (l *LogoutController) Logout(c *gin.Context) {
	id := c.GetString("user_id")
	// TODO: get id from claims
	_, err := l.LogoutUsecase.CheckActiveUser(c, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found, page not found, login before logout"})
		return
	}
	user_agent := c.Request.UserAgent()
	err = l.LogoutUsecase.Logout(c, id, user_agent)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Logout success"})
}
