package routers

// import (
// 	"time"

// 	config "github.com/aait.backend.g5.main/backend/Config"
// 	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
// 	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
// 	repository "github.com/aait.backend.g5.main/backend/Repository"
// 	usecases "github.com/aait.backend.g5.main/backend/UseCases"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func NewBlogRouter(env *config.Env, database mongo.Database, group *gin.RouterGroup) {

// 	blog_repo := repository.NewBlogRepository(&database)

// 	jwt_service := infrastructure.NewJwtService()
// 	redis_cache := infrastructure.NewRedisCache()
// 	blog_helper := infrastructure.NewBlogHelper()
// 	cache_TTL := time.Hour * 1

// 	// instantiate blogController
// 	BlogController := &controllers.BlogController{
// 		BlogUsecase: usecases.NewblogUsecase(blog_repo, redis_cache, *env, cache_TTL, blog_helper),
// 		JwtService:  jwt_service,
// 	}

// 	group.POST("/createBlog", BlogController.CreateBlog)

// 	group.POST("/getBlog/:id", BlogController.GetBlog)
// 	group.POST("/getBlogs", BlogController.GetBlogs)
// 	group.POST("/searchBlogs", BlogController.SearchBlogs)

// 	group.POST("/trackPopularity", BlogController.HandelTrackPopularity)

// 	group.POST("/updateBlog", BlogController.UpdateBlog)
// 	group.POST("/deleteBlog", BlogController.DeleteBlog)
// }
