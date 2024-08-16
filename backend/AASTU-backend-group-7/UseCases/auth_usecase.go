package usecases

import (
	"blogapp/Domain"
	"context"
	"time"

	"github.com/gin-gonic/gin"
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
func (a *authUseCase) Login(c *gin.Context, user *Domain.User) (string, error, int) {
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
