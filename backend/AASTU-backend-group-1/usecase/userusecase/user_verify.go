package userusecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) VerifyUser(token string) error {
	claims, err := config.ValidateToken(token, "register")
	if err != nil {
		return err
	}

	registerClaims, ok := claims.(*domain.RegisterClaims)
	if !ok {
		return config.ErrInvalidToken
	}

	user, err := u.UserRepo.GetUserByUsernameorEmail(registerClaims.Username)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return config.ErrAlreadyVerified
	}

	user.IsVerified = true
	return u.UserRepo.UpdateProfile(registerClaims.Username, user)
}
