package usecases

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"context"
)

type tokenUsecase struct {
	tokenService domain.TokenInfrastructure
	env          *bootstrap.Env
}

func NewTokenUsecase(tokenService domain.TokenInfrastructure, env *bootstrap.Env) domain.TokenUseacses {
	return &tokenUsecase{
		tokenService: tokenService,
		env:          env,
	}
}

func (tu *tokenUsecase) UpdateTokens(c context.Context, id string) (accessToken string, refreshToken string, err error) {
	accessToken, refreshToken, err = tu.tokenService.UpdateTokens(id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
