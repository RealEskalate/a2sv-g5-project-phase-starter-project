package user

import (
	"context"
	"errors"
)

func (u *UserUsecase) DeleteUser(id string) error {
	user, err := u.repo.FindUserById(context.Background(), id)

	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	err = u.repo.DeleteUser(context.Background(), id)

	return err

}

func (u *UserUsecase) AdminRemoveUser(UserId string) error {

	user, err := u.repo.FindUserById(context.Background(), UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	err = u.repo.DeleteUser(context.Background(), UserId)
	if err != nil {
		return err
	}
	return nil
}
