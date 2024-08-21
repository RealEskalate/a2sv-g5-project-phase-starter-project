package usecases

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
)

type IUserUseCase interface {
	CreateUser(user *domain.User) (*domain.User, *domain.CustomError)
	GetUserByID(id uuid.UUID) (*dto.GetUserResponseDto, *domain.CustomError)
	UpdateUser(requesterID uuid.UUID, user *dto.UserUpdate) *domain.CustomError
	PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError
	UploadProfilePic(userID uuid.UUID, file *multipart.File, header *multipart.FileHeader) (string, *domain.CustomError) // GetUserByName(name string) (*domain.User, *domain.CustomError)
}

type UserUseCase struct {
	userRepo interfaces.IUserRepository
}

func NewUserUseCase(repo interfaces.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, *domain.CustomError) {
	return user, u.userRepo.CreateUser(user)
}

func (u *UserUseCase) GetUserByID(id uuid.UUID) (*dto.GetUserResponseDto, *domain.CustomError) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.GetUserResponseDto{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Bio:      user.Bio,
		ImageUrl: user.ImageURL,
		IsAdmin:  user.IsAdmin,
	}, nil
}

func (u *UserUseCase) UpdateUser(requesterID uuid.UUID, user *dto.UserUpdate) *domain.CustomError {
	if requesterID != user.ID {
		return domain.ErrUnAuthorized
	}
	user.UpdatedAt = time.Now().UTC()
	return u.userRepo.UpdateUser(user)
}

func (u *UserUseCase) PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError {
	return u.userRepo.PromoteUser(id, isPromote)
}

func (uc *UserUseCase) UploadProfilePic(userID uuid.UUID, file *multipart.File, header *multipart.FileHeader) (string, *domain.CustomError) {

	CLOUDINARY_API_KEY := os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARY_API_SECRET := os.Getenv("CLOUDINARY_API_SECRET")
	CLOUDINARY_CLOUD_NAME := os.Getenv("CLOUDINARY_CLOUD_NAME")

	// Create the uploads directory if it doesn't exist
	uploadDir := "../uploads/"
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", &domain.CustomError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	// Generate a unique file name (you might want to use the user's ID or a UUID here)
	fileName := fmt.Sprintf("%s-%s", userID, filepath.Base(header.Filename))
	filePath := filepath.Join(uploadDir, fileName)

	// Save the file to the server
	out, err := os.Create(filePath)
	if err != nil {

		return "", &domain.CustomError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}
	defer out.Close()

	_, err = io.Copy(out, *file)
	if err != nil {
		os.Remove(out.Name())

		return "", &domain.CustomError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}
	cld, err := cloudinary.NewFromParams(CLOUDINARY_CLOUD_NAME, CLOUDINARY_API_KEY, CLOUDINARY_API_SECRET)
	if err != nil {

		return "", &domain.CustomError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	// Upload the file to Cloudinary
	uploadParams := uploader.UploadParams{
		PublicID: fmt.Sprintf("%s-%s", userID, filepath.Base(header.Filename)), // Generate a unique PublicID
		Folder:   "profile_pics/",                                              // Store in a specific folder in Cloudinary
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), out.Name(), uploadParams)
	if err != nil {
		os.Remove(out.Name())

		return "", &domain.CustomError{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}
	// Update the user's profile in the database with the file path
	// Assuming you have a function like UpdateUserProfilePic(userID, filePath string)
	user, errs := uc.userRepo.GetUserByID(userID)
	if errs != nil {
		os.Remove(out.Name())

		return "", errs
	}
	user.ImageURL = uploadResult.SecureURL
	errs = uc.userRepo.UpdateUserToken(user)
	if errs != nil {
		os.Remove(out.Name())

		cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: uploadParams.PublicID})
		return "", errs
	}
	os.Remove(out.Name())

	return uploadResult.SecureURL, nil
	// Respond with the file path or URL
}
