package user

import (
	"blogApp/internal/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	userId := user.ID
	dbUser, err := u.repo.FindUserById(context.Background(), userId.Hex())
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("invalid user")
	}
	dbUser.Profile = user.Profile
	dbUser.Updated = primitive.NewDateTimeFromTime(time.Now())
	return u.repo.UpdateUser(context.Background(), dbUser)
}
