package usecase

import (
	"errors"
	"group3-blogApi/domain"
)

func (uc *UserUsecase) GetMyProfile(userID string) (domain.User, error) {
	user, err := uc.UserRepo.GetMyProfile(userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}


func (uc *UserUsecase) GetUsers() ([]domain.User, error) {
	users, err := uc.UserRepo.GetUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (uc *UserUsecase) DeleteUser(userID string) (domain.User, error) {
	user, err := uc.UserRepo.DeleteUser(userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (uc *UserUsecase) UpdateUserRole(userID, role string) (domain.User, error) {
	user, err := uc.UserRepo.UpdateUserRole(userID, role)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (uc *UserUsecase)DeleteMyAccount(userID string) error{
	err := uc.UserRepo.DeleteMyAccount(userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase)UploadImage (userID string, imagePath string) error{
	err := uc.UserRepo.UploadImage(userID, imagePath)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase)UpdateMyProfile(user domain.User, UserID string) error{
	if user.Bio == "" || user.Username == "" {
		return errors.New("Bio and Username are required")
	}
	
	err := uc.UserRepo.UpdateMyProfile(user, UserID)
	if err != nil {
		return err
	}
	return nil
}