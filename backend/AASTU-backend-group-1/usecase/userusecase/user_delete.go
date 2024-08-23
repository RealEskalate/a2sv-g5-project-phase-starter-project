package userusecase

func(u *UserUsecase) DeleteUser(username string) error {
	return u.UserRepo.DeleteUser(username)
}