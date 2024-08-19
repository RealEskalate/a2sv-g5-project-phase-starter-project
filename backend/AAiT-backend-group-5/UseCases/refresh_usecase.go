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
	ctx               context.Context
}

func NewRefreshUsecase(jwtService interfaces.JwtService, sessionRepository interfaces.SessionRepository, userRepository interfaces.UserRepository, ctx context.Context) interfaces.RefreshUsecase {
	return &refreshUsecase{
		jwtService:        jwtService,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		ctx:               ctx,
	}
}

func (uc *refreshUsecase) RefreshToken(userID string, refreshToken string) (string, *models.ErrorResponse) {
	//get the user session
	session, err := uc.sessionRepository.GetToken(uc.ctx, userID)
	if err != nil {
		return "", err
	}

	//check if the refresh token is valid
	if session.RefreshToken != refreshToken {
		return "", models.Unauthorized("Invalid refresh token")
	}

	user, err := uc.userRepository.GetUserByID(uc.ctx, userID)
	if err != nil {
		return "", err
	}
	//generate a new access token
	accessToken, tErr := uc.jwtService.CreateAccessToken(*user, 60)
	newRefresheToken, rErr := uc.jwtService.CreateRefreshToken(*user, 60)

	if tErr != nil || rErr != nil {
		return "", models.InternalServerError("An unexpected error occurred")
	}

	newSession := models.Session{
		UserID:       userID,
		AccessToken:  accessToken,
		RefreshToken: newRefresheToken,
	}

	// store the new access token
	err = uc.sessionRepository.UpdateToken(uc.ctx, &newSession)
	if err != nil {
		return "", err
	}

	return accessToken, err
}
