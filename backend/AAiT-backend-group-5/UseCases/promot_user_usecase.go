package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type userUsecase struct {
	repo interfaces.UserRepository
}

<<<<<<< HEAD
func NewUserUsecase(repo interfaces.UserRepository) interfaces.PromoteDemoteUserUsecase {
=======
func NewUserUsecase(repo interfaces.UserRepository) interfaces.PromoteUserUsecase {
>>>>>>> origin/aait.backend.g5.bisrat.setup-db-and-user-repo
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) PromoteUser(ctx context.Context, userID string) *models.ErrorResponse {

	user, err := uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Check if the user is already an admin
	if user.Role == "admin" {
		return models.BadRequest("User is already an admin")
	}

	// Promote the user

	err = uc.repo.PromoteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUsecase) DemoteUser(ctx context.Context, userID string) *models.ErrorResponse {

	user, err := uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Check if the user is not an admin
	if user.Role != "admin" {
		return models.BadRequest("User is not an admin")
	}

	// Demote the user

	err = uc.repo.DemoteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
