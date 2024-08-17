package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authUseCase struct {
	AuthRepository Domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUseCase(repo Domain.AuthRepository) *authUseCase {
	return &authUseCase{
		AuthRepository: repo,
		contextTimeout: time.Second * 10,
	}
}

// login
func (a *authUseCase) Login(c *gin.Context, user *Domain.User) (Domain.Tokens, error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.AuthRepository.Login(ctx, user)
}

// register
func (a *authUseCase) Register(c *gin.Context, user *Domain.User) (*Domain.OmitedUser, error, int) {
	// return error
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.Register(ctx,user)
}

// logout

func (a *authUseCase) Logout(c *gin.Context, user_id primitive.ObjectID) (error,int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.Logout(ctx, user_id)
}



