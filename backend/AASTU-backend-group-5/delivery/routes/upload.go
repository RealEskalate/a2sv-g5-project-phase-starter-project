package routes

import (
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/repository"
	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

func NewUploadRoute(group *gin.RouterGroup , user_repo repository.UserRepository) {
	repo := repository.NewUploadRepository(user_repo)
	uc := usecase.NewUploadUsecase(*repo)
	ctrl := controller.NewUploadController(*uc)	
	group.POST("api/upload/:id" , ctrl.UplaodImg())
}