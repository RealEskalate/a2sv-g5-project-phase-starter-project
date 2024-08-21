package domain

import "context"



type LogoutResponse struct {
	Message string `json:"message"`
}

type LogoutUsecase interface {
	Logout(ctx context.Context, token string) (LogoutResponse, error)
}

