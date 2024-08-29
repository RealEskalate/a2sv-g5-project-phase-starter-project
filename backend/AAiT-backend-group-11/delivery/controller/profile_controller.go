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
func (controller *ProfileController) GetAllProfiles(ctx *gin.Context) {
	profiles, err := controller.ProfileService.GetAllProfiles()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, profiles)
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
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile created successfully", "profile": profile_})

}
func (controller *ProfileController) GetUserProfile(ctx *gin.Context) {
	userId := ctx.Param("userId")
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
	updUserId := ctx.Param("userId")

	UserID := ctx.GetString("userId")
	if UserID != updUserId {
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
	delId := ctx.Param("userId")
	userId := ctx.GetString("userId")
	if userId != delId {

		ctx.JSON(400, gin.H{"error": "You are only authorized to delete your own profile", "userId": userId, "delId": delId})
		return
	}
	err := controller.ProfileService.DeleteUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile deleted successfully"})
}

func (controller *ProfileController) UpdateOrCreateProfilePicture(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	url, err := controller.ProfileService.UpdateProfilePicture(userId, image)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile picture updated successfully", "url": url})

}

func (controller *ProfileController) GetProfilePicture(ctx *gin.Context) {
	userId := ctx.Param("userId")
	url, err := controller.ProfileService.GetProfilePicture(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"url": url})
}

func (controller *ProfileController) DeleteProfilePicture(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	user_id := ctx.Param("userId")
	if userId != user_id {
		ctx.JSON(400, gin.H{"error": "You cann't delete others profile picture", "userId": userId, "user_id": user_id})
		return
	}
	err := controller.ProfileService.DeleteProfilePicture(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile picture deleted successfully"})
}
