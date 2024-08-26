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

	accessClaims := &domain.LoginClaims{
		Username: claims.Username,
		Role:     claims.Role,
		Type:     "access",
	}

	accessToken, err := config.GenerateToken(accessClaims)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
