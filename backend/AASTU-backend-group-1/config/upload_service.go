package config

import (
	"blogs/bootstrap"
	"context"
	"log"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var GenerateUid = func() string {
	return time.Now().Format("20060102150405")
}

func UploadToCloudinary(file *multipart.FileHeader) (string, error) {
	// Create Cloudinary instance
	cloudinaryURL, err := bootstrap.GetEnv("CLOUDINARY_URL")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Upload the image to Cloudinary
	var ctx = context.Background()
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "my_avatar" + "-" + file.Filename + "-" + GenerateUid(),
	})
	if err != nil {
		log.Println("Failed to upload file to Cloudinary:", err)
		return "", err
	}

	return resp.SecureURL, nil
}
