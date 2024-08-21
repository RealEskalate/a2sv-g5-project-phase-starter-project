package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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
func (u *UserUsecase) DeleteUser(c context.Context, userID string, password string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	// check if the incoming password and users password are the same
	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	//compare the hashed password

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("password incorrect")
	}

	err = u.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
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
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.Role == "admin" {
		return errors.New("user already has an admin role")
	}

	user, err = u.userRepo.UpdateRole(ctx, "admin", userID)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		return err
	}

	return nil

}

// DemoteUser implements domain.UserUsecase.
func (u *UserUsecase) DemoteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.Role == "user" {
		return errors.New("user already has a user role")
	}

	_, err = u.userRepo.UpdateRole(ctx, "user", userID)
	if err != nil {
		return err
	}

	if user.Role != "user" {
		return err
	}

	return nil

}

// UpdateUser implements domain.UserUsecase.
func (u *UserUsecase) UpdateUser(c context.Context, user *domain.UserUpdate, userID string) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	_, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	updatedUser, err := u.userRepo.UpdateProfile(ctx, user, userID)
	if err != nil {
		return nil, err
	}
	return &domain.UserResponse{
		UserID:         updatedUser.UserID,
		Username:       updatedUser.Username,
		Email:          updatedUser.Email,
		Name:           updatedUser.Name,
		Bio:            updatedUser.Bio,
		ContactInfo:    updatedUser.ContactInfo,
		Role:           updatedUser.Role,
		IsActivated:    updatedUser.IsActivated,
		ProfilePicture: updatedUser.ProfilePicture,
	}, nil
}

// UpdateProfilePicture implements domain.UserUsecase.
func (u *UserUsecase) UpdateProfilePicture(c context.Context, profilePicPath string, userID string) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	_, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	updatedUser, err := u.userRepo.UpdateProfilePicture(ctx, profilePicPath, userID)
	if err != nil {
		return nil, err
	}
	return &domain.UserResponse{
		UserID:         updatedUser.UserID,
		Username:       updatedUser.Username,
		Email:          updatedUser.Email,
		Name:           updatedUser.Name,
		Bio:            updatedUser.Bio,
		ContactInfo:    updatedUser.ContactInfo,
		Role:           updatedUser.Role,
		IsActivated:    updatedUser.IsActivated,
		ProfilePicture: updatedUser.ProfilePicture,
	}, nil
}
