package middlewares

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
		userDir := filepath.Join("./uploads", "profileImages")
		if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
			log.Println("Error creating user directory:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			c.Abort()
			return
		}

		// Generate a unique filename for the image
		uniqueId := uuid.New()
		fileExt := filepath.Ext(file.Filename)
		if fileExt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image file extension"})
			c.Abort()
			return
		}
		filename := fmt.Sprintf("%s%s", strings.Replace(uniqueId.String(), "-", "", -1), fileExt)

		// Save the image to the user's directory
		filePath := filepath.Join(userDir, filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			log.Println("Error saving the image:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			c.Abort()
			return
		}

		// Pass the image path to the next handler
		imagePath := filepath.Join("/uploads/profileImages", filename)
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
