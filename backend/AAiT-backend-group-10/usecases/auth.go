package usecases

import (
	"errors"
	"fmt"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/usecases/interfaces"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ForgotPasswordRequestDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestDTO struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
type IAuthUsecase interface {
	RegisterUser(User *domain.RegisterUserDTO) (interface{}, error)
	LoginUser(dto *domain.LoginUserDTO) (*domain.TokenResponseDTO, error)
	RefreshTokens(refreshToken string) (*domain.TokenResponseDTO, error)
	ResetPassword(dto *ResetPasswordRequestDTO) error
	ForgotPassword(dto *ForgotPasswordRequestDTO) error
}

type AuthUsecase struct {
	userRepository  interfaces.UserRepositoryInterface
	Infranstructure infrastructures.InfrastructureInterface
}

func NewAuthUsecase(ur interfaces.UserRepositoryInterface, Infr infrastructures.InfrastructureInterface) IAuthUsecase {
	return &AuthUsecase{
		userRepository:  ur,
		Infranstructure: Infr,
	}
}

func (u *AuthUsecase) RegisterUser(User *domain.RegisterUserDTO) (interface{}, error) {
	existingUser, _ := u.userRepository.GetUserByEmail(User.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := u.Infranstructure.HashPassword(User.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		FullName:  User.FullName,
		Email:     User.Email,
		Bio:       User.Bio,
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  hashedPassword,
		IsAdmin:   false,
	}

	err = u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	type created struct {
		ID       string `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		ImageUrl string `json:"image_url"`
	}
	return &created{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Bio:      user.Bio,
		ImageUrl: user.ImageURL,
	}, nil
}

func (uc *AuthUsecase) LoginUser(dto *domain.LoginUserDTO) (*domain.TokenResponseDTO, error) {
	user, err := uc.userRepository.GetUserByEmail(dto.Email)
	// fmt.Println(user)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	fmt.Print(dto.Password)
	fmt.Println(user.Password)
	errs := uc.Infranstructure.CheckPasswordHash(dto.Password, user.Password)
	if !errs {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	accessToken, refreshToken, err := uc.Infranstructure.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Store refresh token in the database
	user.RefreshToken = refreshToken
	user.AccessToken = accessToken
	err = uc.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &domain.TokenResponseDTO{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (uc *AuthUsecase) RefreshTokens(refreshToken string) (*domain.TokenResponseDTO, error) {
	// Validate the refresh token
	token, err := uc.Infranstructure.ValidateToken(refreshToken)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	email := claims["email"].(string)
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	// Check if the provided refresh token matches the stored one
	token, err = uc.Infranstructure.ValidateToken(user.RefreshToken)
	if err != nil || !token.Valid {
		return nil, errors.New("login required")
	}
	if user.RefreshToken != refreshToken {
		return nil, errors.New("invalid refresh token")
	}

	// Generate new tokens
	accessToken, _, err := uc.Infranstructure.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Update the user's refresh token
	user.AccessToken = accessToken
	err = uc.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &domain.TokenResponseDTO{
		AccessToken: accessToken,
	}, nil
}

func (uc *AuthUsecase) ForgotPassword(dto *ForgotPasswordRequestDTO) error {
	user, err := uc.userRepository.GetUserByEmail(dto.Email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	resetToken, err := uc.Infranstructure.GenerateResetToken(user.Email)
	if err != nil {
		return err
	}

	return uc.Infranstructure.SendResetEmail(user, resetToken)
}

func (uc *AuthUsecase) ResetPassword(dto *ResetPasswordRequestDTO) error {
	token, err := uc.Infranstructure.ValidateToken(dto.Token)
	if err != nil || !token.Valid {
		return errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	email := claims["email"].(string)
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return uc.userRepository.UpdateUser(user)
}
