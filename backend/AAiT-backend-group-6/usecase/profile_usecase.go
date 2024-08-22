package usecase

import (
	"AAiT-backend-group-6/domain"
	"context"
	"time"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) UpdateProfile(c context.Context, userID string, profile *domain.Profile) error {
	panic("unimplemented")
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Username: user.Username, Email: user.Email}, nil
}
