package infrastructure

import (
	"AAiT-backend-group-6/domain"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)



type ImageService struct {
	cloudinaryUrl string
}

func NewImageService(cloudinaryUrl string) *ImageService {
	return &ImageService{cloudinaryUrl: cloudinaryUrl}
}

func (service *ImageService) UploadImage(c context.Context, file *multipart.FileHeader) (string, error) { 
	defer func() {
		os.Remove("assets/uploads/" + file.Filename)
	} ()

	cld, err := cloudinary.NewFromURL(service.cloudinaryUrl)

	if err != nil {
		return "", err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(c, "assets/uploads/" + file.Filename, uploader.UploadParams{PublicID: "my_avatar" + "-" + file.Filename + "-" + id.String()})
	if err != nil{
		return "", err
	}

	return resp.SecureURL, nil
}

func (service *ImageService) SaveProfileImage(dto *domain.UpdateProfileDto) (string, error) {
	if dto.Avatar == nil {
		return "", fmt.Errorf("no image file provided")
	}

	// Open the file stream
	file, err := dto.Avatar.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Create a destination file on the server

	destinationPath := filepath.Join("assets/uploads/", dto.Avatar.Filename)
	out, err := os.Create(destinationPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer out.Close()

	// Copy the file content to the destination file
	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("failed to save image file: %v", err)
	}

	// Return the path where the file is saved
	return destinationPath, nil
}