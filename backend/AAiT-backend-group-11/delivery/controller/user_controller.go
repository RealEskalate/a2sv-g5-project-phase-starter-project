package controller

import (
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}


func (uc *UserController) PromoteUser(c *gin.Context) {
	userId := c.Param("id")
	err := uc.userService.PromoteUserToAdmin(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User promoted successfully"})
}

func (uc *UserController) DemoteUser(c *gin.Context) {
	userId := c.Param("id")
	err := uc.userService.DemoteUserToRegular(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User demoted successfully"})
}