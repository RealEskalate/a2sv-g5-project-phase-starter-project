package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) CreateUser(c context.Context, user *domain.User) error {
	return nil
}

func (uu *userUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (uu *userUsecase) GetUserById(c context.Context, userId string) (*domain.User, error) {
	return nil, nil
}
func (uu *userUsecase) GetUsers(c context.Context, limit int64, page int64) (*[]domain.User, mongopagination.PaginationData, error) {
	return nil, mongopagination.PaginationData{}, nil
}

func (uu *userUsecase) UpdateUser(c context.Context, userID string, updatedUser *domain.User) error {
	return nil
}

func (uu *userUsecase) DeleteUser(c context.Context, userID string) error {
	return nil
}

func (uu *userUsecase) IsUserActive(c context.Context, userID string) (bool, error) {
	return false, nil
}

func (uu *userUsecase) ResetUserPassword(c context.Context, userID string, resetPassword *domain.ResetPassword) error {
	return nil
}

func (uu *userUsecase) UpdateUserPassword(c context.Context, userID string, updatePassword *domain.UpdatePassword) error {
	return nil
}

func (uu *userUsecase) PromoteUserToAdmin(c context.Context, userID string) error {
	return nil
}

func (uu *userUsecase) DemoteAdminToUser(c context.Context, userID string) error {
	return nil
}
