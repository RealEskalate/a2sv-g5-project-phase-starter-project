package controller

import (
	"blog/config"
	"blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *config.Env
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	claims := c.MustGet("claim").(*domain.JwtCustomClaims)
	var user domain.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	returnedUser, _ := uc.UserUsecase.GetUserByEmail(c, user.Email)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	returnedUser, _ = uc.UserUsecase.GetUserByUsername(c, user.Username)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	
	err := uc.UserUsecase.CreateUser(c, &user, claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// UpdateUser updates a user
func (uc *UserController) UpdateUser(c *gin.Context) {
	claims := c.MustGet("claim").(*domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingUser, _ := uc.UserUsecase.GetUserByID(c, objectID)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	returnedUser, _ := uc.UserUsecase.GetUserByEmail(c, user.Email)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	returnedUser, _ = uc.UserUsecase.GetUserByUsername(c, user.Username)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	resp, err := uc.UserUsecase.UpdateUser(c, &user, claims, objectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": resp})
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	claims := c.MustGet("claim").(*domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingUser, _ := uc.UserUsecase.GetUserByID(c, objectID)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	
	err = uc.UserUsecase.DeleteUser(c, objectID, claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUser gets a user
func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, _ := uc.UserUsecase.GetUserByID(c, objectID)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers gets all users
func (uc *UserController) GetUsers(c *gin.Context) {
	users, _ := uc.UserUsecase.GetAllUsers(c)
	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No users found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// PromoteUser promotes a user
func (uc *UserController) PromoteUser(c *gin.Context) {
	claims := c.MustGet("claim").(*domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingUser, _ := uc.UserUsecase.GetUserByID(c, objectID)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	err = uc.UserUsecase.PromoteUser(c, objectID, claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}

// DemoteUser demotes a user
func (uc *UserController) DemoteUser(c *gin.Context) {
	claims := c.MustGet("claim").(*domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingUser, _ := uc.UserUsecase.GetUserByID(c, objectID)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	err = uc.UserUsecase.DemoteUser(c, objectID, claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User demoted successfully"})
}
