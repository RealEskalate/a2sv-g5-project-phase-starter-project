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
	promote := struct {
		ID uuid.UUID `json:"id"`
		IsPromote bool `json:"isPromote"`
	}{}
	if err := c.BindJSON(&promote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cerr := u.userUseCase.PromoteUser(promote.ID, promote.IsPromote)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})	
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
	cerr := u.userUseCase.UpdateUser(&user)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}