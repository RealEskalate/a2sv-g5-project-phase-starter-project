package controllers

import (
	"group3-blogApi/domain"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetMyProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	user, err := uc.UserUsecase.GetMyProfile(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}


func(uc* UserController) GetUsers(c *gin.Context){
	Role := c.GetString("role")
	if Role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	users, err := uc.UserUsecase.GetUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func(uc* UserController) GetUser(c *gin.Context){
	var user domain.User
	Role := c.GetString("role")
	userID := c.Param("id")
	user, err := uc.UserUsecase.GetMyProfile(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	checkUserID := user.ID.Hex()

	if Role != "admin"  || userID != checkUserID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	
	c.JSON(200, user)
}


func(uc* UserController) DeleteUser(c *gin.Context){
	var user domain.User
	Role := c.GetString("role")
	userID := c.Param("id")
	user, err := uc.UserUsecase.GetMyProfile(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	checkUserID := user.ID.Hex()

	if Role != "admin"  || userID != checkUserID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	
	deletedUser, err := uc.UserUsecase.DeleteUser(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully", "user": deletedUser})
}


func(uc* UserController) UpdateUserRole(c *gin.Context){
	Role := c.GetString("role")
	userID := c.Param("id")
	if Role != "admin" {	
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var role struct {
		Role string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
		

	updatedUser, err := uc.UserUsecase.UpdateUserRole(userID, role.Role)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedUser.Role = role.Role
	c.JSON(200, gin.H{"message": "User role updated successfully", "user": updatedUser})
}


func(uc* UserController) DeleteMyAccount(c *gin.Context){
	userID := c.GetString("user_id")

	user, err := uc.UserUsecase.GetMyProfile(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	if user.ID.Hex() != userID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}


	err = uc.UserUsecase.DeleteMyAccount(userID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Account deleted successfully"})
}


func (uc *UserController) UploadImage(c *gin.Context) {
	var Image struct {
		Image string `json:"image" binding:"required"`
	}

	
	userID := c.GetString("user_id")

	user, err := uc.UserUsecase.GetMyProfile(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	if user.ID.Hex() != userID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&Image); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	err = uc.UserUsecase.UploadImage(userID, Image.Image)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Image uploaded successfully"})
}

func (uc *UserController) UpdateMyProfile(c *gin.Context) {
	var user domain.User
	userID := c.GetString("user_id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.UserUsecase.UpdateMyProfile(user, userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Profile updated successfully"})
}