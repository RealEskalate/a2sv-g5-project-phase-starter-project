package usecase

import (
	"Blog_Starter/domain"
	"context"
	"time"
)


type LogoutUseCase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}


// LogOut implements domain.LogoutUsecase.
func (l *LogoutUseCase) LogOut(c context.Context, userID string) error {
	ctx,cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	_,err := l.UserRepository.UpdateToken(ctx, "", "" ,userID)
	if err != nil {
		return err
	}
	return nil
}

