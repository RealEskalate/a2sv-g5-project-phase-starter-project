package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type LogoutUsecase interface {
	LogoutUser(ctx context.Context, userID string) *models.ErrorResponse
}
