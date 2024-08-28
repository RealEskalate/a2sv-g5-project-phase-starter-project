package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (userUT *UserUtils) IsValidFileFormat(header *multipart.FileHeader, formats ...string) bool {
	fileFormat := header.Header.Get("Content-Type")
	for _, file := range formats {
		if file == fileFormat {
			return true
		}
	}
	return false
}

func (userUT *UserUtils) SaveImage(file multipart.File, name string, cxt context.Context) (*uploader.UploadResult, error) {
	cld, err := userUT.SetupCloudinary()
	if err != nil {
		return &uploader.UploadResult{}, err
	}

	uploadResult, err := cld.Upload.Upload(cxt, file, uploader.UploadParams{
		Folder:       "a2sv_blog_project",
		UploadPreset: "profile_picture",
	})
	if err != nil {
		return &uploader.UploadResult{}, err
	}
	return uploadResult, nil
}

func (userUT *UserUtils) DeleteImage(publicID string, cxt context.Context) error {
	cld, err := userUT.SetupCloudinary()
	if err != nil {
		return err
	}

	invalidate := true
	resp, err := cld.Upload.Destroy(cxt, uploader.DestroyParams{
		PublicID:   publicID,
		Invalidate: &invalidate,
	})

	if err != nil {
		return err
	}

	if resp.Result != "ok" {
		return fmt.Errorf("error deleting image")
	}
	return nil
}

func (userUT *UserUtils) SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cld, errCld := cloudinary.NewFromURL(fmt.Sprintf("%v%v:%v@%v", os.Getenv("CLOUDINARY_URL"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"), os.Getenv("CLOUDINARY_NAME")))
	if errCld != nil {
		return nil, errCld
	}
	return cld, nil
}
