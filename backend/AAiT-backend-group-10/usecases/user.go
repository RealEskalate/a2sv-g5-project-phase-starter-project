package usecases

import (
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IUserUseCase interface {
	CreateUser(user *domain.User) (*domain.User, *domain.CustomError)
	GetUserByID(id uuid.UUID) (*dto.GetUserResponseDto, *domain.CustomError)
	UpdateUser(requesterID uuid.UUID, user *dto.UserUpdate) *domain.CustomError
	PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError
}

type UserUseCase struct {
	userRepo interfaces.IUserRepository
}

func NewUserUseCase(repo interfaces.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, *domain.CustomError) {
	user.ID = uuid.New()
	return user, u.userRepo.CreateUser(user)
}

func (u *UserUseCase) GetUserByID(id uuid.UUID) (*dto.GetUserResponseDto, *domain.CustomError) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.GetUserResponseDto{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Bio: 	   user.Bio,
		ImageUrl:  user.ImageURL,
		IsAdmin:   user.IsAdmin,
	}, nil
}

func (u *UserUseCase) UpdateUser(requesterID uuid.UUID, user *dto.UserUpdate) *domain.CustomError {
	if requesterID != user.ID {
		return domain.ErrUnAuthorized
	}
	user.UpdatedAt = time.Now().UTC()
	return u.userRepo.UpdateUser(user)
}

func (u *UserUseCase) PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError {
	return u.userRepo.PromoteUser(id, isPromote)
}
