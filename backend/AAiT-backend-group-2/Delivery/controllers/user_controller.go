package controllers

import (
	domain "AAiT-backend-group-2/Domain"
	"AAiT-backend-group-2/Infrastructure/dtos"
	"AAiT-backend-group-2/Infrastructure/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

func (ctr *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctr.UserUsecase.GetAllUsers(c)

	if err != nil {
		c.IndentedJSON(utils.GetHttpErrorCodes(err), gin.H{
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
		c.IndentedJSON(utils.GetHttpErrorCodes(err), gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
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
		c.IndentedJSON(utils.GetHttpErrorCodes(err), gin.H{
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
		c.IndentedJSON(utils.GetHttpErrorCodes(err), gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User deleted successfully",
	})
}


func (ctr *UserController) PromoteUser(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.PromoteUser(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully!"})
}

func (ctr *UserController) DemoteAdmin(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.DemoteAdmin(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User demoted successfully!"})
}

func (ctr *UserController) Logout(c *gin.Context) {
	id := c.Param("id")

	err := ctr.UserUsecase.Logout(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged out successfully!"})
}

func (ctr *UserController) ForgotPassword(c *gin.Context) {
	var forgotPasswordDto dtos.ForgotPasswordDto

	if err := c.ShouldBindJSON(&forgotPasswordDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	id, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	idStr, ok := id.(string)
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	err := ctr.UserUsecase.ForgotPassword(c,  idStr, forgotPasswordDto.Email)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Password reset link sent to your email!"})
}

func (ctr *UserController) ResetPassword(c *gin.Context) {
	var passwordResetDto dtos.PasswordResetDto

	if err := c.ShouldBindJSON(&passwordResetDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, exists := c.Get("userID")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	idStr, ok := id.(string)
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	err := ctr.UserUsecase.ResetPassword(c, idStr, &passwordResetDto)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Password reset successfully!"})
}

func (ctr *UserController) ChangePassword(c *gin.Context){
	var changePasswordDto dtos.ChangePasswordDto

	if err := c.ShouldBindJSON(&changePasswordDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userID")
	if !exists{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	err := ctr.UserUsecase.ChangePassword(c, userIdStr, &changePasswordDto)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Password changed successfully!"})
}

func (ctr *UserController) UpdateProfile(c *gin.Context) {
	var updateProfileDto dtos.UpdateProfileDto

	if err := c.ShouldBindWith(&updateProfileDto, binding.FormMultipart); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userID")
	if !exists{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	err := ctr.UserUsecase.UpdateProfile(c, userIdStr, &updateProfileDto)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Profile updated successfully!"})
}