package controller

import (
	"backend-starter-project/domain/dto"
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
	var response dto.Response
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		c.JSON(500,response)
		return
	}

	response.Success = true
	response.Message = "User promoted successfully"

	c.JSON(200, response)
}

func (uc *UserController) DemoteUser(c *gin.Context) {
	userId := c.Param("id")
	err := uc.userService.DemoteUserToRegular(userId)

	if err != nil {
		c.JSON(500, dto.Response{
			Success: false,
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, dto.Response{
		Success: true,
		Message: "User demoted successfully",
	})
}