package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type UserUsecase interface {
	PromoteUser(ctx context.Context, user *models.User) *models.ErrorResponse
	DemoteUser(ctx context.Context, user *models.User) *models.ErrorResponse
}
