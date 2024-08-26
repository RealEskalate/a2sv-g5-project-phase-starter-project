package controllers

import (
	"group3-blogApi/domain"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetMyProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	user, uerr := uc.UserUsecase.GetMyProfile(userID)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}
	c.JSON(200, gin.H{"message": "Welcome to your profile", "user": user})
}


func(uc* UserController) GetUsers(c *gin.Context){
	Role := c.GetString("role")
	if Role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	users, uerr := uc.UserUsecase.GetUsers()
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

func(uc* UserController) GetUser(c *gin.Context){
	var user domain.User
	Role := c.GetString("role")
	userID := c.Param("id")
	user, uerr := uc.UserUsecase.GetMyProfile(userID)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
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
	user, uerr := uc.UserUsecase.GetMyProfile(userID)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}

	checkUserID := user.ID.Hex()

	if Role != "admin"  || userID != checkUserID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	
	deletedUser, uerr := uc.UserUsecase.DeleteUser(userID)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
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
		

	updatedUser, uerr := uc.UserUsecase.UpdateUserRole(userID, role.Role)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}
	updatedUser.Role = role.Role
	c.JSON(200, gin.H{"message": "User role updated successfully", "user": updatedUser})
}


func(uc* UserController) DeleteMyAccount(c *gin.Context){
	userID := c.GetString("user_id")

	user, uerr := uc.UserUsecase.GetMyProfile(userID)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}

	if user.ID.Hex() != userID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if uerr = uc.UserUsecase.DeleteMyAccount(userID); uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}
	c.JSON(200, gin.H{"message": "Account deleted successfully"})
}


func (uc *UserController) UploadImage(c *gin.Context) {
	var Image struct {
		Image string `json:"image" binding:"required"`
	}

	
	userID := c.GetString("user_id")

	user, uerr := uc.UserUsecase.GetMyProfile(userID)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
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


	if uerr = uc.UserUsecase.UploadImage(userID, Image.Image); uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
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

	if uerr := uc.UserUsecase.UpdateMyProfile(user, userID); uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}
	c.JSON(200, gin.H{"message": "Profile updated successfully"})
}


