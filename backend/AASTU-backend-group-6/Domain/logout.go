package domain

import "context"

type LogoutUsecase interface {
	Logout(c context.Context, id string) error
	CheckActiveUser(c context.Context, id string) (ActiveUser, error)
}
