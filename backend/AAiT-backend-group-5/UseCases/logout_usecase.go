package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type logoutUsecase struct {
	jwtService interfaces.JwtService
	repository interfaces.SessionRepository
}

func NewLogoutUsecase(jwtService interfaces.JwtService, repository interfaces.SessionRepository) interfaces.LogoutUsecase {
	return &logoutUsecase{
		jwtService: jwtService,
		repository: repository,
	}
}

func (uc *logoutUsecase) LogoutUser(ctx context.Context, userID string) *models.ErrorResponse {

	if err := uc.repository.RemoveToken(ctx, userID); err != nil {
		return models.InternalServerError("Error removing user token")
	}

	return nil
}

