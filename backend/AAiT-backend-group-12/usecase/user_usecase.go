package usecase

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	infrastructure "blog_api/infrastructure/cryptography"
	"blog_api/infrastructure/jwt"
	"context"
	"net/mail"
	"strings"
	"time"
)

type UserUsecase struct {
	userRepository domain.UserRepositoryInterface
}

func NewUserUsecase(userRepository domain.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) Signup(c context.Context, user *domain.User) domain.CodedError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Username = strings.TrimSpace(strings.ToLower(user.Username))

	user.CreatedAt = time.Now().Round(0)
	if len(user.Username) < 3 {
		return domain.NewError("Username too short", domain.ERR_BAD_REQUEST)
	}

	if len(user.Username) > 20 {
		return domain.NewError("Username too short", domain.ERR_BAD_REQUEST)
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return domain.NewError("Invalid email", domain.ERR_BAD_REQUEST)
	}

	if len(user.Password) < 8 {
		return domain.NewError("Password too short", domain.ERR_BAD_REQUEST)
	}

	hashedPw, err := infrastructure.HashString(user.Password)
	if err != nil {
		return domain.NewError("Internal server error", domain.ERR_INTERNAL_SERVER)
	}
	user.Password = hashedPw
	user.Role = "user"

	err = u.userRepository.CreateUser(c, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Login(c context.Context, user *domain.User) (string, string, domain.CodedError) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Username = strings.TrimSpace(strings.ToLower(user.Username))

	foundUser, err := u.userRepository.FindUser(c, user)
	if err != nil {
		return "", "", err
	}

	err = infrastructure.ValidateHashedString(foundUser.Password, user.Password)
	if err != nil {
		return "", "", domain.NewError("Incorrect password", domain.ERR_UNAUTHORIZED)
	}

	accessToken, err := jwt.SignJWTWithPayload(user.Username, user.Role, "accessToken", time.Hour*time.Duration(env.ENV.ACCESS_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.SignJWTWithPayload(user.Username, user.Role, "refreshToken", time.Hour*time.Duration(env.ENV.REFRESH_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	// set the new refresh token in the database after hashing it
	hashedRefreshToken, err := infrastructure.HashString(strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	err = u.userRepository.SetRefreshToken(c, user, hashedRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
