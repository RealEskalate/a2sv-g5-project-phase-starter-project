package user

import (
	"blogApp/internal/domain"
	"context"
	"errors"
	"strings"
)

func (u *UserUsecase) FindUserById(id string) (*domain.User, error) {
	user, err := u.repo.FindUserById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	users, err := u.repo.GetAllUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserUsecase) FindUserByEmail(email string) (*domain.User, error) {
	user, err := u.repo.FindUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) FindUserByUserName(username string) (*domain.User, error) {
	user, err := u.repo.FindUserByUserName(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) FilterUsers(filter map[string]interface{}) ([]*domain.User, error) {
	validFilters := []string{"role", "email", "username", "firstName", "lastName"}
	//check the filter
	for k, v := range filter {
		if !strings.Contains(strings.Join(validFilters, " "), k) {
			return nil, errors.New("invalid filter")
		}
		if v == nil {
			return nil, errors.New("invalid filter value")
		}
	}

	users, err := u.repo.FilterUsers(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return users, nil
}
