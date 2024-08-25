package route

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/delivery/middleware"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"

	"github.com/gin-gonic/gin"
)

func NewReactionRouter(env *bootstrap.Env, db mongo.Database, gin *gin.Engine) {
	rr := repository.NewReactionRepository(db, domain.CollectionReactions)
	br := repository.NewBlogRepository(db, domain.CollectionBlogs)
	// ur := repository.NewUserRepository(db, domain.UserCollection)
	Ru := usecase.NewReactionUseCase(rr, br)
	Rc := controller.ReactionController{
		ReactionUseCase: Ru,
	}
	protectedRoute := gin.Group("")
	// publicRoute := gin.Group("")
	protectedRoute.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	protectedRoute.POST("/blog/:blog_id/like", Rc.LikeBlog)
	protectedRoute.POST("/blog/:blog_id/unlike", Rc.UnLikeBlog)
	protectedRoute.DELETE("/blog/:blog_id/delete", Rc.DeleteLike)

}
