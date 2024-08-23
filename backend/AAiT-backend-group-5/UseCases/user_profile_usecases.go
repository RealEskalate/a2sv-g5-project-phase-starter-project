package usecases

import (
	"context"
	"fmt"
	"mime/multipart"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type userProfileUpdateUsecase struct {
	UserRepository  interfaces.UserRepository
	PasswordService interfaces.PasswordService
	cloudinary      *cloudinary.Cloudinary
}

func NewUserProfileUpdateUsecase(UserRepository interfaces.UserRepository, PasswordService interfaces.PasswordService, cld *cloudinary.Cloudinary) interfaces.UserProfileUpdateUsecase {
	return &userProfileUpdateUsecase{
		UserRepository:  UserRepository,
		PasswordService: PasswordService,
		cloudinary:      cld,
	}
}

func (uc *userProfileUpdateUsecase) UpdateUserProfile(ctx context.Context, userID string, user *dtos.ProfileUpdateRequest, file *multipart.FileHeader) *models.ErrorResponse {
	var updatedUser models.User
	fmt.Println("You were here")

	// Check if the user exists
	_, err := uc.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// handle password update
	if user.Password != "" {
		if err := uc.PasswordService.ValidatePasswordStrength(user.Password); err != nil {
			return err
		}

		hashedPassword, err := uc.PasswordService.EncryptPassword(user.Password)
		if err != nil {
			return models.InternalServerError("Something went wrong")

		}

		// set the password field of user with a hashed password
		updatedUser.Password = hashedPassword
	}

	// handel profile picture update
	if file.Filename != "" {
		src, e := file.Open()
		if e != nil {
			return models.BadRequest("")
		}
		defer src.Close()

		// upload the image to Cloudinary
		uploadResult, e := uc.cloudinary.Upload.Upload(context.TODO(), src, uploader.UploadParams{})
		if e != nil {
			return models.InternalServerError("image upload failed")
		}

		// set the image public key to 'ImageKey' of user
		updatedUser.ImageKey = uploadResult.PublicID
	}

	updatedUser.Name = user.Name
	updatedUser.Bio = user.Bio
	updatedUser.PhoneNumber = user.PhoneNumber

	fmt.Println(updatedUser)
	err = uc.UserRepository.UpdateUser(ctx, &updatedUser, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userProfileUpdateUsecase) GetUserProfile(ctx context.Context, userID string) (*models.User, *models.ErrorResponse) {
	user, err := uc.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userProfileUpdateUsecase) DeleteUserProfile(ctx context.Context, userID string) *models.ErrorResponse {

	if _, err := uc.UserRepository.GetUserByID(ctx, userID); err != nil {
		return err
	}

	err := uc.UserRepository.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
