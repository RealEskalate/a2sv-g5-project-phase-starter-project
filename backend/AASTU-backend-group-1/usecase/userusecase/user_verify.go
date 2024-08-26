package userusecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) VerifyUser(token string) error {
	claims := &domain.RegisterClaims{}
	err := config.ValidateToken(token, claims)
	if err != nil {
		return err
	}

	err = u.UserRepo.CheckUsernameAndEmail(claims.User.Username, claims.User.Email)
	if err != nil {
		return config.ErrUserAlreadyVerified
	}

	err = u.UserRepo.RegisterUser(&claims.User)
	return err
}
