package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	controllers "aait-backend-group4/Delivery/Controllers"
	repositories "aait-backend-group4/Repositories"
	usecases "aait-backend-group4/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBlogRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup, rc redis.Client) {
	// Initialize repositories
	blogRepo := repositories.NewBlogRepository(db, env.BlogCollection)
	userRepo := repositories.NewUserRepository(db, env.UserCollection)

	// Initialize use cases
	blogUsecase := usecases.NewBlogUsecase(blogRepo, userRepo, 2*time.Second, &rc)

	// Initialize the controller with the use case
	bc := controllers.BlogController{
		BlogUsecase: blogUsecase,
	}

	// Define the blog routes with pagination considerations
	blogRoutes := group.Group("/blogs")
	{
		blogRoutes.GET("/:id", bc.FetchByBlogID)
		blogRoutes.POST("", bc.CreateBlog)                         // works
		blogRoutes.PUT("/:id", bc.UpdateBlog)                      // works
		blogRoutes.DELETE("/:id", bc.DeleteBlog)                   // works
		blogRoutes.GET("/", bc.FetchAll)                           // works
		blogRoutes.GET("/popular", bc.FetchByPageAndPopularity)    // works
		blogRoutes.GET("/author/:author_id", bc.FetchByBlogAuthor) // works
		blogRoutes.POST("/tags", bc.FetchByTags)                   // works
		blogRoutes.GET("/search", bc.SearchBlogs)                  // works
		blogRoutes.GET("/title", bc.FetchByBlogTitle)              // works
		blogRoutes.POST("/comments/:id", bc.AddComment)            // works
		blogRoutes.PUT("/comments/:id", bc.UpdateComment)
		blogRoutes.DELETE("/comments/:id", bc.RemoveComment)
	}

}
