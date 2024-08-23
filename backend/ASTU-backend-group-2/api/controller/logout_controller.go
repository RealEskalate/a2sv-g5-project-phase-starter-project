package controller

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase domain.LoginUsecase
	Env           *bootstrap.Env
}

func (lc *LogoutController) Logout(c *gin.Context) {
	// // Get the user ID from the context
	// userID := c.GetString("x-user-id")

	// // Call the usecase
	// err := lc.LogoutUsecase.Logout(userID)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(200, gin.H{
	// 	"message": "Successfully logged out",
	// })
}
