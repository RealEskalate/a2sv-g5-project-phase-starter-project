package infrastructure

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"strings"
)

func SaveAvatar(photo multipart.FileHeader) (string, error) {
	// Save the avatar to disk
	if photo.Size == 0 {
		return "", errors.New("No photo provided")
	}
	fmt.Println(photo.Header.Get("Content-Type"))

	if photo.Header.Get("Content-Type") != "image/jpeg" && photo.Header.Get("Content-Type") != "image/png" {
		return "", errors.New("Invalid file type")
	}

	file, err := photo.Open()
	if err != nil {
		return "", fmt.Errorf("Error opening photo: %w", err)
	}
	defer file.Close()

	// Generate a UUID for a unique filename
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", fmt.Errorf("Error generating UUID: %w", err)
	}

	//Get the extension of the file
	filetype := strings.Split(photo.Header.Get("Content-Type"), "/")[1]

	// Create a new file in the "uploads" directory with the UUID
	destination := "./uploads/avatars/" + strings.TrimSpace(string(uuid)) + "." + filetype
	address := "public/avatars/" + strings.TrimSpace(string(uuid)) + "." + filetype
	out, err := os.Create(destination)
	if err != nil {
		return "", fmt.Errorf("Error creating avatar file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("Error copying photo to destination: %w", err)
	}

	return address, nil
}
