package assetutil

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file multipart.File, fileName string, cloudinary *cloudinary.Cloudinary) (string, error) {

	resp, err := cloudinary.Upload.Upload(context.Background(), file, uploader.UploadParams{PublicID: fileName})

	if err != nil {
		return "", errors.New("failed to upload file")
	}

	return resp.SecureURL, nil

	// return f.SaveFile(file, header)
}
