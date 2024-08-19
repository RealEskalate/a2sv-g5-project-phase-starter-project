package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type userUsecase struct {
	repo interfaces.UserRepository
}

func NewUserUsecase(repo interfaces.UserRepository) interfaces.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) PromoteUser(ctx context.Context, user *models.User) *models.ErrorResponse {
	user, err := uc.repo.GetUserByID(ctx, user.ID.Hex())
	if err != nil {
		return err
	}

	// Check if the user is already an admin
	if user.Role == "admin" {
		return models.BadRequest("User is already an admin")
	}

	// Promote the user
	user.Role = "admin"
	err = uc.repo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUsecase) DemoteUser(ctx context.Context, user *models.User) *models.ErrorResponse {

	user, err := uc.repo.GetUserByID(ctx, user.ID.Hex())
	if err != nil {
		return err
	}

	// Check if the user is not an admin
	if user.Role != "admin" {
		return models.BadRequest("User is not an admin")
	}

	// Demote the user
	user.Role = "user"
	err = uc.repo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
