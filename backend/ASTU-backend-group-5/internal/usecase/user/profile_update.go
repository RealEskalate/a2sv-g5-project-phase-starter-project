package user

import (
	"blogApp/internal/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	email := user.Email
	dbUser, err := u.repo.FindUserByEmail(context.Background(),email)
	if err != nil {
		return err
	}
	dbUser.Profile = user.Profile
	dbUser.UserName = user.UserName
	dbUser.Email = user.Email

	dbUser.Updated = primitive.NewDateTimeFromTime(time.Now())
	return u.repo.UpdateUser(context.Background(), dbUser)
}
