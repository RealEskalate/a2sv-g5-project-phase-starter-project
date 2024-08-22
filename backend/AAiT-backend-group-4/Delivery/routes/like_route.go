package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	controllers "aait-backend-group4/Delivery/Controllers"
	repositories "aait-backend-group4/Repositories"
	usecases "aait-backend-group4/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewLikeRoute sets up the routes for handling like and dislike actions.
func NewLikeRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	// Initialize repositories
	likeRepo := repositories.NewLikeRepository(db, env.LikeCollection)
	blogRepo := repositories.NewBlogRepository(db, env.BlogCollection)

	// Initialize use cases
	likeUsecase := usecases.NewLikeUsecase( blogRepo, likeRepo, timeout)

	// Initialize the controller with the use case
	lctrl := controllers.LikeController{
		LikeUsecase: likeUsecase,
	}

	// Define the like and dislike routes
	likeRoutes := group.Group("/blogs")
	{
		likeRoutes.POST("/:blog_id/like/:user_id", lctrl.Like)      // Route for adding a like
		likeRoutes.POST("/:blog_id/dislike/:user_id", lctrl.Dislike) // Route for adding a dislike
	}
}
