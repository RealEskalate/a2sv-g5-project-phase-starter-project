package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func IsValidFileFormat(header *multipart.FileHeader, formats ...string) bool {
	fileFormat := header.Header.Get("Content-Type")
	for _, file := range formats {
		if file == fileFormat {
			return true
		}
	}
	return false
}

func SaveImage(file multipart.File, name string, cxt context.Context) (string, error) {
	cld, err := setupCloudinary()
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(cxt, file, uploader.UploadParams{
		PublicID: name,
		Folder:   "a2sv_blog_project",
	})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}

func setupCloudinary() (*cloudinary.Cloudinary, error) {
	cld, errCld := cloudinary.NewFromURL(fmt.Sprintf("%v%v:%v@%v", os.Getenv("CLOUDINARY_URL"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"), os.Getenv("CLOUDINARY_NAME")))
	if errCld != nil {
		return nil, errCld
	}
	return cld, nil
}
