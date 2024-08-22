package infrastructure

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

type MediaUpload interface {
	FileUpload(file File) (string, error)
}

type media struct{}

func NewMediaUpload() MediaUpload {
	return &media{}
}

func (*media) FileUpload(file File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}
