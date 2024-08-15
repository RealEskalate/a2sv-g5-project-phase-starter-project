package usecases

import (
	"blogapp/Domain"
	"time"

	"github.com/gin-gonic/gin"
)

type authUseCase struct {
	AuthRepository Domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUseCase(repo Domain.AuthRepository) (*authUseCase) {
	return &authUseCase{
		AuthRepository: repo,
		contextTimeout: time.Second * 10,
	}
}

// login
func (a *authUseCase) Login(c *gin.Context, user *Domain.User) (string, error, int) {
	// return error
	return "", nil, 0
}

// register
func (a *authUseCase) Register(c *gin.Context, user *Domain.User) (*Domain.OmitedUser, error, int) {
	// return error
	return nil, nil, 0
}
