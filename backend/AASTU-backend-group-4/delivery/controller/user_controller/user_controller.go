package user_controller

import (
	"blog-api/domain"
	"blog-api/infrastructure/bootstrap"
)

type UserController struct {
	userUsecase domain.UserUsecase
	authService domain.AuthService
	Env         *bootstrap.Env
}

func NewUserController(userUsecase domain.UserUsecase, authService domain.AuthService, env *bootstrap.Env) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		authService: authService,
		Env:         env,
	}
}
