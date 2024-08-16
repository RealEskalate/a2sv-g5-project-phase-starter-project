func (u *userUsecase) Login(*Domain.LoginInput) (string, error) {
	user, err := u.userRepository.FindByEmail(email)
    if err!= nil {
        return " ", err
    }

	storedPassword := user.Password

	err = infrastructure.ComparePasswords(storedPassword, password)

	if err != nil{
		return  " ", err
	}

	access_token, err  := infrastructure.GenerateToken(user.Username , user.Role)

	if err != nil{
		return " ", err
	}
	refresh_token , err := infrastructure.GenerateRefreshToken(user.Username)

	if err != nil{
		return " ", err
	}

	err =  u.userRepository.InsertToken(user.Username , access_token , refresh_token)
	if err != nil{
		return " ", err
	}
    
	return access_token, nil
}

func (u *userUsecase) Logout( username  string) error {
	u.userRepository.DeleteToken (username )
}

