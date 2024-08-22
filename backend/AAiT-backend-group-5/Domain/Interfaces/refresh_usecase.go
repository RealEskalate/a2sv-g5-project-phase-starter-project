package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type RefreshUsecase interface {
	RefreshToken(c context.Context, userID string, refreshToken string) (string, *models.ErrorResponse)
}
