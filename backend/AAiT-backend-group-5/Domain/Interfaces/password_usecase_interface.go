package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type PasswordUsecase interface {
	GenerateResetURL(ctx context.Context, email string) (string, *models.ErrorResponse)
	SendResetEmail(ctx context.Context, email string, resetURL string) *models.ErrorResponse
	UpdateUserPassword(ctx context.Context, newPassword string, userID string) *models.ErrorResponse
}
