package infrastructure

import (
	"context"
	"mime/multipart"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinaryService(cloudinary *cloudinary.Cloudinary) interfaces.CloudinaryInterface {
	return &CloudinaryService{
		cloudinary: cloudinary,
	}
}

func (uc *CloudinaryService) UploadFile(file multipart.FileHeader, ctx context.Context) (string, *models.ErrorResponse) {

	src, e := file.Open()
	if e != nil {
		return "", models.BadRequest("invalid file")
	}
	defer src.Close()

	uploadResult, e := uc.cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{})
	if e != nil {
		return "", models.InternalServerError("image upload failed")
	}

	return uploadResult.PublicID, nil
}
