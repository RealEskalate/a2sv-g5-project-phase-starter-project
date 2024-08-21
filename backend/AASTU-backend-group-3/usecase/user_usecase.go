package usecase

import "group3-blogApi/domain"

type UserUsecase struct {
	UserRepo       domain.UserRepository
	TokenGen       domain.TokenGenerator
	PasswordSvc    domain.PasswordService
}

func NewUserUsecase(userRepo domain.UserRepository, tokenGen domain.TokenGenerator, passwordSvc domain.PasswordService) domain.UserUsecase {
	return &UserUsecase{
		UserRepo:    userRepo,
		TokenGen:    tokenGen,
		PasswordSvc: passwordSvc,
	}
}