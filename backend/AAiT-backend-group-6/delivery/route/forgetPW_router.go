package route

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/infrastructure"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"
	"time"

	"github.com/gin-gonic/gin"
)


func NewFogetPWRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)

	fpu := usecase.NewForgetPWUsecase(ur, timeout, *emailService)
	uu := usecase.NewUserUsecase(ur, timeout)

	fpc := controller.ForgetPWController{
		Userusecase: uu,
		ForgetPWUsecase: fpu,
		Env: env,
	}

	group.POST("/forget-password", fpc.ForgetPW)
	group.POST("/recover-password", fpc.ResetPW)
}