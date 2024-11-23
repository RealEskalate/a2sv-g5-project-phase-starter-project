package user_usecase

import (
	"blog-api/domain"
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error) {
	// Find user by username or email
	user, err := u.userRepo.GetByUsernameOrEmail(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Generate the access token
	accessToken, err := u.authService.GenerateAccessToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	// Generate and store the refresh token
	_, err = u.authService.GenerateAndStoreRefreshToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken: accessToken,
	}, nil
}
