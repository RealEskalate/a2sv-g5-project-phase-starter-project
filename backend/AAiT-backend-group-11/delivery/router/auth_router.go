package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup) {
	token_repo := repository.NewTokenRepository(db)
	acc_tok_secret := env.AccessTokenSecret
	ref_tok_secret := env.RefreshTokenSecret
	pass_reset_secret := env.PasswordResetSecret

	user_repo := repository.NewUserRepository(db.Collection("users"))
	token_service := service.NewTokenService(acc_tok_secret, ref_tok_secret, token_repo, user_repo)

	user_service := service.NewUserService(user_repo)
	otpRepo := repository.NewOtpRepository(db.Collection("otp"))
	otpService := service.NewOtpService(otpRepo)

	password_token_repo := repository.NewPasswordTokenRepository(db)

	pass_service := utils.NewPasswordService()
	pass_reset_service := service.NewPasswordResetService(pass_reset_secret, user_repo, password_token_repo)
	auth_service := service.NewAuthService(user_service, pass_reset_service, pass_service, token_service, otpService)
	auth_controller := controller.NewAuthController(auth_service, pass_reset_service)

	group.POST("login", auth_controller.Login)
	group.POST("logout", auth_controller.Logout)
	group.GET("refresh", auth_controller.RefreshAccessToken)
	group.POST("register", auth_controller.RegisterUser)
	group.POST("/verify-email", auth_controller.VerifyEmail)
	group.POST("/forgot-password", auth_controller.RequestPasswordReset)
	group.POST("/reset-password", auth_controller.ResetPassword)
	group.POST("/resend-otp", auth_controller.ResendOtp)
}
