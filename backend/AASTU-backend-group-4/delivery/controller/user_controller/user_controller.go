package user_controller

import (
	"blog-api/domain/user"
)

type userController struct {
	usecase user.UserUsecase
}
