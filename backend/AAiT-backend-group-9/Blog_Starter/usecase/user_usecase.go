package usecase

import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type UserUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		userRepo:       userRepo,
		contextTimeout: timeout,
	}
}

// DeleteUser implements domain.UserUsecase.
func (u *UserUsecase) DeleteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepo.DeleteUser(ctx, userID)
}

// GetAllUser implements domain.UserUsecase.
func (u *UserUsecase) GetAllUser(c context.Context) ([]*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	users, err := u.userRepo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	var userResponses []*domain.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &domain.UserResponse{
			UserID:         user.UserID,
			Username:       user.Username,
			Email:          user.Email,
			Name:           user.Name,
			Bio:            user.Bio,
			ContactInfo:    user.ContactInfo,
			Role:           user.Role,
			IsActivated:    user.IsActivated,
			ProfilePicture: user.ProfilePicture,
		})
	}
	return userResponses, nil
}

// GetUserByEmail implements domain.UserUsecase.
func (u *UserUsecase) GetUserByEmail(c context.Context, email string) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &domain.UserResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		Name:           user.Name,
		Bio:            user.Bio,
		ContactInfo:    user.ContactInfo,
		Role:           user.Role,
		IsActivated:    user.IsActivated,
		ProfilePicture: user.ProfilePicture,
	}, nil

}

// GetUserByID implements domain.UserUsecase.
func (u *UserUsecase) GetUserByID(c context.Context, userID string) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &domain.UserResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		Name:           user.Name,
		Bio:            user.Bio,
		ContactInfo:    user.ContactInfo,
		Role:           user.Role,
		IsActivated:    user.IsActivated,
		ProfilePicture: user.ProfilePicture,
	}, nil
}

// PromoteUser implements domain.UserUsecase.
func (u *UserUsecase) PromoteUser(c context.Context, userID string) error {
	panic("unimplemented")
}

// DemoteUser implements domain.UserUsecase.
func (u *UserUsecase) DemoteUser(c context.Context, userID string) error {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUsecase.
func (u *UserUsecase) UpdateUser(c context.Context, user *domain.UserUpdate, userID string) (*domain.UserResponse, error) {
	panic("unimplemented")
}
