package usecase

import "group3-blogApi/domain"

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

