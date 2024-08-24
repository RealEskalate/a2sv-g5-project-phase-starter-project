package user_usecase

import (
	"blog-api/domain"
	"context"
)

func (uc *userUsecase) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.userRepo.GetByUsernameOrEmail(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	return *u, nil
}
