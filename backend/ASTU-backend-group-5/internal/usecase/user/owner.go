package user

import (
	"context"
	"errors"
)

func (u *UserUsecase) PromoteToAdmin(UserId string) error {

	user, err := u.repo.FindUserById(context.Background(), UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	user.Role = "admin"
	err = u.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) DemoteFromAdmin(UserId string) error {

	user, err := u.repo.FindUserById(context.Background(), UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	user.Role = "user"
	err = u.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
