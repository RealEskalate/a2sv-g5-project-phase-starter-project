package usecase

import (
	"blogs/domain"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		UserRepo: ur,
	}
}
