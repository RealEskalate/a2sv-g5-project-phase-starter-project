package usecases

import (
	domain "AAiT-backend-group-2/Domain"
	infrastructure "AAiT-backend-group-2/Infrastructure"
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	validator *infrastructure.ValidatorService
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration, validator *infrastructure.ValidatorService) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		validator: validator,
	}
}

func (uu *userUsecase) GetAllUsers(c context.Context,) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	users, err := uu.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uu *userUsecase) GetUserByID(c context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) CreateUser(c context.Context, user domain.User) error {
	if err := infrastructure.ValidateStruct(uu.validator, user); err != nil {
		return fmt.Errorf("validation error: %v", err.Error())
	}

	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err := uu.userRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) UpdateUser(c context.Context, id string, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err := uu.userRepository.Update(ctx, id, user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) DeleteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err := uu.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) Login(c context.Context, user *domain.User) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	token, err := uu.userRepository.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (uu *userUsecase) RefreshToken(c context.Context, token string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	newToken, err := uu.userRepository.RefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func (uu *userUsecase) PromoteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.PromoteUser(ctx, id)
}

func (uu *userUsecase) DemoteAdmin(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.DemoteAdmin(ctx, id)
}