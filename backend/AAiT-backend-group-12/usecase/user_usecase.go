package usecase

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/infrastructure/cryptography"
	jwt_service "blog_api/infrastructure/jwt"
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

func (u *UserUsecase) SanitizeAndValidateUserFields(user *domain.User) domain.CodedError {
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

	if len(user.Password) > 71 {
		return domain.NewError("Password too long", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func (u *UserUsecase) Signup(c context.Context, user *domain.User) domain.CodedError {
	err := u.SanitizeAndValidateUserFields(user)
	if err != nil {
		return err
	}

	hashedPw, err := cryptography.HashString(user.Password)
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
	err := u.SanitizeAndValidateUserFields(user)
	if err != nil {
		return "", "", err
	}

	foundUser, err := u.userRepository.FindUser(c, user)
	if err != nil {
		return "", "", err
	}

	err = cryptography.ValidateHashedString(foundUser.Password, user.Password)
	if err != nil {
		return "", "", domain.NewError("Incorrect password", domain.ERR_UNAUTHORIZED)
	}

	accessToken, err := jwt_service.SignJWTWithPayload(user.Username, user.Role, "accessToken", time.Hour*time.Duration(env.ENV.ACCESS_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt_service.SignJWTWithPayload(user.Username, user.Role, "refreshToken", time.Hour*time.Duration(env.ENV.REFRESH_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	// set the new refresh token in the database after hashing it
	hashedRefreshToken, err := cryptography.HashString(strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	err = u.userRepository.SetRefreshToken(c, user, hashedRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *UserUsecase) RenewAccessToken(c context.Context, user *domain.User, refreshToken string) (string, domain.CodedError) {
	token, err := jwt_service.ValidateAndParseToken(refreshToken, env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", domain.NewError("Invalid token", domain.ERR_UNAUTHORIZED)
	}

	// check expiry date of the refresh token
	expiresAtTime, err := jwt_service.GetExpiryDate(token)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	if expiresAtTime.Compare(time.Now()) == -1 {
		return "", domain.NewError("Token expired", domain.ERR_UNAUTHORIZED)
	}

	// get the hashed refresh token from the database
	foundUser, err := u.userRepository.FindUser(c, user)
	err = cryptography.ValidateHashedString(foundUser.RefreshToken, strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	accessToken, err := jwt_service.SignJWTWithPayload(user.Username, user.Role, "accessToken", time.Hour*time.Duration(env.ENV.ACCESS_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return accessToken, nil
}
