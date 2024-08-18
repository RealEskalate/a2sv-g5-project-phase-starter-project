package domain

import "context"

type LogoutUsecase interface {
	LogOut(c context.Context, userID string) error
}