package userusecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) PromoteUser(username string, promoted bool, claims *domain.LoginClaims) error {
	user, err := u.UserRepo.GetUserByUsernameorEmail(username)
	if err != nil {
		return err
	}

	if claims.Role != "root" {
		return config.ErrUserCantPromote
	}

	if promoted {
		if user.Role == "admin" {
			return config.ErrAlreadyAdmin
		}

		user.Role = "admin"
	} else {
		if user.Role == "user" {
			return config.ErrAlreadyUser
		}

		user.Role = "user"
	}

	err = u.UserRepo.UpdateProfile(username, user)
	if err != nil {
		return err
	}

	return nil
}
