package controller

import (
	domain "AAiT-backend-group-8/Domain"
	usecase "AAiT-backend-group-8/Usecase"
)

type Controller struct {
	blogUseCase    domain.IBlogUseCase
	commentUseCase usecase.CommentUseCase
	UserUsecase    domain.IUserUseCase
	LikeUseCase    usecase.LikeUseCase
}

func NewController(commentUseCase usecase.CommentUseCase, userUseCase domain.IUserUseCase, likeUseCase usecase.LikeUseCase, blogUseCase domain.IBlogUseCase) *Controller {
	return &Controller{
		blogUseCase:    blogUseCase,
		commentUseCase: commentUseCase,
		UserUsecase:    userUseCase,
		LikeUseCase:    likeUseCase,
	}
}
