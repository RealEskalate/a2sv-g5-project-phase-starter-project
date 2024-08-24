package services

import (
    "context"
    "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
    "log"
)

type ICloudinaryService interface{
	UploadProfilePicture(ctx context.Context, file interface{}) (string, error) 
}

type CloudinaryService struct {
    cld          *cloudinary.Cloudinary
    uploadFolder string
}

func NewCloudinaryService(cloudName, apiKey, apiSecret, uploadFolder string) ICloudinaryService {
    cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        log.Fatalf("Failed to initialize Cloudinary: %v", err)
    }
    return &CloudinaryService{cld: cld, uploadFolder: uploadFolder}
}

func (cs *CloudinaryService) UploadProfilePicture(ctx context.Context, file interface{}) (string, error) {
    uploadParams := uploader.UploadParams{Folder: cs.uploadFolder}
    result, err := cs.cld.Upload.Upload(ctx, file, uploadParams)
    if err != nil {
        return "", err
    }
    return result.SecureURL, nil
}