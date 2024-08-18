package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type LoginUsecase interface {
	LoginUser(ctx context.Context, emailOrUsername string, password string) (*models.User, *models.ErrorResponse)
	GenerateAccessToken(user *models.User, expiry int) (string, *models.ErrorResponse)
	GenerateRefreshToken(user *models.User, expiry int) (string, *models.ErrorResponse)
}

