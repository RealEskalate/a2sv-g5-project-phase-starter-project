package usecase

import (
	"blogs/config"
	"blogs/domain"
	"errors"
)

func (u *UserUsecase) VerifyUser(token string) error {
	claims, err := config.ValidateToken(token, "register")
	if err != nil {
		return err
	}

	registerClaims, ok := claims.(*domain.RegisterClaims)
	if !ok {
		return errors.New("invalid token claims type")
	}

	user, err := u.UserRepo.GetUserByUsernameorEmail(registerClaims.Username)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return errors.New("user is already verified")
	}

	user.IsVerified = true
	return u.UserRepo.UpdateProfile(registerClaims.Username, user)
}
