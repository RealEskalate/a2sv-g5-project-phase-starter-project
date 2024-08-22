package routes

import (
	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewVerifyEmialRoute(group *gin.RouterGroup, user_collection database.CollectionInterface) {
	repo := repository.NewUserRepository(user_collection)
	user_usecase := usecase.NewUserUseCase(repo)
	

	email_repo := repository.NewEmailVRepo(*repo)
	email_usecase := usecase.NewEmailVUsecase(user_usecase , email_repo)
	email_ctrl := controller.NewEmailVController(email_usecase , user_usecase)

	group.POST("api/verify-email/:id" , email_ctrl.SendVerificationEmail())
	group.GET("api/verify-email/:token" ,email_ctrl.VerifyEmail())
	group.POST("api/forget-password/:id" , email_ctrl.SendForgetPasswordEmail())
	group.GET("/api/forget-password/" , email_ctrl.ForgetPasswordValidate())

}