package user

import (
	"blogApp/internal/domain"
	"blogApp/pkg/hash"
	"blogApp/pkg/jwt"
	"context"
	"errors"
)

func (u *UserUsecase) Login(email string, password string) (*domain.User, *domain.Token, error) {
	user, err := u.repo.FindUserByEmail(context.TODO(), email)

	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, errors.New("invalid credentials")
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
	accessToken, err := jwt.GenerateJWT(user.ID.Hex(), user.UserName, user.Email,  user.Role)
	
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID.Hex(), user.UserName, user.Email, user.Role)
	if err != nil {
		return nil, nil, err
	}

	return user, &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

// GoogleCallback is a usecase that handles the google oauth callback
func (u *UserUsecase) GoogleCallback(code string) (*domain.User, *domain.Token, error) {
	googleUser, err := u.repo.GoogleCallback(context.Background(), code)
	if err != nil {
		return nil, nil, err
	}
	if googleUser == nil {
		return nil, nil, errors.New("invalid credentials")
	}

	accessToken, err := jwt.GenerateJWT(googleUser.ID.Hex(), googleUser.Email, googleUser.UserName, googleUser.Role)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(googleUser.ID.Hex(), googleUser.UserName, googleUser.Email, googleUser.Role)
	if err != nil {
		return nil, nil, err
	}

	return googleUser, &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}