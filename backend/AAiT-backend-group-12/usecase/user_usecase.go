package usecase

import (
	"blog_api/delivery/env"
	"blog_api/domain"
	"blog_api/domain/dtos"
	"blog_api/infrastructure/cryptography"
	jwt_service "blog_api/infrastructure/jwt"
	"context"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

type UserUsecase struct {
	userRepository domain.UserRepositoryInterface
}

var PhoneRegex = regexp.MustCompile(`^\+?[1-9][0-9]{7,14}$`)

func NewUserUsecase(userRepository domain.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) ValidatePassword(password string) domain.CodedError {
	if len(password) < 8 {
		return domain.NewError("Password too short", domain.ERR_BAD_REQUEST)
	}

	if len(password) > 71 {
		return domain.NewError("Password too long", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func (u *UserUsecase) ValidateUsername(username string) domain.CodedError {
	if len(username) < 3 {
		return domain.NewError("Username too short", domain.ERR_BAD_REQUEST)
	}

	if len(username) > 20 {
		return domain.NewError("Username too short", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func (u *UserUsecase) ValidateEmail(email string) domain.CodedError {
	if _, err := mail.ParseAddress(email); err != nil {
		return domain.NewError("Invalid email", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func (u *UserUsecase) SantizeUserFields(user *domain.User) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Username = strings.TrimSpace(strings.ToLower(user.Username))
	user.Bio = strings.TrimSpace(user.Bio)
	user.PhoneNumber = strings.TrimSpace(user.PhoneNumber)
}

func (u *UserUsecase) SanitizeAndValidateNewUser(user *domain.User) domain.CodedError {
	u.SantizeUserFields(user)
	user.CreatedAt = time.Now().Round(0)

	err := u.ValidateUsername(user.Username)
	if err != nil {
		return err
	}

	err = u.ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	err = u.ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	if len(user.Bio) == 0 {
		return domain.NewError("Bio can not be empty", domain.ERR_BAD_REQUEST)
	}

	if !PhoneRegex.MatchString(user.PhoneNumber) {
		return domain.NewError("Invalid phone number: must be informat +XXXXXXXXXX", domain.ERR_BAD_REQUEST)
	}

	return nil
}

func (u *UserUsecase) Signup(c context.Context, user *domain.User) domain.CodedError {
	err := u.SanitizeAndValidateNewUser(user)
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
	u.SantizeUserFields(user)
	err := u.ValidateUsername(user.Username)
	if user.Username != "" && err != nil {
		return "", "", err
	}

	err = u.ValidateEmail(user.Email)
	if user.Email != "" && err != nil {
		return "", "", err
	}

	// if both username and email are empty return an error
	if user.Email == "" && user.Username == "" {
		return "", "", domain.NewError("Username or email required", domain.ERR_BAD_REQUEST)
	}

	err = u.ValidatePassword(user.Password)
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

	accessToken, err := jwt_service.SignJWTWithPayload(foundUser.Username, foundUser.Role, "accessToken", time.Hour*time.Duration(env.ENV.ACCESS_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt_service.SignJWTWithPayload(foundUser.Username, foundUser.Role, "refreshToken", time.Hour*time.Duration(env.ENV.REFRESH_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
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

func (u *UserUsecase) RenewAccessToken(c context.Context, refreshToken string) (string, domain.CodedError) {
	token, err := jwt_service.ValidateAndParseToken(refreshToken, env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", domain.NewError("Invalid token", domain.ERR_UNAUTHORIZED)
	}

	// get the username from the token
	username, err := jwt_service.GetUsername(token)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	// check expiry date of the refresh token
	expiresAtTime, err := jwt_service.GetExpiryDate(token)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	if expiresAtTime.Compare(time.Now()) == -1 {
		u.userRepository.SetRefreshToken(c, &domain.User{Username: username}, "")
		return "", domain.NewError("Token expired", domain.ERR_UNAUTHORIZED)
	}

	// get the hashed refresh token from the database
	foundUser, qErr := u.userRepository.FindUser(c, &domain.User{Username: username})
	if qErr != nil {
		return "", qErr
	}

	if foundUser.RefreshToken == "" {
		return "", domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	err = cryptography.ValidateHashedString(foundUser.RefreshToken, strings.Split(refreshToken, ".")[2])
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_UNAUTHORIZED)
	}

	accessToken, err := jwt_service.SignJWTWithPayload(foundUser.Username, foundUser.Role, "accessToken", time.Hour*time.Duration(env.ENV.ACCESS_TOKEN_LIFESPAN), env.ENV.JWT_SECRET_TOKEN)
	if err != nil {
		return "", domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return accessToken, nil
}

func (u *UserUsecase) UpdateUser(c context.Context, requestUsername string, tokenUsername string, user *dtos.UpdateUser) (map[string]string, domain.CodedError) {
	if requestUsername != tokenUsername {
		return nil, domain.NewError("Only the owner of the account can update its details", domain.ERR_FORBIDDEN)
	}

	if user.PhoneNumber != "" && !PhoneRegex.MatchString(user.PhoneNumber) {
		return nil, domain.NewError("Invalid phone number: must be informat +XXXXXXXXXX", domain.ERR_BAD_REQUEST)
	}

	if user.Bio == "" && user.PhoneNumber == "" {
		return nil, domain.NewError("No fields to update", domain.ERR_BAD_REQUEST)
	}

	return u.userRepository.UpdateUser(c, requestUsername, user)
}
