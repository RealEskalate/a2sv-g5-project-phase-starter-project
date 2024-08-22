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

func NewForgotPasswordRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, env.UserCollection)
	oi := infrastructure.NewOTPService()
	ps := infrastructure.NewPasswordService()

	// Initialize the controller with the use case
	fc := controllers.ForgotPasswordController{
		OtpService:            infrastructure.NewOTPService(),
		ForgotPasswordUsecase: usecases.NewPasswordUsecase(ur, oi, timeout, ps),
		Env:                   env,
	}

	group.POST("/forgot-password", fc.ForgotPassword)
	group.GET("/forgot-password", fc.ServePage)
	group.POST("/submit-change-password", fc.VerifyForgotPassowrd)
}
