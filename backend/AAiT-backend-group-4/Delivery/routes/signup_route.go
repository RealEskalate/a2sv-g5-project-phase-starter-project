package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	controllers "aait-backend-group4/Delivery/Controllers"
	infrastructure "aait-backend-group4/Infrastructure"
	repositories "aait-backend-group4/Repositories"
	usecases "aait-backend-group4/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	// User repository
	ur := repositories.NewUserRepository(db, env.UserCollection)
	// otp repository
	or := repositories.NewOTPRepository(db, env.OtpCollection)
	// otp Service
	os := infrastructure.NewOTPService()
	// password service
	ps := infrastructure.NewPasswordService()
	// otp usecase
	ou := usecases.NewOtpUsecase(or, timeout, os, ps, *env, ur)
	// token Service
	ts := infrastructure.NewTokenService()

	// signup controller
	sc := controllers.SignupController{
		SingupUsecase: usecases.NewSingupUsecase(ur, ou, timeout, ps, ts),
		Env:           env,
	}

	group.POST("/user/register", sc.SignUp)
}
