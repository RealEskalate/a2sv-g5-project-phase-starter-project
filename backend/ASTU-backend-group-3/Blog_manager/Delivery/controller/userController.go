package controller

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/Usecases"
	"net/http"
	"ASTU-backend-group-3/Blog_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related endpoints
type UserController struct {
	UserUsecase Usecases.UserUsecase
}

// NewUserController creates a new instance of UserController
func NewUserController(userUsecase Usecases.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

// Register handles user registration
func (uc *UserController) Register(c *gin.Context) {
	var input Domain.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uc.UserUsecase.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// UpdateUser handles updating user information
func (uc *UserController) UpdateUser(c *gin.Context) {
	username := c.Param("username")

	var input Domain.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.UpdateUser(username, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("username")

	err := uc.UserUsecase.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var input Domain.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	access_token, err := uc.UserUsecase.Login(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

<<<<<<<<< Temporary merge branch 1
    c.JSON(http.StatusOK, gin.H{"access_token": access_token})
}

func (uc *UserController) Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	username, err := infrastructure.ParseUsernameToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	err = uc.UserUsecase.Logout(username, tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logout failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
=========
	c.JSON(http.StatusOK, gin.H{"access_token": access_token})
}
>>>>>>>>> Temporary merge branch 2
