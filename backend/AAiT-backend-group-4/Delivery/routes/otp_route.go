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

func NewOtpRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	or := repositories.NewOTPRepository(db, env.OtpCollection)
	os := infrastructure.NewOTPService()
	ur := repositories.NewUserRepository(db, env.UserCollection)
	ps := infrastructure.NewPasswordService()
	ou := usecases.NewOtpUsecase(or, timeout, os, ps, *env, ur)
	oc := controllers.OtpController{
		OtpUsecase: ou,
		Env:        *env,
	}

	group.POST("/user/verifyAccount", oc.VerifyOtp)
	group.POST("/user/resendOtp", oc.ResendOtp)
}
