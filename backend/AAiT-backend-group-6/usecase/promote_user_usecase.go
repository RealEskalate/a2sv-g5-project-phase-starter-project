package usecase

import (
	"AAiT-backend-group-6/domain"
	"context"
	"errors"
	"time"
)

type promoteUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewPromoteUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.PromoteUsecase {
	return &promoteUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *promoteUsecase) PromoteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	user, err := pu.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return errors.New("user with the given userID is not found")
	}
	updatedUser := &domain.User{
		ID: user.ID,
		User_type: "ADMIN",
	}
	
	return pu.userRepository.UpdateUser(ctx, updatedUser)
}

func (pu *promoteUsecase) DemoteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	user, err := pu.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return errors.New("user with the given userID is not found")
	}
	updatedUser := &domain.User{
		ID: user.ID,
		User_type: "USER",
	}
	
	return pu.userRepository.UpdateUser(ctx, updatedUser)
}