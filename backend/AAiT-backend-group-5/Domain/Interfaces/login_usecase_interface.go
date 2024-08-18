package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type LoginUsecase interface {
	GetUserByID(ctx context.Context, id string) (*models.User, models.ErrorResponse)
	GenerateAccessToken(user *models.User, secret string, expiry int) (string, models.ErrorResponse)
	GenerateRefreshToken(user *models.User, secret string, expiry int) (string, models.ErrorResponse)
}

type LogoutUsecase interface {
	LogoutUser(ctx context.Context, userID string) models.ErrorResponse
}
