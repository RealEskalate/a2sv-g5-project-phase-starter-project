package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"time"
)

type promoteUserUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// NewPromoteUsecase creates a new instance of the PromoteUsecase struct.
// It takes a userRepository of type domain.UserRepository and a timeout of type time.Duration as parameters.
// It returns a pointer to the PromoteUsecase struct.
// The PromoteUsecase struct implements the domain.UserUsecase interface.
func NewPromoteUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &promoteUserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// Promote promotes a user identified by the given ID.
// It returns the promoted user and any error encountered during the process.
func (pu *promoteUserUsecase) Promote(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.Promote(ctx, id)
}
