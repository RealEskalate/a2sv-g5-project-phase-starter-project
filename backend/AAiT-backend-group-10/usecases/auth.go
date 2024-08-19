package usecases

import (
	"errors"
	"fmt"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/repositories"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserUsecase interface {
	RegisterUser(User *domain.RegisterUserDTO) (interface{}, error)
	LoginUser(dto *domain.LoginUserDTO) (*domain.TokenResponseDTO, error)
	RefreshTokens(refreshToken string) (*domain.TokenResponseDTO, error)
}

type userUsecase struct {
	userRepository  repositories.UserRepositoryInterface
	Infranstructure infrastructures.InfrastructureInterface
}

func NewUserUsecase(ur repositories.UserRepositoryInterface, Infr infrastructures.InfrastructureInterface) UserUsecase {
	return &userUsecase{
		userRepository:  ur,
		Infranstructure: Infr,
	}
}

func (u *userUsecase) RegisterUser(User *domain.RegisterUserDTO) (interface{}, error) {
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

func (uc *userUsecase) LoginUser(dto *domain.LoginUserDTO) (*domain.TokenResponseDTO, error) {
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
		AccessToken: accessToken,
	}, nil
}

func (uc *userUsecase) RefreshTokens(refreshToken string) (*domain.TokenResponseDTO, error) {
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
	if user.RefreshToken != refreshToken {
		return nil, errors.New("invalid refresh token")
	}

	// Generate new tokens
	accessToken, newRefreshToken, err := uc.Infranstructure.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Update the user's refresh token
	user.RefreshToken = newRefreshToken
	err = uc.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &domain.TokenResponseDTO{
		AccessToken: accessToken,
	}, nil
}
