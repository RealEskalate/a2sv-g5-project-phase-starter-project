package services

import (
	"context"
	"time"
	"your_project/domain"
	"your_project/infrastructure/cloudinary"
)

type mediaService struct{}

func NewMediaService() domain.MediaUpload {
	return &mediaService{}
}

func (s *mediaService) FileUpload(file domain.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return cloudinary.UploadFile(ctx, file.File)
}

func (s *mediaService) RemoteUpload(url domain.Url) (string, error) {
	// Cloudinary doesn't directly support URL uploads in the same way,
	// so this example assumes the URL is used directly if needed.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return cloudinary.UploadFile(ctx, url.Url)
}
