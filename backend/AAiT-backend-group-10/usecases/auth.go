package usecases

import (
<<<<<<< HEAD
	"errors"
=======
	"fmt"
>>>>>>> e77e9cdd (aait.backend.g10.Yordanos: add image upload)
	"math/rand"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type IAuthUsecase interface {
	RegisterUser(User *dto.RegisterUserDTO) (interface{}, *domain.CustomError)
	LoginUser(dto *dto.LoginUserDTO) (*dto.TokenResponseDTO, *domain.CustomError)
	RefreshTokens(refreshToken string) (*dto.TokenResponseDTO, *domain.CustomError)
	ResetPassword(dto *dto.ResetPasswordRequestDTO) *domain.CustomError
	VerifyUserAccessToken(token string, userId uuid.UUID) (bool, *domain.CustomError)
	ForgotPassword(dto *dto.ForgotPasswordRequestDTO) *domain.CustomError
	HandleGoogleCallback(userDto *domain.User) (string, string, error)
}

type AuthUsecase struct {
	userRepository interfaces.IUserRepository
	jwtService     interfaces.IJwtService
	pwdService     interfaces.IHashingService
	emailService   interfaces.IEmailService
}

func NewAuthUsecase(ur interfaces.IUserRepository, jwt interfaces.IJwtService, pwdService interfaces.IHashingService, emailService interfaces.IEmailService) IAuthUsecase {
	return &AuthUsecase{
		userRepository: ur,
		jwtService:     jwt,
		pwdService:     pwdService,
		emailService:   emailService,
	}
}

func (u *AuthUsecase) RegisterUser(User *dto.RegisterUserDTO) (interface{}, *domain.CustomError) {
	existingUser, _ := u.userRepository.GetUserByEmail(User.Email)

	if existingUser != nil && !existingUser.GoogleSignIn {
		return nil, domain.ErrUserEmailExists
	}

	hashedPassword, err := u.pwdService.HashPassword(User.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		FullName:  User.FullName,
		Email:     User.Email,
		Bio:       User.Bio,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  hashedPassword,
		IsAdmin:   false,
	}
	if existingUser != nil && existingUser.GoogleSignIn {
		existingUser.Password = hashedPassword
		existingUser.FullName = User.FullName
		existingUser.Email = User.Email
		existingUser.Bio = User.Bio
		existingUser.UpdatedAt = time.Now()
		existingUser.GoogleSignIn = false
		err = u.userRepository.UpdateUserToken(existingUser)
		if err != nil {
			return nil, &domain.CustomError{
				Message:    err.Error(),
				StatusCode: 400,
			}
		}
		return &dto.CreatedResponseDto{
			ID:       existingUser.ID,
			FullName: existingUser.FullName,
			Email:    existingUser.Email,
			Bio:      existingUser.Bio,
			ImageUrl: existingUser.ImageURL,
		}, nil
	}

	count, err := u.userRepository.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		user.IsAdmin = true
	}

	err = u.userRepository.CreateUser(user)
	if err != nil {
		return nil, domain.ErrUserCreationFailed
	}

	return &dto.CreatedResponseDto{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Bio:      user.Bio,
		ImageUrl: user.ImageURL,
	}, nil
}

func (uc *AuthUsecase) LoginUser(loginUser *dto.LoginUserDTO) (*dto.TokenResponseDTO, *domain.CustomError) {
	user, err := uc.userRepository.GetUserByEmail(loginUser.Email)
	// fmt.Println(user)
	if err != nil || user == nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Check password
	if user.GoogleSignIn {
		return nil, &domain.CustomError{
			Message:    "please create account",
			StatusCode: 400,
		}
	}
	errs := uc.pwdService.CheckPasswordHash(loginUser.Password, user.Password)
	if !errs {
		return nil, domain.ErrInvalidCredentials
	}

	// Generate tokens
	accessToken, refreshToken, err := uc.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Store refresh token in the database
	user.RefreshToken = refreshToken
	user.AccessToken = accessToken
	err = uc.userRepository.UpdateUserToken(user)
	if err != nil {
		return nil, err
	}

	tokenResponse := &dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenResponse, nil
}

func (uc *AuthUsecase) RefreshTokens(refreshToken string) (*dto.TokenResponseDTO, *domain.CustomError) {
	// Validate the refresh token
	token, err := uc.jwtService.ValidateToken(refreshToken)
	if err != nil || !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, domain.ErrInvalidRefreshToken
	}

	userId := claims["id"]
	userId, errs := uuid.Parse(userId.(string))
	if errs != nil {
		return nil, domain.ErrInvalidToken
	}
	user, err := uc.userRepository.GetUserByID(userId.(uuid.UUID))
	if err != nil || user == nil {
		return nil, domain.ErrUserNotFound
	}

	// Check if the provided refresh token matches the stored one
	token, err = uc.jwtService.ValidateToken(user.RefreshToken)
	if err != nil || !token.Valid {
		return nil, domain.ErrInvalidToken
	}
	if user.RefreshToken != refreshToken {
		fmt.Println("refresh token not equal")
		return nil, domain.ErrInvalidRefreshToken
	}

	// Generate new tokens
	accessToken, refreshToken, err := uc.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Update the user's refresh token
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	err = uc.userRepository.UpdateUserToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *AuthUsecase) ForgotPassword(dto *dto.ForgotPasswordRequestDTO) *domain.CustomError {
	user, err := uc.userRepository.GetUserByEmail(dto.Email)
	if err != nil || user == nil {
		return domain.ErrUserNotFound
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	min := 10000
	max := 100000

	randomNumber := r.Int63n(int64(max-min)) + int64(min)
	resetToken, err := uc.jwtService.GenerateResetToken(user.Email, randomNumber)

	if err != nil {
		return err
	}

	user.ResetCode = randomNumber
	user.ResetToken = resetToken
	err = uc.userRepository.UpdateUserToken(user)
	if err != nil {
		return err
	}
	return uc.emailService.SendResetEmail(user.Email, resetToken)
}

func (uc *AuthUsecase) ResetPassword(dto *dto.ResetPasswordRequestDTO) *domain.CustomError {
	token, err := uc.jwtService.ValidateToken(dto.Token)
	if err != nil || !token.Valid {
		return domain.ErrUserNotFound
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return domain.ErrInvalidToken
	}
	code, ok := claims["code"].(float64)
	if !ok {
		return domain.ErrInvalidToken
	}

	email := claims["email"].(string)
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil || user == nil {
		return domain.ErrUserNotFound
	}
	if user.ResetToken != dto.Token {
		return domain.ErrInvalidToken
	}
	if user.ResetCode != int64(code) {
		return domain.ErrInvalidResetCode
	}

	hashedPassword, err := uc.pwdService.HashPassword(dto.NewPassword)
	if err != nil {
		return err
	}
	user.ResetCode = 0
	user.ResetToken = ""
	user.Password = string(hashedPassword)
	return uc.userRepository.UpdateUserToken(user)
}

<<<<<<< HEAD
func (uc *AuthUsecase) HandleGoogleCallback(userDto *domain.User) (string, string, error) {
	// Get the authorization code from the query parameters

	// Handle the user info (e.g., create a new user, login the user, etc.)
	user, errs := uc.userRepository.GetUserByEmail(userDto.Email)
	if errs != nil && errs.Message != "User not found" {
		return "", "", errors.New(errs.Message)
	}
	if user != nil {
		if !user.GoogleSignIn {
			return "", "", errors.New("login required")
		}
		accessToken, refreshToken, err := uc.jwtService.GenerateToken(user)
		if err != nil {
			return "", "", errors.New(err.Message)
		}
		user.AccessToken = accessToken
		user.RefreshToken = refreshToken
		err = uc.userRepository.UpdateUserToken(user)
		if err != nil {

			return "", "", errors.New(err.Message)
		}
		return accessToken, refreshToken, nil
	}

	accessToken, refreshToken, err := uc.jwtService.GenerateToken(userDto)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}
	userDto.AccessToken = accessToken
	userDto.RefreshToken = refreshToken
	userDto.ID = uuid.New()
	userDto.CreatedAt = time.Now()
	userDto.UpdatedAt = time.Now()
	err = uc.userRepository.CreateUser(userDto)
	if err != nil {
		return "", "", errors.New("failed to register user")
	}
	return accessToken, refreshToken, nil
	// For example, redirect to the home page with the user info
	// c.Redirect(http.StatusTemporaryRedirect, "/")
=======
func (uc *AuthUsecase) VerifyUserAccessToken(token string, userId uuid.UUID) (bool, *domain.CustomError) {
	user, err := uc.userRepository.GetUserByID(userId)
	if err != nil || user == nil {
		return false, domain.ErrUserNotFound
	}
	return (token == user.AccessToken), nil
>>>>>>> e77e9cdd (aait.backend.g10.Yordanos: add image upload)
}
