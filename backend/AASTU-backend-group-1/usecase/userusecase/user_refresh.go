package userusecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) RefreshToken(claims *domain.LoginClaims) (string, error) {
	_, err := u.UserRepo.GetTokenByUsername(claims.Username)
	if err != nil {
		if err == config.ErrTokenNotFound {
			return "", config.ErrTokenBlacklisted
		}
		return "", err
	}

	accessToken, _, err := config.GenerateToken(claims, "access")
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
