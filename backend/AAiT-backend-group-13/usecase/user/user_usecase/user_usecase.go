package userusecase

import (
	er "github.com/group13/blog/domain/errors"
	iuser "github.com/group13/blog/usecase/user/i_user"
)

type ChangeStatus struct {
	ToAdmin  bool
	UserRepo iuser.IUserRepository
}

func (cs ChangeStatus) Handle(username string) (bool, error) {
	err := cs.UserRepo.Save(username)

	if err != nil {
		return false, er.NewUnexpected("server error")
	}

	return true, nil
}
