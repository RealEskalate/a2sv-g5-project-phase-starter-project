package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type SignupUsecase interface {
	CreateUser(ctx context.Context, user *models.User) models.ErrorResponse
	GetUserByID(ctx context.Context, id string) (*models.User, models.ErrorResponse)
}
