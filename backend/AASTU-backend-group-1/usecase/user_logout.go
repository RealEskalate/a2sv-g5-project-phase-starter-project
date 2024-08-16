package usecase

func (u *UserUsecase) LogoutUser(username string) error {
	// Check if the user exists
	user, err := u.UserRepo.GetUserByUsernameorEmail(username)
	if err != nil {
		return err
	}

	// Get the token for the user
	token, err := u.UserRepo.GetTokenByUsername(user.Username)
	if err != nil {
		return err
	}

	// Delete the token
	err = u.UserRepo.DeleteToken(token.Username)
	return err
}
