package infrastructure

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ImageUploadMiddleware checks and processes the uploaded image
func ImageUploadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username") // Assuming the username is passed as a URL parameter
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			c.Abort()
			return
		}

		file, err := c.FormFile("image")
		if err != nil {
			log.Println("Error retrieving the image:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image file"})
			c.Abort()
			return
		}

		// Validate image type (you can customize this list)
		allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
		if !isValidImageType(file, allowedTypes) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
			c.Abort()
			return
		}

		// Create a directory for the user if it doesn't exist
		userDir := filepath.Join("./images", username)
		if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
			log.Println("Error creating user directory:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			c.Abort()
			return
		}

		// Generate a unique filename for the image
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := strings.Split(file.Filename, ".")[1]
		image := fmt.Sprintf("%s.%s", filename, fileExt)

		// Save the image to the user's directory
		filePath := filepath.Join(userDir, image)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			log.Println("Error saving the image:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			c.Abort()
			return
		}

		// Pass the image path to the next handler
		imagePath := filepath.Join("/images", username, image)
		c.Set("profileImagePath", imagePath)

		c.Next()
	}
}

// isValidImageType checks if the file is of an allowed image type
func isValidImageType(file *multipart.FileHeader, allowedTypes []string) bool {
	for _, t := range allowedTypes {
		if file.Header.Get("Content-Type") == t {
			return true
		}
	}
	return false
}
