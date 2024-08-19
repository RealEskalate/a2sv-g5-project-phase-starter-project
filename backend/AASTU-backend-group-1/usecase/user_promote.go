package usecase

import (
	"blogs/domain"
	"errors"
)

func (u *UserUsecase) PromoteUser(username string, promoted bool, claims *domain.LoginClaims) error {
	user, err := u.UserRepo.GetUserByUsernameorEmail(username)
	if err != nil {
		return err
	}

	if claims.Role != "root" {
		return errors.New("only root can promote or demote users")
	}

	if promoted {
		if user.Role == "admin" {
			return errors.New("user is already an admin")
		}

		user.Role = "admin"
	} else {
		if user.Role == "user" {
			return errors.New("user is already a regular user")
		}

		user.Role = "user"
	}

	err = u.UserRepo.UpdateProfile(username, user)
	if err != nil {
		return err
	}

	return nil
}
