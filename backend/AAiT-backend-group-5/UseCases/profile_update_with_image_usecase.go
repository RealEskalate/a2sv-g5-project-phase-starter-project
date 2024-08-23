package usecases

import (
	"context"
	"mime/multipart"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type profileUsecase struct {
	profileRepo interfaces.ProfileUpdateRepository
	cloudinary  *cloudinary.Cloudinary
}

func NewProfileUsecase(pr interfaces.ProfileUpdateRepository, cld *cloudinary.Cloudinary) interfaces.ProfileUpdateUsecase {
	return &profileUsecase{
		profileRepo: pr,
		cloudinary:  cld,
	}
}

func (pu *profileUsecase) UploadImageToCloudinary(userId string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uploadResult, err := pu.cloudinary.Upload.Upload(context.TODO(), src, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	if err := pu.profileRepo.SaveProfileImageKey(userId, uploadResult.PublicID); err != nil {
		return "", err
	}

	return uploadResult.PublicID, nil
}
