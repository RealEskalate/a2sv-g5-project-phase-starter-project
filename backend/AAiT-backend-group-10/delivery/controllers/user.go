package controllers

import (
	"net/http"

	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userUseCase usecases.IUserUseCase
}

func NewUserController(u usecases.IUserUseCase) *UserController {
	return &UserController{
		userUseCase: u,
	}
}

func (u *UserController) PromoteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	makeAdmin := c.Query("makeAdmin")
	if makeAdmin != "true" && makeAdmin != "false" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid makeAdmin value"})
		return
	}

	err = u.userUseCase.PromoteUser(id, makeAdmin == "true")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User promoted successfully"})	
}

func (u *UserController) UpdateProfile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var user dto.UserUpdate
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	err = u.userUseCase.UpdateUser(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User profile updated successfully"})
}