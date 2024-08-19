package usecase

import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type LoginUseCase struct {
	LoginRepository domain.LoginRepository
	ContextTimeout  time.Duration
}


func NewLoginUseCase(loginRepository domain.LoginRepository, timeout time.Duration) domain.LoginUsecase {
	return &LoginUseCase{
		LoginRepository: loginRepository,
		ContextTimeout:  timeout,
	}
}

// Login implements domain.LoginUsecase.
func (l *LoginUseCase) Login(c context.Context, user *domain.UserLogin) (*domain.LoginResponse, error) {
	ctx, cancel:= context.WithTimeout(c, l.ContextTimeout)
	defer cancel()
	return l.LoginRepository.Login(ctx, user)
}

// UpdatePassword implements domain.LoginUsecase.
func (l *LoginUseCase) UpdatePassword(c context.Context, req domain.ChangePasswordRequest, userID string) error {
	ctx, cancel := context.WithTimeout(c, l.ContextTimeout )
	defer cancel()
	return l.LoginRepository.UpdatePassword(ctx, req, userID)
}

