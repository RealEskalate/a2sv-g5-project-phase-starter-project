package interfaces

import (
	"context"
	"mime/multipart"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type CloudinaryInterface interface {
	UploadFile(file multipart.FileHeader, ctx context.Context) (string, *models.ErrorResponse)
}