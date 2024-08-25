package user_usecase

import (
	"blog-api/domain"
	"blog-api/infrastructure/auth"
	"blog-api/infrastructure/bootstrap"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUsecase) LoginUser(c context.Context, loginRequest domain.LoginRequest, Env *bootstrap.Env) (domain.LoginResponse, error) {
	var u domain.User
	var err error

	if loginRequest.Email != "" {
		u, err = uc.repo.GetByEmail(c, loginRequest.Email)
	} else if loginRequest.Username != "" {
		u, err = uc.repo.GetByUsername(c, loginRequest.Username)
	}

	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid credentials. User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginRequest.Password)); err != nil {
		return domain.LoginResponse{}, errors.New("invalid credentials. Incorrect password")
	}

	accessToken, err := auth.CreateAccessToken(&u, Env.AccessTokenSecret, Env.AccessTokenExpiryHour)
	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to generate access token")
	}

	refreshToken, err := auth.CreateRefreshToken(&u, Env.RefreshTokenSecret, Env.RefreshTokenExpiryHour)
	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to generate refresh token")
	}

	u.RefreshToken = refreshToken

	updateRequest := domain.UpdateRequest{
		Firstname:          u.Firstname,
		Lastname:           u.Lastname,
		Username:           u.Username,
		Bio:                u.Bio,
		ProfilePicture:     u.ProfilePicture,
		ContactInformation: u.ContactInformation,
		RefreshToken:       refreshToken,
	}
	err = uc.repo.UpdateUser(c, u.ID, &updateRequest)

	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to update user with refresh token")

	}

	return domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
