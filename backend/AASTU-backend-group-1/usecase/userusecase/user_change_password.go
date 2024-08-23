package userusecase

import "blogs/config"

func (u *UserUsecase) ChangePassword(username, oldPassword, newPassword string) error {
	user, err := u.UserRepo.GetUserByUsernameorEmail(username)
	if err != nil {
		return err
	}

	err = config.ComparePassword(user.Password, oldPassword)
	if err != nil {
		return config.ErrIncorrectPassword
	}

	hashedPassword, err := config.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if oldPassword == newPassword {
		return config.ErrSamePassword
	}

	err = u.UserRepo.Resetpassword(user.Username, hashedPassword)
	return err
}
