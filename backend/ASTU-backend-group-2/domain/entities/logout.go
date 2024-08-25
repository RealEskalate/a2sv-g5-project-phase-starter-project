package entities

import "context"

type LogOutUsecase interface {
	LogOut(c context.Context, refreshid string)error
}
	
