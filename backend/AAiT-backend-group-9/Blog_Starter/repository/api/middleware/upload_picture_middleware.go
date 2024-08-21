package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// uplpadPictureMiddleware will save an image to upload folder and save the path on the context for the next handler am using gin web freamwork
func UploadPictureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		fileName := strings.Split(file.Filename, ".")
		filePath := "upload/" + fileName[0] + "." + fileName[1]
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("profile_picture", filePath)
		c.Next()
	}
}
