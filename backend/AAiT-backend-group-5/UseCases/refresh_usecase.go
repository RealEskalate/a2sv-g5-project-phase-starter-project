package usecases

import (
	"context"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type refreshUsecase struct {
	jwtService        interfaces.JwtService
	sessionRepository interfaces.SessionRepository
	userRepository    interfaces.UserRepository
}

func NewRefreshUsecase(jwtService interfaces.JwtService, sessionRepository interfaces.SessionRepository, userRepository interfaces.UserRepository) interfaces.RefreshUsecase {
	return &refreshUsecase{
		jwtService:        jwtService,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

func (uc *refreshUsecase) RefreshToken(ctx context.Context, userID string, refreshToken string) (string, *models.ErrorResponse) {
	session, err := uc.sessionRepository.GetToken(ctx, userID)
	if err != nil {
		return "", err
	}

	if session.RefreshToken != refreshToken {
		return "", models.Unauthorized("Invalid refresh token")
	}

	user, err := uc.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return "", err
	}
	accessToken, tErr := uc.jwtService.CreateAccessToken(*user, 60)

	if tErr != nil {
		return "", models.InternalServerError("An unexpected error occurred")
	}

	newSession := models.Session{
		UserID:       userID,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	err = uc.sessionRepository.UpdateToken(ctx, &newSession)
	if err != nil {
		return "", err
	}

	return accessToken, err
}
