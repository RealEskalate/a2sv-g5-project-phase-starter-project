package route

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewChatRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	cc := controller.ChatController{
		Env: env,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	group.POST("/chat", cc.Chat(ctx))
}
