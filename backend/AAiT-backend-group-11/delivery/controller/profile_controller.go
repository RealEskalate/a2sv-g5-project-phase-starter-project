package controller

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileService interfaces.ProfileService
}

func NewProfileController(service interfaces.ProfileService) ProfileController {
	return ProfileController{ProfileService: service}
}

func (controller *ProfileController) CreateUserProfile(ctx *gin.Context) {
	var profile dto.CreateProfileDto
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetString("userId")
	if userID == "" {
		ctx.JSON(400, gin.H{"error": "user id is required"})
		return
	}
	profile.UserID = userID
	
	profile_, err := controller.ProfileService.CreateUserProfile(&profile)
	if err!=nil{
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile created successfully", "profile": profile_})

}

func (controller *ProfileController) GetUserProfile(ctx *gin.Context) {
	userId:=ctx.Param("userId")
	profile, err := controller.ProfileService.GetUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, profile)
}

func (controller *ProfileController) UpdateUserProfile(ctx *gin.Context) {
	var profile dto.UpdateProfileDto
	err := ctx.ShouldBindJSON(&profile)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updUserId:=ctx.Param("userId")

	UserID := ctx.GetString("userId")
	if UserID!=updUserId {
		ctx.JSON(400, gin.H{"error": "You are only authorized to update your own profile"})
		return
	}
	profile.UserID = UserID
	updated, err := controller.ProfileService.UpdateUserProfile(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile updated successfully", "updated profile": updated})
}

func (controller *ProfileController) DeleteUserProfile(ctx *gin.Context) {
	delId:=ctx.Param("userId")
	userId := ctx.GetString("userId")
	if userId!=delId {
		ctx.JSON(400, gin.H{"error": "You are only authorized to delete your own profile"})
		return
	}
	err := controller.ProfileService.DeleteUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile deleted successfully"})
}
