package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/repository"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// func AuthHandlers(r *gin.Engine, ctrl controller.AuthHandlers) {

// 	r.POST("/login", ctrl.Login())
// 	r.POST("/logout", ctrl.Logout())
// 	r.POST("/signup", ctrl.Signup())
// 	r.POST("/forgot-password", ctrl.ForgotPassword())

// 	r.POST("/reset-password/:id/:token", ctrl.ResetPassword())
// 	r.GET("verify-email", ctrl.VerifyEmail())
// }

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(*db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(*db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}

func NewLogoutRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(*db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLogoutUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/logout", lc.Login)
}

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(*db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
