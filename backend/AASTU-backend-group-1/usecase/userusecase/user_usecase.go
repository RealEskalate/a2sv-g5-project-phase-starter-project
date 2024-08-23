package userusecase

import (
	"blogs/domain"
)

type UserUsecase struct {
	UserRepo   domain.UserRepository
	Oauth2Repo domain.OAuthStateRepository
}

func NewUserUsecase(ur domain.UserRepository, or domain.OAuthStateRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo:   ur,
		Oauth2Repo: or,
	}
}
