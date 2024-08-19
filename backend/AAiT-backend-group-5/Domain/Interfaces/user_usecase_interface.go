package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type PromoteUserUsecase interface {
	PromoteUser(ctx context.Context, userID string) *models.ErrorResponse
	DemoteUser(ctx context.Context, userID string) *models.ErrorResponse
}

type UserProfileUpdateUsecase interface {
	UpdateUserProfile(ctx context.Context, userID string, user *dtos.ProfileUpdateRequest) *models.ErrorResponse
}
