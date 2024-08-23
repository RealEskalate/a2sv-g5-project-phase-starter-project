package interfaces

import "mime/multipart"

type ProfileUpdateUsecase interface {
	UploadImageToCloudinary(userId string, file *multipart.FileHeader) (string, error)
}
