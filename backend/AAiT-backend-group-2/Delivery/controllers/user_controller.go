package controllers

import (
	domain "AAiT-backend-group-2/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) UserController {
	return UserController{
		UserUsecase: userUsecase,
	}
}

func (ctr *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctr.UserUsecase.GetAllUsers(c)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (ctr *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := ctr.UserUsecase.GetUserByID(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "User Not Found!",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func (ctr *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := ctr.UserUsecase.CreateUser(c, user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User Registerd successfully!"})
}

func (ctr *UserController) UpdateUser(c *gin.Context) {
	var user domain.User

	id := c.Param("id")

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err := ctr.UserUsecase.UpdateUser(c, id, &user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func (ctr *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.DeleteUser(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User deleted successfully",
	})
}

func (ctr *UserController) Login(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	token, err := ctr.UserUsecase.Login(c, &user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"token": token,
	})
}


func (ctr *UserController) PromoteUser(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.PromoteUser(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully!"})
}

func (ctr *UserController) DemoteAdmin(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.DemoteAdmin(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User demoted successfully!"})
}
