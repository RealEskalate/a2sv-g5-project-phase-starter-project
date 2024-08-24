package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type SignupUsecase interface {
	CreateUser(ctx context.Context, user *models.User, agent string) *models.ErrorResponse
}
