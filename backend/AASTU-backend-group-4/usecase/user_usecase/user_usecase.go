package user_usecase

import (
	"time"

	"blog-api/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	authService    domain.AuthService
	emailService   domain.EmailService
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, authService domain.AuthService, emailService domain.EmailService, timeout time.Duration) *userUsecase {
	return &userUsecase{
		userRepo:       userRepository,
		emailService:  emailService,
		authService:    authService,
		contextTimeout: timeout,
	}
}

// func (uc *UserUsecase) GetByEmail(ctx context.Context, email string) (domain.User, error) {
// 	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
// 	defer cancel()

// 	u, err := uc.repo.GetByEmail(ctx, email)
// 	if err != nil {
// 		return domain.User{}, err
// 	}

// 	return u, nil
// }

// func (uc *UserUsecase) GetByUsername(ctx context.Context, username string) (domain.User, error) {
// 	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
// 	defer cancel()

// 	u, err := uc.repo.GetByUsername(ctx, username)
// 	if err != nil {
// 		return domain.User{}, err
// 	}

// 	return u, nil
// }
