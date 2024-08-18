package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileUseCases struct {
	userRepository Domain.ProfileRepository
	contextTimeout time.Duration
}

func NewProfileUseCase(service_reference Domain.ProfileRepository) *profileUseCases {
	return &profileUseCases{
		userRepository: service_reference,
		contextTimeout: time.Second * 10,
	}
}

func (uc *profileUseCases) GetProfile(c *gin.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetProfile(ctx, id, current_user)

}

func (uc *profileUseCases) UpdateProfile(c *gin.Context, id primitive.ObjectID, user Domain.User, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.UpdateProfile(ctx, id, user, current_user)

}

func (uc *profileUseCases) DeleteProfile(c *gin.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.DeleteProfile(ctx, id, current_user)

}
