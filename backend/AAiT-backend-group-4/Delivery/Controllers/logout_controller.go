package controllers

import (
	"net/http"
	"aait-backend-group4/Domain"
	"github.com/gin-gonic/gin"
)

type LogOutController struct {
	LogoutUsecase domain.LogoutUsecase
}

// func NewLogOutController(logoutUsecase domain.LogoutUsecase) *LogOutController {
// 	return &LogOutController{
// 		LogoutUsecase: logoutUsecase,
// 	}
// }

func (c *LogOutController)Logout(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Token is required"})
		return
	}


	response, err := c.LogoutUsecase.Logout(ctx, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
