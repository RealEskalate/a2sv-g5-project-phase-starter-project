package account

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *UserHandler) UploadProfilePic(c *gin.Context) {
    file, header, err := c.Request.FormFile("profile_pic")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
        return
    }
    defer file.Close()

    uploadPath := "uploads"
    if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
        if err := os.Mkdir(uploadPath, os.ModePerm); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload folder"})
            return
        }
    }

    fileName := fmt.Sprintf("%s_%s", primitive.NewObjectID().Hex(), header.Filename)
    filePath := filepath.Join(uploadPath, fileName)

    out, err := os.Create(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }
    defer out.Close()

    if _, err := io.Copy(out, file); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    profilePicUrl := fmt.Sprintf("/%s/%s", uploadPath, fileName)
    c.JSON(http.StatusOK, gin.H{"profile_url": profilePicUrl})
}
