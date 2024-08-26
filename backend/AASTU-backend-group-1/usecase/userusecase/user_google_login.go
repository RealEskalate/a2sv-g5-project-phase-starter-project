package userusecase

import (
	"blogs/config"
	"blogs/domain"
	"context"
	"time"
)

func (u *UserUsecase) GoogleLogin() (string, error) {
	stateString := config.GenerateState()

	state := &domain.OAuthState{
		ID:        stateString,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 10),
	}

	err := u.Oauth2Repo.InsertState(state)
	if err != nil {
		return "", err
	}

	return config.GetGoogleLoginURL(stateString), nil
}

func (u *UserUsecase) GoogleCallback(stateString, code string) (string, string, error) {
	state, err := u.Oauth2Repo.GetState(stateString)
	if err != nil {
		return "", "", err
	}

	if time.Now().After(state.ExpiresAt) {
		return "", "", config.ErrStateExpired
	}

	user, err := config.HandleGoogleCallback(context.Background(), code)
	if err != nil {
		return "", "", err
	}

	_, err = u.UserRepo.GetUserByUsernameorEmail(user.Email)
	if err != nil && err != config.ErrUserNotFound {
		err = u.UserRepo.RegisterUser(user)
		if err != nil {
			return "", "", err
		}
	}

	// Generate access token
	accessToken, err := config.GenerateToken(
		&domain.LoginClaims{
			Username: user.Username,
			Role:     user.Role,
			Type:     "access",
		})

	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshClaims := &domain.LoginClaims{
		Username: user.Username,
		Role:     user.Role,
		Type:     "refresh",
	}

	refreshToken, err := config.GenerateToken(refreshClaims)
	if err != nil {
		return "", "", err
	}

	// Save the refresh token in the repository
	err = u.UserRepo.InsertToken(refreshClaims.ToToken())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
