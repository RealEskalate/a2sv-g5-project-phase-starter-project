package user_controller

import (
	"blog-api/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *UserController) UpdateUser(c *gin.Context) {
	var updateRequest domain.UpdateRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userIDParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	email := c.Param("email")

	user_, err := uc.userUsecase.GetByEmail(c, email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if updateRequest.Firstname == "" {
		updateRequest.Firstname = user_.Firstname
	}
	if updateRequest.Lastname == "" {
		updateRequest.Lastname = user_.Lastname
	}
	if updateRequest.Username == "" {
		updateRequest.Username = user_.Username
	}
	if updateRequest.Bio == "" {
		updateRequest.Bio = user_.Bio
	}
	if updateRequest.ContactInformation == "" {
		updateRequest.ContactInformation = user_.ContactInformation
	}
	if updateRequest.ProfilePicture == "" {
		updateRequest.ProfilePicture = user_.ProfilePicture
	}

	if err := uc.userUsecase.UpdateUser(context.Background(), userID, &updateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}
	c.Set("username", updateRequest.Username)
	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}
