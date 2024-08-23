package controller

import (
	interfaces "AAiT-backend-group-8/Interfaces"
	usecase "AAiT-backend-group-8/Usecase"
)

type Controller struct {
	blogUseCase    interfaces.IBlogUseCase
	commentUseCase *usecase.CommentUseCase
	UserUseCase    interfaces.IUserUseCase
	LikeUseCase    *usecase.LikeUseCase
	AiUseCase      interfaces.IAiUsecase
}

func NewController(commentUseCase *usecase.CommentUseCase, userUseCase interfaces.IUserUseCase, likeUseCase *usecase.LikeUseCase, blogUseCase interfaces.IBlogUseCase, aiUseCase interfaces.IAiUsecase) *Controller {
	return &Controller{
		blogUseCase:    blogUseCase,
		commentUseCase: commentUseCase,
		UserUseCase:    userUseCase,
		LikeUseCase:    likeUseCase,
		AiUseCase:      aiUseCase,
	}
}
