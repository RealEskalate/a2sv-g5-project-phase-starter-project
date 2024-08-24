package user_usecase

import (
	"blog-api/domain"
	"context"
)

func (uc *userUsecase) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.userRepo.GetByUsernameOrEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return *u, nil
}
