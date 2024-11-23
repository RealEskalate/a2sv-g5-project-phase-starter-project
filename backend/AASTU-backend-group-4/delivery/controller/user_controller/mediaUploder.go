package user_controller

import (
	"blog-api/domain"
	infrastructure "blog-api/infrastructure/cloudinary"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateRequest struct {
	Firstname          string `json:"firstname" bson:"firstname"`
	Lastname           string `json:"lastname" bson:"lastname"`
	Username           string `json:"username" bson:"username"`
	Bio                string `json:"bio" bson:"bio"`
	ProfilePicture     string `json:"profile_picture" bson:"profile_picture"`
	ContactInformation string `json:"contact_information" bson:"contact_information"`
	RefreshToken       string `json:"refresh_token" bson:"refresh_token"`
}

func (bc *UserController) FileUpload(c *gin.Context) {
	// Upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		return
	}

	// Get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
		return
	}

	// Upload file and get URL
	uploadUrl, err := bc.Medcont.FileUpload(infrastructure.File{File: formFile})
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
		return
	}

	// Create and populate the UpdateRequest struct
	updateRequest := domain.UpdateRequest{
		ProfilePicture: uploadUrl, // Set the profile picture URL
	}

	// Update user's profile picture with the uploaded URL
	userID := c.GetString("user_id")
	userId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = bc.userUsecase.UpdateUser(context.Background(), userId, &updateRequest)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Failed to update profile picture: " + err.Error()},
			})
		return
	}

	// Return success response
	c.JSON(
		http.StatusOK,
		infrastructure.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}
