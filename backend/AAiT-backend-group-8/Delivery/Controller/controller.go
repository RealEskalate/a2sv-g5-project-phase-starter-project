package controller

import (
	domain "AAiT-backend-group-8/Domain"
	usecase "AAiT-backend-group-8/Usecase"
)

type Controller struct {
	commentUseCase usecase.CommentUseCase
	UserUsecase    domain.IUserUseCase
	LikeUseCase    usecase.LikeUseCase
}

func NewController(commentUseCase usecase.CommentUseCase, userUseCase domain.IUserUseCase, likeUseCase usecase.LikeUseCase) *Controller {
	return &Controller{
		commentUseCase: commentUseCase,
		UserUsecase:    userUseCase,
		LikeUseCase:    likeUseCase,
	}
}
