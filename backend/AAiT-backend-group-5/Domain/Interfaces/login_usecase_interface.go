package interfaces

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type LoginUsecase interface {
	LoginUser(ctx context.Context, userReqest dtos.LoginRequest) (*dtos.LoginResponse, *models.ErrorResponse)
	GenerateAccessToken(user *models.User, expiry int) (string, *models.ErrorResponse)
	GenerateRefreshToken(user *models.User, expiry int) (string, *models.ErrorResponse)
}

