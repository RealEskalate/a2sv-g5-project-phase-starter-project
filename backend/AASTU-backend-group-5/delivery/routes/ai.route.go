package routes

import (
	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/infrastructure"
	"github.com/gin-gonic/gin"
)

func NewAiRequestRoute(group *gin.RouterGroup) {
	var Ai infrastructure.AI
	ctrl := controller.AI_controller{
		Ai_func: Ai,
	}
	group.POST("api/generate-blog/", ctrl.GenerateBlog())
}