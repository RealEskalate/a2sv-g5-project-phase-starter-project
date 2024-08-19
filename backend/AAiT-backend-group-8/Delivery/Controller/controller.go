package controllers

import (
	"AAiT-backend-group-8/Domain"
	usecase "AAiT-backend-group-8/Usecase"
)

type Controller struct {
	commentUseCase usecase.CommentUseCase
	blogUseCase    Domain.IBlogUseCase
}

func NewController(blogCase Domain.IBlogUseCase, commentUseCase usecase.CommentUseCase) *Controller {
	return &Controller{
		blogUseCase:    blogCase,
		commentUseCase: commentUseCase,
	}
}
