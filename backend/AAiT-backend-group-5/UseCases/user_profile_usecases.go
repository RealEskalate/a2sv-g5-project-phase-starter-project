package usecases

import (
	"context"
	"fmt"
	"mime/multipart"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type userProfileUpdateUsecase struct {
	UserRepository    interfaces.UserRepository
	PasswordService   interfaces.PasswordService
	sessionRepo       interfaces.SessionRepository
	cloudinaryService interfaces.CloudinaryInterface
}

func NewUserProfileUpdateUsecase(
	UserRepository interfaces.UserRepository,
	PasswordService interfaces.PasswordService,
	cloudinaryService interfaces.CloudinaryInterface,
	sessionRepo interfaces.SessionRepository,
) interfaces.UserProfileUpdateUsecase {
	return &userProfileUpdateUsecase{
		UserRepository:    UserRepository,
		PasswordService:   PasswordService,
		cloudinaryService: cloudinaryService,
		sessionRepo:       sessionRepo,
	}
}

func (uc *userProfileUpdateUsecase) UpdateUserProfile(ctx context.Context, userID string, user *dtos.ProfileUpdateRequest, file *multipart.FileHeader) *models.ErrorResponse {
	var updatedUser models.User

	currUser, err := uc.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.Password != "" {
		if err := uc.PasswordService.ValidatePasswordStrength(user.Password); err != nil {
			return err
		}

		hashedPassword, err := uc.PasswordService.EncryptPassword(user.Password)
		if err != nil {
			return models.InternalServerError("Something went wrong")

		}

		updatedUser.Password = hashedPassword
	}

	if file.Filename != "" {
		uploadResult, err := uc.cloudinaryService.UploadFile(*file, ctx)
		if err != nil {
			return err
		}

		fmt.Println(currUser.ImageKey)
		err = uc.cloudinaryService.DeleteFile(currUser.ImageKey, ctx)
		if err != nil {
			return err
		}

		updatedUser.ImageKey = uploadResult
	}

	updatedUser.Username = user.Username
	updatedUser.Name = user.Name
	updatedUser.Bio = user.Bio
	updatedUser.PhoneNumber = user.PhoneNumber

	err = uc.UserRepository.UpdateUser(ctx, &updatedUser, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userProfileUpdateUsecase) GetUserProfile(ctx context.Context, userID string) (*dtos.Profile, *models.ErrorResponse) {
	user, err := uc.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	imageURL := uc.cloudinaryService.GetProfileImageURL(user.ImageKey, ctx)
	userResponse := dtos.Profile{
		Email:       user.Email,
		Username:    user.Username,
		Name:        user.Name,
		Bio:         user.Bio,
		PhoneNumber: user.PhoneNumber,
		ImageURL:    imageURL,
	}

	return &userResponse, nil
}

func (uc *userProfileUpdateUsecase) DeleteUserProfile(ctx context.Context, userID string) *models.ErrorResponse {

	if _, err := uc.UserRepository.GetUserByID(ctx, userID); err != nil {
		return err
	}
	session := models.Session{
		UserID: userID,
	}

	if err := uc.UserRepository.DeleteUser(ctx, userID); err != nil {

		return err
	}
	if err := uc.sessionRepo.UpdateToken(ctx, &session); err != nil {
		return err
	}
	return nil
}
