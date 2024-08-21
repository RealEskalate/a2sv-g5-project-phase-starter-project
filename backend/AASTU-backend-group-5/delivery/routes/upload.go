package routes

import (
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/gin-gonic/gin"
)

func NewIploadRoute(group *gin.Engine) {
	ctrl := controller.UploadController{}
	group.POST("api/upload/" , ctrl.UplaodImg())
}