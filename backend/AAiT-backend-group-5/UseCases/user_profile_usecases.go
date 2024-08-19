package usecases

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type userProfileUpdateUsecase struct {
	repository      interfaces.UserRepository
	passwordService interfaces.PasswordService
}

func NewUserProfileUpdateUsecase(repository interfaces.UserRepository, passwordService interfaces.PasswordService) interfaces.UserProfileUpdateUsecase {
	return &userProfileUpdateUsecase{
		repository:      repository,
		passwordService: passwordService,
	}
}

func (uc *userProfileUpdateUsecase) UpdateUserProfile(ctx context.Context, userID string, user *dtos.ProfileUpdateRequest) *models.ErrorResponse {
	// Check if the user exists

	existUser, err := uc.repository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Update the user profile
	if user.Username != "" {
		existUser.Username = user.Username
	}
	if user.Name != "" {
		existUser.Name = user.Name
	}

	if user.Password != "" {
		// Hash the password
		hashedPassword, err := uc.passwordService.EncryptPassword(user.Password)
		if err != nil {
			return models.InternalServerError("Something went wrong")

		}
		existUser.Password = hashedPassword
	}

	err = uc.repository.UpdateUser(ctx, existUser, userID)
	if err != nil {
		return err
	}

	return nil
}
