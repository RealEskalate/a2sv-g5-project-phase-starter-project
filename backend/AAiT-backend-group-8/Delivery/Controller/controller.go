package controllers

import usecase "AAiT-backend-group-8/Usecase"

type Controller struct {
	commentUseCase usecase.CommentUseCase
}

func NewController(commentUseCase usecase.CommentUseCase) *Controller {
	return &Controller{
		commentUseCase: commentUseCase,
	}
}
