package user_usecase

import (
	"context"
)

func (u *UserUsecase) DeleteRefreshTokenByUserID(ctx context.Context, userID string) error {
	return u.repo.DeleteRefreshTokenByUserID(ctx, userID)
}
