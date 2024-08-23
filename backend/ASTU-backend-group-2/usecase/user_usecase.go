package usecase

import (
	"context"
	"errors"
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

func (uu *userUsecase) CreateUser(c context.Context, user *domain.User) (*domain.User, error) {
	return uu.userRepository.CreateUser(c, user)
}

func (uu *userUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	return uu.userRepository.GetUserByEmail(c, email)
}

func (uu *userUsecase) GetUserById(c context.Context, userId string) (*domain.User, error) {
	return uu.userRepository.GetUserById(c, userId)
}
func (uu *userUsecase) GetUsers(c context.Context, limit int64, page int64) (*[]domain.User, mongopagination.PaginationData, error) {
	return nil, mongopagination.PaginationData{}, nil
}

func (uu *userUsecase) UpdateUser(c context.Context, userID string, updatedUser *domain.User) error {
	_, err := uu.userRepository.UpdateUser(c, userID, updatedUser)
	return err
}

func (uu *userUsecase) ActivateUser(c context.Context, userID string) error {
	_, err := uu.userRepository.ActivateUser(c, userID)
	return err
}

func (uu *userUsecase) DeleteUser(c context.Context, userID string) error {
	return uu.userRepository.DeleteUser(c, userID)
}

func (uu *userUsecase) IsUserActive(c context.Context, userID string) (bool, error) {
	return uu.userRepository.IsUserActive(c, userID)
}

func (uu *userUsecase) IsOwner(c context.Context) (bool, error) {
	return uu.userRepository.IsOwner(c)
}

func (uu *userUsecase) UpdateUserPassword(c context.Context, userID string, updatePassword *domain.UpdatePassword) error {
	return uu.userRepository.UpdateUserPassword(c, userID, updatePassword)
}

func (uu *userUsecase) PromoteUserToAdmin(c context.Context, userID string) error {
	return uu.userRepository.PromoteUserToAdmin(c, userID)
}

func (uu *userUsecase) DemoteAdminToUser(c context.Context, userID string) error {
	user, err := uu.GetUserById(c, userID)
	if err != nil {
		return err
	}
	if user.IsOwner {
		return errors.New("cannot demote owner")
	}
	return uu.userRepository.DemoteAdminToUser(c, userID)
}
func (uu *userUsecase) UpdateProfilePicture(c context.Context, userID string, filename string) error {
	_, err := uu.GetUserById(c, userID)
	if err != nil {
		return err
	}
	return uu.userRepository.UpdateProfilePicture(c, userID, filename)
}
