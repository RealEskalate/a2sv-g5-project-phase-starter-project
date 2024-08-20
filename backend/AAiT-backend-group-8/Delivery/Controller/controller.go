package controller

import (
	domain "AAiT-backend-group-8/Domain"
	usecase "AAiT-backend-group-8/Usecase"
)

type Controller struct {
	commentUseCase usecase.CommentUseCase
	UserUsecase    domain.IUserUseCase
}

func NewController(commentUseCase usecase.CommentUseCase, userUseCase domain.IUserUseCase) *Controller {
	return &Controller{
		commentUseCase: commentUseCase,
		UserUsecase:    userUseCase,
	}
}
