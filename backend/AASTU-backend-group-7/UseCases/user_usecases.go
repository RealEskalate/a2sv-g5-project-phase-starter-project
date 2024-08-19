package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUseCases struct {
	userRepository Domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(service_reference Domain.UserRepository) *userUseCases {
	return &userUseCases{
		userRepository: service_reference,
		contextTimeout: time.Second * 10,
	}
}

func (uc *userUseCases) GetUsers(c *gin.Context) ([]*Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetUsers(ctx)

}

func (uc *userUseCases) GetUsersById(c *gin.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetUsersById(ctx, id, current_user)

}

func (uc *userUseCases) CreateUser(c *gin.Context, user *Domain.User) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.CreateUser(ctx, user)

}

func (uc *userUseCases) UpdateUsersById(c *gin.Context, id primitive.ObjectID, user Domain.User, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.UpdateUsersById(ctx, id, user, current_user)

}

func (uc *userUseCases) DeleteUsersById(c *gin.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.DeleteUsersById(ctx, id, current_user)

}

func (uc *userUseCases) PromoteUser(c *gin.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.PromoteUser(ctx, id, current_user)

}

func (uc *userUseCases) DemoteUser(c *gin.Context, id primitive.ObjectID, current_user Domain.AccessClaims) (Domain.OmitedUser, error, int) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.DemoteUser(ctx, id, current_user)

}
