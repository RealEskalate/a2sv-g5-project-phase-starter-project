package controllers

import (
	"mime/multipart"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type UserProfileController struct {
	UserProfileUC interfaces.UserProfileUpdateUsecase
}

func NewUserProfileController(userProfileUC interfaces.UserProfileUpdateUsecase) *UserProfileController {
	return &UserProfileController{
		UserProfileUC: userProfileUC,
	}
}

func (userProfileController *UserProfileController) ProfileUpdate(ctx *gin.Context) {
	var updatedUser dtos.ProfileUpdateRequest

	userID := ctx.GetString("id")

	file, err := ctx.FormFile("profileImage")
	if err != nil {
		file = &multipart.FileHeader{}
	}

	updatedUser.Username = ctx.PostForm("user_name")
	updatedUser.Name = ctx.PostForm("name")
	updatedUser.PhoneNumber = ctx.PostForm("phone_number")
	updatedUser.Password = ctx.PostForm("password")
	updatedUser.Bio = ctx.PostForm("bio")

	if err := updatedUser.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}

	e := userProfileController.UserProfileUC.UpdateUserProfile(ctx, userID, &updatedUser, file)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user profile successfully updated"})
}

func (userProfileController *UserProfileController) ProfileGet(ctx *gin.Context) {
	userID := ctx.GetString("id")
	user, e := userProfileController.UserProfileUC.GetUserProfile(ctx, userID)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (userProfileController UserProfileController) ProfileDelete(ctx *gin.Context) {
	userID := ctx.GetString("id")
	e := userProfileController.UserProfileUC.DeleteUserProfile(ctx, userID)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user profile successfully deleted"})
}
