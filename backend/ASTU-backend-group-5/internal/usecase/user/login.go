package user

import (
	"blogApp/internal/domain"
	"blogApp/pkg/hash"
	"blogApp/pkg/jwt"
	"context"
	"errors"
)

func (u *UserUsecase) Login(email, password string) (*domain.User, *domain.Token, error) {
	user, err := u.repo.FindUserByEmail(context.TODO(), email)

	if err != nil {
		return nil, nil, err
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, nil, errors.New("invalid credentials")
	}
	// hashedPassword, err := hash.HashPassword(password)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// if user.Password != hashedPassword {
	// 	return nil, nil, errors.New("invalid credentials")
	// }
	accessToken, err := jwt.GenerateJWT(user.ID.Hex(), user.Email, user.UserName, user.Role)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID.Hex(), user.UserName, user.Email, user.Role)
	if err != nil {
		return nil, nil, err
	}

	return user, &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
