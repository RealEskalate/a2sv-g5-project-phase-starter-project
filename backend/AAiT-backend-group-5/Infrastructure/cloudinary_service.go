package infrastructure

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cloudinary *cloudinary.Cloudinary
	env        *config.Env
}

func NewCloudinaryService(cloudinary *cloudinary.Cloudinary, env *config.Env) interfaces.CloudinaryInterface {
	return &CloudinaryService{
		cloudinary: cloudinary,
		env:        env,
	}
}

func (uc *CloudinaryService) UploadFile(file multipart.FileHeader, ctx context.Context) (string, *models.ErrorResponse) {

	src, e := file.Open()
	if e != nil {
		return "", models.BadRequest("invalid file")
	}
	defer src.Close()

	// Validate the file type
	buffer := make([]byte, 512) // Read the first 512 bytes to determine the file type
	if _, e := src.Read(buffer); e != nil {
		return "", models.BadRequest("invalid file")
	}

	// Reset the read pointer after reading
	src.Seek(0, 0)

	// Get the content type
	contentType := http.DetectContentType(buffer)

	// Check if the content type is an image
	if !strings.HasPrefix(contentType, "image/") {
		return "", models.BadRequest("only image files are allowed")
	}

	uploadResult, e := uc.cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{ResourceType: "image"})
	if e != nil {
		return "", models.InternalServerError("image upload failed")
	}

	return uploadResult.PublicID, nil
}

func (uc *CloudinaryService) DeleteFile(publicID string, ctx context.Context) *models.ErrorResponse {
	_, e := uc.cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})
	if e != nil {
		return models.InternalServerError("image deletion failed")
	}
	return nil
}

func (uc *CloudinaryService) GetProfileImageURL(publicID string, c context.Context) string {
	cloudName := uc.env.CLOUDINARY_CLOUD_NAME
	imageURL := fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", cloudName, publicID)
	return imageURL
}
