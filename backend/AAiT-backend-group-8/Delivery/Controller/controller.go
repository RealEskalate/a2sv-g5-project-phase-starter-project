package controller

import (
	interfaces "AAiT-backend-group-8/Interfaces"
	usecase "AAiT-backend-group-8/Usecase"

	"github.com/redis/go-redis/v9"
)

type Controller struct {
	blogUseCase    interfaces.IBlogUseCase
	commentUseCase *usecase.CommentUseCase
	UserUseCase    interfaces.IUserUseCase
	LikeUseCase    *usecase.LikeUseCase
	AiUseCase      interfaces.IAiUsecase
	rdb            *redis.Client
	cacheUseCase   interfaces.ICacheUseCase
}

func NewController(commentUseCase *usecase.CommentUseCase, userUseCase interfaces.IUserUseCase, likeUseCase *usecase.LikeUseCase, blogUseCase interfaces.IBlogUseCase, aiUseCase interfaces.IAiUsecase, rdb *redis.Client, cacheUseCase interfaces.ICacheUseCase) *Controller {
	return &Controller{
		blogUseCase:    blogUseCase,
		commentUseCase: commentUseCase,
		UserUseCase:    userUseCase,
		LikeUseCase:    likeUseCase,
		rdb:            rdb,
		cacheUseCase:   cacheUseCase,
		AiUseCase:      aiUseCase,
	}
}
