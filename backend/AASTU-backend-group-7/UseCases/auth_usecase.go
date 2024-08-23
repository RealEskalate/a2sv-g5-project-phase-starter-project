package usecases

import (
	"blogapp/Domain"
	"blogapp/Dtos"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUseCase struct {
	AuthRepository Domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUseCase(repo Domain.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		AuthRepository: repo,
		contextTimeout: time.Second * 10,
	}
}

// login
func (a *AuthUseCase) Login(c *gin.Context, user *Dtos.LoginUserDto) (Domain.Tokens, error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.AuthRepository.Login(ctx, user)
}

// register
func (a *AuthUseCase) Register(c *gin.Context, user *Dtos.RegisterUserDto) (*Domain.OmitedUser, error, int) {
	// return error
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.Register(ctx, user)
}

// logout

func (a *AuthUseCase) Logout(c *gin.Context, user_id primitive.ObjectID) (error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.Logout(ctx, user_id)
}

func (a *AuthUseCase) ForgetPassword(c *gin.Context, email string) (error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.ForgetPassword(ctx, email)
}

func (a *AuthUseCase) ResetPassword(c *gin.Context, email string, password string, resetToken string) (error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.ResetPassword(ctx, email, password, resetToken)
}

func (a *AuthUseCase) GoogleLogin(c *gin.Context) string {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.GoogleLogin(ctx)
}

func (a *AuthUseCase) CallbackHandler(c *gin.Context, code string) (Domain.Tokens, error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.CallbackHandler(ctx, code)
}

func (a *AuthUseCase) ActivateAccount(c *gin.Context, token string) (error, int) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.AuthRepository.ActivateAccount(ctx, token)
}
