package router

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup)  {
	userRepo := repository.NewUserRepository(db.Collection("users"))
	usr := service.NewUserService(userRepo)
	tokenRepo := repository.NewTokenRepository(db)
	otpRepo := repository.NewOtpRepository(db.Collection("otp"))
	otpService := service.NewOtpService(otpRepo)
	passService := utils.NewPasswordService()
	tokenService := service.NewTokenService(env.AccessTokenSecret, env.RefreshTokenSecret, tokenRepo, userRepo)

	ac := controller.AuthController{
		AuthService : service.NewAuthService(usr, tokenRepo, otpService, passService, tokenService),
		TokenRepo : repository.NewTokenRepository(db),
		Env: env,
	}

	group.POST("/register", ac.RegisterUser)
	group.POST("/verify-email", ac.VerifyEmail)

}
