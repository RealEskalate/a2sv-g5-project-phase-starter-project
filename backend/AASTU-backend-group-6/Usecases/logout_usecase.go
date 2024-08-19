package usecases

import (
	domain "blogs/Domain"
	"context"
	"time"
)

type logoutUsecase struct {
	activeUserRepository domain.ActiveUserRepository
	contextTimeout       time.Duration
}

func NewLogoutUsecase(activeUserRepository domain.ActiveUserRepository, timeout time.Duration) domain.LogoutUsecase {
	return &logoutUsecase{
		activeUserRepository: activeUserRepository,
		contextTimeout:       timeout,
	}
}

// Logout implements domain.LogoutUsecase.
func (l *logoutUsecase) Logout(c context.Context, id string, user_agent string) error {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	return l.activeUserRepository.DeleteActiveUser(id, user_agent, ctx)
}
func (l *logoutUsecase) CheckActiveUser(c context.Context, id string) (domain.ActiveUser, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	return l.activeUserRepository.FindActiveUserById(id, ctx)
}
