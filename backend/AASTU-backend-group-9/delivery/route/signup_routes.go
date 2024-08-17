package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"
	"blog/domain"
	"time"
	"github.com/gin-gonic/gin"
	"blog/repository"
	"blog/usecase"
)

// Setup sets up the routes for the application

func NewSignupRouter(env *config.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	tr := repository.NewMongoTokenRepository(db, domain.TokenCollection)
	or := repository.NewOTPRepository(db, domain.CollectionOTP)
	config.GoogleConfig(env)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur,tr,or,timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
	group.POST("/verify-otp", sc.VerifyOTP)
	group.GET("/google_login", sc.GoogleLogin)
	group.GET("/google_callback", sc.GoogleCallback)
}
