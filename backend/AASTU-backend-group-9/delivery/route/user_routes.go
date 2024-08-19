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

func NewUserRouter(env *config.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur,timeout),
		Env:           env,
	}
	group.POST("/create_user", uc.CreateUser)
	group.PUT("/update_user/:id", uc.UpdateUser)
	group.DELETE("/delete_user/:id", uc.DeleteUser)
	group.GET("/get_user/:id", uc.GetUser)
	group.GET("/get_all_users", uc.GetUsers)
	group.PATCH("/promote_user/:id", uc.PromoteUser)
	group.PATCH("/demote_user/:id", uc.DemoteUser)
	
}