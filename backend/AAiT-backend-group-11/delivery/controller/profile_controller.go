package controller

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileService interfaces.ProfileService
}

func NewProfileController(service interfaces.ProfileService) ProfileController {
	return ProfileController{ProfileService: service}
}

func NewProfileController(ps interfaces.ProfileService ) *ProfileController {
		return &ProfileController{ profileService: ps }
	}


func (controller *ProfileController) CreateUserProfile(ctx *gin.Context) {
	var profile entities.Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile created successfully"})

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
	var profile entities.Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile updated successfully"})
}

func (controller *ProfileController) DeleteUserProfile(ctx *gin.Context) {
	userId := ctx.Param("userId")
	err := controller.ProfileService.DeleteUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile deleted successfully"})
}