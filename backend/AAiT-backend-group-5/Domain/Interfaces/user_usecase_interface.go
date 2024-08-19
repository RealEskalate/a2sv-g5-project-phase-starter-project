package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type PromoteUserUsecase interface {
	PromoteUser(ctx context.Context, userID string) *models.ErrorResponse
	DemoteUser(ctx context.Context, userID string) *models.ErrorResponse
}
