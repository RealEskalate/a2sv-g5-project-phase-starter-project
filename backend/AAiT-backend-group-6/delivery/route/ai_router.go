package route

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/infrastructure"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"
	"AAiT-backend-group-6/utils"
	"time"

	"github.com/gin-gonic/gin"
)


func NewAiRouter(env *bootstrap.Env,timeout time.Duration,db mongo.Database,group *gin.RouterGroup){
	air := repository.NewAIRepository(db)
	Llc := infrastructure.NewLlmClient(utils.MESSAGE_TELL_ROLE)
	aiu := usecase.NewChatUseCase(air,Llc)
	aic := controller.NewAIController(aiu)

	group.GET("/chat", aic.GetChats)
	group.GET("/chat/:id", aic.GetChat)
	group.POST("/chat", aic.CreateChat)
	group.PUT("/chat/:id", aic.UpdateChat)
	// group.DELETE("/chat/:id", aic.DeleteChat)
}