package usecases

import (
	"errors"
	"meleket/domain"
	"meleket/infrastructure"
    "time"
)

type TokenUsecase struct {
    tokenRepo domain.TokenRepositoryInterface
    jwtSvc   infrastructure.JWTService
}

// func NewTokenRepository(tr repository.UserRepositoryInterface, js domain.JWTService) *UserUsecase {
func NewTokenUsecase(tr domain.TokenRepositoryInterface, js infrastructure.JWTService) *TokenUsecase{
    return &TokenUsecase{
        tokenRepo: tr,
        jwtSvc:   js,
    }
}

// RefreshToken refreshes a user's JWT token
func (u *TokenUsecase) RefreshToken(refreshToken *domain.RefreshToken) (string, error) {
    // Check if the refresh token is expired
    if refreshToken.ExpiresAt.Before(time.Now()) {
        return "", errors.New("refresh token expired!! Login again")
    }

    _, err := u.tokenRepo.FindRefreshToken(refreshToken.UserID)
    if err != nil {
        return "", errors.New("invalid refresh token")
    }


    // Generate a new JWT token
    newToken, err := u.jwtSvc.GenerateToken(refreshToken.UserID, refreshToken.Role)
    if err != nil {
        return "", errors.New("could not generate new JWT token!")
    }

    return newToken, nil
}
