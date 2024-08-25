package controller

import (
	"AAiT-backend-group-6/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (ctr *ProfileController) UpdateProfile(c *gin.Context) {
	var updateProfileDto  domain.UpdateProfileDto

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

	err := ctr.ProfileUsecase.UpdateProfile(c, userIdStr, &updateProfileDto)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Profile updated successfully!"})
}