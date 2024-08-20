package Controller

import (
	"AAiT-backend-group-8/Domain"
	usecase "AAiT-backend-group-8/Usecase"
)

type Controller struct {
	blogUseCase    Domain.IBlogUseCase
	commentUseCase *usecase.CommentUseCase
	UserUseCase    Domain.IUserUseCase
}

func NewController(blogUseCase Domain.IBlogUseCase, commentUseCase *usecase.CommentUseCase, userUseCase Domain.IUserUseCase) *Controller {
	return &Controller{
		blogUseCase:    blogUseCase,
		commentUseCase: commentUseCase,
		UserUseCase:    userUseCase,
	}
}
