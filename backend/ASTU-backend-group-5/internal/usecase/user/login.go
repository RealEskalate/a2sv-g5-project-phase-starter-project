package user

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"blogApp/pkg/hash"
	"blogApp/pkg/jwt"
	"context"
	"errors"
)

type LoginUseCase struct {
	userRepository repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{userRepository: userRepo}
}

func (uc *LoginUseCase) Login(email, password string) (*domain.User, *domain.Token, error) {
	user, err := uc.userRepository.FindByEmail(context.TODO(), email)
	if err != nil {
		return nil, nil, err
	}
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return nil, nil, err
	}
	
	if user.Password != hashedPassword {
		return nil, nil, errors.New("invalid credentials")
	}
	accessToken, err := jwt.GenerateJWT(user.ID.Hex(), user.UserName, user.Email, user.Role)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID.Hex(), user.UserName, user.Email, user.Role)
	if err != nil {
		return nil, nil, err
	}

	return user, &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
