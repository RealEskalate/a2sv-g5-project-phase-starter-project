package user_usecase

import (
	"blog-api/domain"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) SignUp(ctx context.Context, req domain.SignupRequest) (domain.SignupResponse, error) {
	existingUser, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.SignupResponse{}, err
	}
	if existingUser != nil {
		return domain.SignupResponse{}, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	user := &domain.User{
		Firstname:          req.Firstname,
		Lastname:           req.Lastname,
		Username:           req.Username,
		Password:           string(hashedPassword),
		Email:              req.Email,
		Bio:                "",
		ProfilePicture:     "",
		ContactInformation: "",
		IsAdmin:            false,
		Active:             true,
		CreatedAt:          time.Now(),
	}
	if err := u.userRepo.CreateUser(ctx, user); err != nil {
		return domain.SignupResponse{}, err
	}

	accessToken, err := u.authService.GenerateAccessToken(ctx, *user)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	refreshToken, err := u.authService.GenerateAndStoreRefreshToken(ctx, user.ID.Hex())
	if err != nil {
		return domain.SignupResponse{}, err
	}

	return domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
