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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)
	ur := repository.NewUserRepository(db, domain.UserCollection)
	su := usecase.NewSignupUsecase(ur, timeout, *emailService)
	uu := usecase.NewUserUsecase(ur, timeout)
	sc := controller.NewSignupController(uu, su, env)

	group.POST("/signup", sc.Signup)
	group.POST("/verify_email", sc.VerifyEmail)
}