package blog

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *BlogHandler) UploadBlogPhotos(c *gin.Context) {
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload files"})
        return
    }

    files := form.File["photos"]
    if len(files) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
        return
    }

    if len(files) > 5 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 5 photos allowed"})
        return
    }

    uploadPath := "uploads"
    if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
        if err := os.Mkdir(uploadPath, os.ModePerm); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload folder"})
            return
        }
    }

    var profilePicUrls []string
    for _, file := range files {
        if file.Size > 10*1024*1024 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum file size exceeded (10 MB)"})
            return
        }

        fileName := fmt.Sprintf("%s_%s", primitive.NewObjectID().Hex(), file.Filename)
        filePath := filepath.Join(uploadPath, fileName)

        out, err := os.Create(filePath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }
        defer out.Close()

        fileReader, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
            return
        }
        defer fileReader.Close()

        if _, err := io.Copy(out, fileReader); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }

        profilePicUrl := fmt.Sprintf("/%s/%s", uploadPath, fileName)
        profilePicUrls = append(profilePicUrls, profilePicUrl)
    }

    c.JSON(http.StatusOK, gin.H{"profile_urls": profilePicUrls})
}