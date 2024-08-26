package usecase

import (
	"group3-blogApi/domain"
)

func (uc *UserUsecase) GetMyProfile(userID string) (domain.User, *domain.CustomError) {
	user, err := uc.UserRepo.GetMyProfile(userID)
	if err != nil {
		return domain.User{}, domain.ErrNotFound
	}
	return user, &domain.CustomError{}
}

func (uc *UserUsecase) GetUsers() ([]domain.User, *domain.CustomError) {
	users, err := uc.UserRepo.GetUsers()
	if err != nil {
		return nil, domain.ErrNotFound
	}
	
	return users, &domain.CustomError{} 
}

func (uc *UserUsecase) DeleteUser(userID string) (domain.User, *domain.CustomError) {
	user, err := uc.UserRepo.DeleteUser(userID)
	if err != nil {
		return domain.User{}, domain.ErrFailedToDeleteUser
	}
	return user, &domain.CustomError{}
}

func (uc *UserUsecase) UpdateUserRole(userID, role string) (domain.User, *domain.CustomError) {
	if role == "" || userID == "" {
		return domain.User{}, domain.ErrMissingRequiredFields
	}
	if role != "admin" && role != "user" {
		return domain.User{}, domain.ErrInvalidUpdateRequest
	}

	user, err := uc.UserRepo.UpdateUserRole(userID, role)
	if err != nil {
		return domain.User{}, domain.ErrFailedToUpdateUser
	}
	return user, &domain.CustomError{}
}

func (uc *UserUsecase) DeleteMyAccount(userID string) *domain.CustomError {
	err := uc.UserRepo.DeleteMyAccount(userID)
	if err != nil {
		return domain.ErrFailedToDeleteAccount
	}
	return &domain.CustomError{}
}

func (uc *UserUsecase) UploadImage(userID string, imagePath string) *domain.CustomError {
	err := uc.UserRepo.UploadImage(userID, imagePath)
	if err != nil {
		return domain.ErrFailedToUploadImage
	}
	return &domain.CustomError{}
}

func (uc *UserUsecase) UpdateMyProfile(user domain.User, UserID string) *domain.CustomError {
	if user.Bio == "" || user.Username == "" {
		return domain.ErrMissingRequiredFields
	}

	err := uc.UserRepo.UpdateMyProfile(user, UserID)
	if err != nil {
		return domain.ErrFailedToUpdateProfile
	}
	return &domain.CustomError{}
}
