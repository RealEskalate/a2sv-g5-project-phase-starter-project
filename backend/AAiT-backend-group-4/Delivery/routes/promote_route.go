package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	controllers "aait-backend-group4/Delivery/Controllers"
	repositories "aait-backend-group4/Repositories"
	usecases "aait-backend-group4/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPromoteRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, env.UserCollection)
	pc := controllers.PromoteController{
		UserUsecase: usecases.NewPromoteUsecase(ur, timeout),
		Env:         env,
	}

	group.PUT("/user/promote/:id", pc.Promote)
}
