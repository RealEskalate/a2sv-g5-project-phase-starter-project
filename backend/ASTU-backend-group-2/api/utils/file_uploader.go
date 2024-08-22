package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type FileUploader struct {
}

func (f *FileUploader) UploadImgFile(ctx *gin.Context, fileName string) (string, error) {

	file, header, err := ctx.Request.FormFile(fileName)
	if err != nil {
		log.Println("error while reading file", err.Error())
		return "", errors.New("error while reading file")
	}
	ext := filepath.Ext(header.Filename)
	if ext == "" || ext != ".jpg" && ext != ".png" {
		return "", errors.New("the file is not an image")
	}
	defer file.Close()
	if header.Size > 3<<20 { // max size accepted 3MB
		return "", errors.New("file size too big")
	}

	return f.SaveFile(file, header)

}

type UploadConfig struct {
	Type    string
	Name    string
	Dir     string
	Ext     []string
	Maxsize int64 //in MB
}

// Valid checks if the provided file header is valid according to the upload configuration.
func (u UploadConfig) Valid(header multipart.FileHeader) error {
	//check file type
	ok := false
	for _, ext := range u.Ext {
		if ext == filepath.Ext(header.Filename) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("the file is not type %s", u.Type)
	} else if header.Size <= u.Maxsize {
		return errors.New("file size too big")
	}

	return nil
}
func (f *FileUploader) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {

	imgDir := "./assets/images"

	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		os.MkdirAll(imgDir, os.ModePerm)
	}
	ext := filepath.Ext(header.Filename) // returns .<extension> eg: .jpg
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dest, err := os.Create(imgDir + "/" + filename)
	if err != nil {
		return "", errors.New("failed to create file")
	}
	defer dest.Close()
	if _, err := io.Copy(dest, file); err != nil {
		return "", errors.New("failed to save file")
	}
	return filename, nil

}
