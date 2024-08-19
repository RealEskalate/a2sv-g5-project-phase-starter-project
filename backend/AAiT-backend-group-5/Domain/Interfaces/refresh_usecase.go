package interfaces

import models "github.com/aait.backend.g5.main/backend/Domain/Models"

type RefreshUsecase interface {
	RefreshToken(userID string, refreshToken string) (string, *models.ErrorResponse)
}
