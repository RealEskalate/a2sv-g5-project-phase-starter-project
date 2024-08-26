package user_controller

import (
	"blog-api/domain"
	"blog-api/infrastructure/bootstrap"
	infrastructure "blog-api/infrastructure/cloudinary"
)

type UserController struct {
	userUsecase domain.UserUsecase
	authService domain.AuthService
	Env         *bootstrap.Env
	Medcont     infrastructure.MediaUpload
}

func NewUserController(userUsecase domain.UserUsecase, authService domain.AuthService, env *bootstrap.Env, Medcont infrastructure.MediaUpload) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		authService: authService,
		Env:         env,
		Medcont:     Medcont,
	}
}
