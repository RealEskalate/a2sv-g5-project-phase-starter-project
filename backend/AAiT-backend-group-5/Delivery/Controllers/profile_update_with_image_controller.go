package controllers

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUsecase interfaces.ProfileUpdateUsecase
}

func NewProfileController(pu interfaces.ProfileUpdateUsecase) *ProfileController {
	return &ProfileController{
		profileUsecase: pu,
	}
}

func (pc *ProfileController) UploadProfileImage(c *gin.Context) {
	file, err := c.FormFile("profileImage")
	userId := c.GetString("id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	imageKey, err := pc.profileUsecase.UploadImageToCloudinary(userId, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_key": imageKey})
}
