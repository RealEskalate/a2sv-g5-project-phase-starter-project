package controllers

import (
	"net/http"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type ProfileUpdateController struct {
	ProfileUsecase interfaces.ProfileUpdateUsecase
}

func NewProfileController(pu interfaces.ProfileUpdateUsecase) *ProfileUpdateController {
	return &ProfileUpdateController{
		ProfileUsecase: pu,
	}
}

func (pc *ProfileUpdateController) UploadProfileImage(c *gin.Context) {
	file, err := c.FormFile("profileImage")
	userId := c.GetString("id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	imageKey, err := pc.ProfileUsecase.UploadImageToCloudinary(userId, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_key": imageKey})
}
