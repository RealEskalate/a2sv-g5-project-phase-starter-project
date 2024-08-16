package usecase

import (
	"blog/domain"
	"blog/internal/tokenutil"
	"blog/internal/userutil"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) RegisterUser(c context.Context, user *domain.AuthSignup) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	if !userutil.ValidateEmail(user.Email) {
		return nil, errors.New("invalid email")
	}
	if !userutil.ValidatePassword(user.Password) {
		return nil, errors.New("invalid password")
	}

	hashedPassword, err := userutil.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	adduser := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		Role:     "user",
	}
	err = su.userRepository.CreateUser(ctx, adduser)
	return &adduser.ID, err
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (su *signupUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (su *signupUsecase) CreateAccessToken(user *domain.AuthSignup, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.AuthSignup, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
