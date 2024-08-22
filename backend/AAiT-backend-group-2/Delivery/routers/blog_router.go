package routers

import (
  "AAiT-backend-group-2/Delivery/controllers"
  domain "AAiT-backend-group-2/Domain"
  infrastructure "AAiT-backend-group-2/Infrastructure"
  "AAiT-backend-group-2/Repositories/blog_repository"
  "AAiT-backend-group-2/Usecases/blog_usecase"
  "AAiT-backend-group-2/Infrastructure/cache"
  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
)



func NewBlogRouter(db *mongo.Database, group *gin.RouterGroup,configs *domain.Config) {
  redisCache := cache.NewRedisCache(configs.RedisAdr,configs.RedisPass,0)
  blogRepo := blog_repository.NewBlogRepository(db,redisCache)
  blogUsecase := blog_usecase.NewBlogUsecase(blogRepo)
  blogController := controllers.NewBlogController(blogUsecase)
  
  blogRoutes := group.Group("/blogs")
  blogRoutes.GET("", blogController.GetAllBlogs)
  blogRoutes.GET("/:id", blogController.GetBlogByID)
  blogRoutes.GET("/blogFilter", blogController.FilterBlogs)
  blogRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey))
  {
    blogRoutes.POST("", blogController.CreateBlog)
    // blogRoutes.GET("/:id", blogController.GetBlogByID)
    blogRoutes.PUT("/:id", blogController.UpdateBlog)
  }
  blogRoutes.Use(infrastructure.AuthMiddleWare(configs.SecretKey), infrastructure.RoleMiddleware())
  {
    blogRoutes.DELETE("/:id", blogController.DeleteBlog)
  }
  
}
