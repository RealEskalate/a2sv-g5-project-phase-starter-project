package route

import (
	"blog/config"
	"blog/database"
	"blog/delivery/controller"
	// "blog/delivery/middleware"
	"blog/domain"
	"blog/repository"
	"blog/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {
	Profilerepo := repository.NewUserRepository(db, domain.CollectionUser)
	Profileusecase := usecase.NewProfileUsecase(Profilerepo, timeout)
	ProfileController := &controller.ProfileController{
		ProfileUsecase: Profileusecase,
		Env:          env,
	}

	router.PATCH("/update_profile", ProfileController.UpdateProfile)
}