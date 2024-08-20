package router

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(db *mongo.Database, group *gin.RouterGroup) {
	token_repo := repository.NewTokenRepository(db)
	acc_tok_secret := ""
	ref_tok_secret := ""
	user_repo := repository.NewUserRepository(db.Collection("users"))
	token_service := service.NewTokenService(acc_tok_secret, ref_tok_secret, token_repo, user_repo)

	user_service := service.NewUserService(user_repo)

	pass_service := utils.NewPasswordService()
	pass_reset_service := service.NewPasswordResetService()
	auth_service := service.NewAuthService(user_service, pass_reset_service, pass_service, token_service)
	auth_controller := controller.NewAuthController(auth_service)

	group.POST("/login", auth_controller.Login)
	group.POST("/logout", auth_controller.Logout)
	group.POST("/refresh", auth_controller.RefreshAccessToken)
	group.POST("/register", auth_controller.RegisterUser)

}
