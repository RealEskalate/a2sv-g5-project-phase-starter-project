package Utils

import (
	"blogapp/Config"
	"blogapp/Domain"
	"context"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	// "errors"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadJSON(ctx *gin.Context) (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func ObjectIdToString(objID primitive.ObjectID) string {
	return primitive.ObjectID.Hex(objID)
}

func StringToObjectId(str string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}

// genreate slug from title
func GenerateSlug(title string) string {
	slug := title
	slug = strings.ToLower(slug)
	slug = strings.Replace(slug, " ", "-", -1)
	return slug
}

// is user author of post or admin
func IsAuthorOrAdmin(claim Domain.AccessClaims, authorID primitive.ObjectID) (bool, error) {

	if claim.Role == "admin" {
		return true, nil
	}

	if claim.ID == authorID {
		return true, nil
	}

	return false, nil
}

func getCurrentTimeString() string {
	// Get the current time
	currentTime := time.Now()

	// Format the time as a string
	timeString := currentTime.Format("2006-01-02 15:04:05")

	return timeString
}

func SetProfilePicture(file *multipart.FileHeader) (string, error) {
	// Supported file types
	supportedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !supportedTypes[ext] {
		log.Println("Unsupported file type:", ext)
		return "", http.ErrNotSupported
	}

	cld, err := cloudinary.NewFromURL("cloudinary://" + Config.Cloud_api_key + ":" + Config.Cloud_api_secret + "@dncnqaztp")
	if err != nil {
		log.Println("Cloudinary config error:", err)
		return "", err
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Upload parameters with dynamic public ID
	uploadParams := uploader.UploadParams{
		PublicID: "my_PP_" + file.Filename + getCurrentTimeString(),
	}

	resp, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		log.Println("Upload error:", err)
		return "", err
	}

	return resp.SecureURL, nil
}

// func ExtractUser(c *gin.Context) (Domain.OmitedUser, error) {
// 	userID, ok := c.Get("user_id")
// 	if !ok {
// 		return Domain.OmitedUser{}, errors.New("Failed to retrieve user ID")
// 	}
// 	UserobjectID, err := primitive.ObjectIDFromHex(userID.(string))
// 	if err != nil {
// 		return Domain.OmitedUser{}, errors.New("invalid user ID")
// 	}
// 	is_admin, ok := c.Get("is_admin")
// 	if !ok {

// 		return Domain.OmitedUser{}, errors.New("Failed to retrieve role")
// 	}

// 	return Domain.OmitedUser{
// 		ID: UserobjectID,
// 		Is_Admin: is_admin.(bool),
// 	}, nil
// }
