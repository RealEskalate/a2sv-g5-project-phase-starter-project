package controllers

import (
	"blog_project/domain"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase domain.IUser_Usecases
}

func NewUserController(userUsecase domain.IUser_Usecases) domain.IUser_Controller {
	return &userController{UserUsecase: userUsecase}
}

func (uc *userController) GetAllUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

// func (uc *userController) GetUserByID(c *gin.Context) {
// 	id := c.Param("id")
// 	user, err := uc.UserUsecase.GetUserByID(id)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, user)
// }

func (uc *userController) CreateUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.UpdateUser(id, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}
