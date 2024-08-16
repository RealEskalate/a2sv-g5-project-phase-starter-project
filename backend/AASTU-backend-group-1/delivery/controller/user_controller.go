package controller

import "blogs/domain"

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(uu domain.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: uu,
	}
}
