package user_controller

import (
	"blog-api/domain"
	"blog-api/infrastructure/bootstrap"
)

type UserController struct {
	usecase domain.UserUsecase
	Env     *bootstrap.Env
}
