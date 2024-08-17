package usecase

import (
	"errors"
	"group3-blogApi/repository"
)

type OAuthUsecase interface {
	GetLoginURL() string
	HandleCallback(state, code string) (string, error)
}

type oauthUsecase struct {
	oauthRepo repository.OAuthRepository
}

func NewOAuthUsecase(oauthRepo repository.OAuthRepository) OAuthUsecase {
	return &oauthUsecase{oauthRepo: oauthRepo}
}

func (u *oauthUsecase) GetLoginURL() string {
	return u.oauthRepo.GenerateAuthURL()
}

func (u *oauthUsecase) HandleCallback(state, code string) (string, error) {
	if state != "random" {
		return "", errors.New("invalid OAuth state")
	}

	token, err := u.oauthRepo.ExchangeCodeForToken(code)
	if err != nil {
		return "", err
	}

	content, err := u.oauthRepo.GetUserInfo(token)
	if err != nil {
		return "", err
	}
	

	return content, nil
}