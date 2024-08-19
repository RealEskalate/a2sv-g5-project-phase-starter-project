package controllers

import (
	"blog_project/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase domain.IUserUsecase
}

func NewUserController(userUsecase domain.IUserUsecase) domain.IUserController {
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

func (uc *userController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	user, err := uc.UserUsecase.GetUserByID(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (uc *userController) CreateUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.CreateUser(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user domain.User
	err = c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.UpdateUser(c, idInt, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = uc.UserUsecase.DeleteUser(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (uc *userController) AddBlog(c *gin.Context) {
	userID := c.Param("userID")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var Blog domain.Blog
	err = c.BindJSON(&Blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.AddBlog(c, userIDInt, Blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) Login(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uc.UserUsecase.Login(c, user.Username, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (uc *userController) Logout(c *gin.Context) {

}

func (uc *userController) ForgetPassword(c *gin.Context) {
	email := c.Param("email")
	err := uc.UserUsecase.ForgetPassword(c, email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset link sent to email"})
}

func (uc *userController) ResetPassword(c *gin.Context) {

	username := c.Param("username")
	password := c.Param("password")
	err := uc.UserUsecase.ResetPassword(c, username, password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset successfully"})
}

func (uc *userController) PromoteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.PromoteUser(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User promoted successfully", "user": user})
}

func (uc *userController) DemoteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.DemoteUser(c, idInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User demoted successfully", "user": user})
}
