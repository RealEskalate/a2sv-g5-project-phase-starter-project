package utils

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Utils interface {
	GenerateTokenWithLength(legth int) (int, error)
	SaveImage(file multipart.File, name string, cxt context.Context) (*uploader.UploadResult, error)
	DeleteImage(publicID string, cxt context.Context) error
	SetupCloudinary() (*cloudinary.Cloudinary, error)
	IsValidFileFormat(header *multipart.FileHeader, formats ...string) bool
}

type UserUtils struct{}

func NewUserUtils() *UserUtils {
	return &UserUtils{}
}

func (userUT *UserUtils) GenerateTokenWithLength(length int) (int, error) {
	if length <= 0 {
		return 0, errors.New("length must be a positive integer")
	}

	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	code, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return int(code.Int64()), nil
}
