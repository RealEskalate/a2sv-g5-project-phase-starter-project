package userusecase

import "blogs/domain"

func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	user, err := u.UserRepo.GetUserByUsernameorEmail(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
