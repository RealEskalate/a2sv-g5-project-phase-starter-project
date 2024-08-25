package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"log"
	"time"
)

type userProfileUsecases struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// NewPromoteUsecase creates a new instance of the PromoteUsecase struct.
// It takes a userRepository of type domain.UserRepository and a timeout of type time.Duration as parameters.
// It returns a pointer to the PromoteUsecase struct.
// The PromoteUsecase struct implements the domain.UserUsecase interface.
func NewUserProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userProfileUsecases{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// Promote promotes a user identified by the given ID.
// It returns the promoted user and any error encountered during the process.
func (pu *userProfileUsecases) Promote(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.Promote(ctx, id)
}

func (pu *userProfileUsecases) GetUserProfile(c context.Context, id string) (domain.User, error) {
	user, err := pu.userRepository.GetByID(c, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (pu *userProfileUsecases) UpdateUserProfile(c context.Context, id string, user domain.UserUpdate) (domain.User, error) {
	log.Println("Updating user profile, ", user)
	updatedUser, err := pu.userRepository.UpdateUser(c, id, user)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (pu *userProfileUsecases) DeleteUserprofile(c context.Context, id string) error {
	return nil
}

func (pu *userProfileUsecases) GetProfileImage(c context.Context, id string) (string, error) {
	user, err := pu.userRepository.GetByID(c, id)
	if err != nil {
		return "", err
	}

	return *user.ProfileImage, nil

}
