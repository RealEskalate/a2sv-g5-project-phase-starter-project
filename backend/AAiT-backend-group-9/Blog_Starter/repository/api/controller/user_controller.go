package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if user.Role != "superAdmin" && user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Only an admin or super-admin can get all users information"})
		return
	}

	users, err := uc.userUsecase.GetAllUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if user.Role != "superAdmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Only the super-admin can promote a user"})
		return
	}

	userID := c.Param("id")

	err = uc.userUsecase.PromoteUser(c, userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

func (uc *UserController) DemoteUser(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if user.Role != "superAdmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Only the super-admin can demote a user"})
		return
	}

	userID := c.Param("id")

	err = uc.userUsecase.DemoteUser(c, userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User demoted to user"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var password string

	if err := c.ShouldBindJSON(&password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if user.Role != "user" && user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Only a user or admin can delete their account"})
		return
	}

	err = uc.userUsecase.DeleteUser(c, user.UserID, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}

func (uc *UserController) UpdateUser(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var userUpdate domain.UserUpdate

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updatedUser, err := uc.userUsecase.UpdateUser(c, &userUpdate, user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

func (uc *UserController) UpdateProfilePicture(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	profilePicPath, ok := c.Get("profile_picture")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updatedUser, err := uc.userUsecase.UpdateProfilePicture(c, profilePicPath.(string), user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

func (uc *UserController) DeleteProfilePicture(c *gin.Context) {
	// Get authenticated user from gin context
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := uc.userUsecase.UpdateProfilePicture(c, "", user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}
