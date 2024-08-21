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

func NewBlogRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	// Initialize repositories
	blogRepo := repositories.NewBlogRepository(db, env.BlogCollection)
	userRepo := repositories.NewUserRepository(db, env.UserCollection)

	// Initialize use cases
	blogUsecase := usecases.NewBlogUsecase(blogRepo, userRepo, 2*time.Second)

	// Initialize the controller with the use case
	bc := controllers.BlogController{
		BlogUsecase: blogUsecase,
	}

	// Set up the router
	router := gin.Default()

	// Define the blog routes with pagination considerations
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.POST("/", bc.CreateBlog)
		blogRoutes.PUT("/:id", bc.UpdateBlog)
		blogRoutes.DELETE("/:id", bc.DeleteBlog)
		blogRoutes.GET("/", bc.FetchAll)
		blogRoutes.GET("/popular", bc.FetchByPageAndPopularity)
		blogRoutes.GET("/author/:author_id", bc.FetchByBlogAuthor)
		blogRoutes.POST("/tags", bc.FetchByTags)
		blogRoutes.GET("/search", bc.SearchBlogs)
		blogRoutes.GET("/title", bc.FetchByBlogTitle)
		blogRoutes.POST("/:id/comments", bc.AddComment)
		blogRoutes.PUT("/:id/comments", bc.UpdateComment)
		blogRoutes.DELETE("/:id/comments", bc.RemoveComment)
	}

	// Start the server
	router.Run(":8080")
}
