package usecase

import "blogs/config"

func (u *UserUsecase) ResetPassword(usernameoremail string, newPassword string) error {
	// Get user by username or email
	_, err := u.UserRepo.GetUserByUsernameorEmail(usernameoremail)
	if err != nil {
		return err
	}

	// Hash the new password
	hashedPassword, err := config.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update the password in the repository
	err = u.UserRepo.Resetpassword(usernameoremail, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
