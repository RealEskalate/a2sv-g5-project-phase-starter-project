package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	gemini "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/aiutil"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/repository"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewChatRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewChatRepository(*db, entities.CollectionChat)
	aiService := gemini.NewAIUtil(env)
	cc := controller.ChatController{
		ChatUsecase: usecase.NewChatUsecase(cr, *aiService, timeout),
		Env:         env,
	}

	group.GET("chats/", cc.GetChats)
	group.POST("chat/", cc.CreateChat)
	group.GET("chat/:id", cc.GetChat)
	group.DELETE("chat/:id", cc.DeleteChat)
	group.POST("chat/:id/interact", cc.SendMessage)
}
