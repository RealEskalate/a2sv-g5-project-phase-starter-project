package user_controller

import (
	"blog-api/domain/user"
	"blog-api/infrastructure/bootstrap"
)

type UserController struct {
	usecase user.UserUsecase
	env     *bootstrap.Env
}
