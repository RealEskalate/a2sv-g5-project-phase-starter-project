package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/repository"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPublicResetPasswordRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewResetPasswordRepository(*db, entities.CollectionUser, entities.CollectionResetPassword)
	sc := controller.ResetPasswordController{
		ResetPasswordUsecase: usecase.NewResetPasswordUsecase(ur, timeout),
		Env:                  env,
	}
	group.POST("/forgot-password", sc.ForgotPassword)
	group.POST("/reset-password", sc.ResetPassword)
}
