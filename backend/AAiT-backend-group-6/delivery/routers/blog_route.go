package routers

import (
	delivery "AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"

	"github.com/gin-gonic/gin"
)

func NewBlogRouter(db mongo.Database, gin *gin.Engine) {
	tr := repository.NewBlogRepository(db, domain.CollectionBlogs)
	tc := &delivery.BlogController{
		BlogUsecase: usecase.NewBlogUseCase(
			tr,
		),
	}
	// protectedRoute := gin.Group("")
	publicRoute := gin.Group("")
	// protectedRoute.Use(infrastructure.AdminOnlyMiddleware(), infrastructure.JWTAuthMiddleware())
	publicRoute.GET("/blogs", tc.GetBlogs)
	publicRoute.GET("/blogs/:id", tc.GetBlog)
	publicRoute.POST("/blogs", tc.CreateBlog)
	publicRoute.PUT("/blogs/:id", tc.UpdateBlog)
	publicRoute.DELETE("/blogs/:id", tc.DeleteBlog)
	publicRoute.POST("/blogs/:id", tc.LikeBlog)
	publicRoute.POST("/blogs/:id", tc.UnlikeBlog)
	publicRoute.POST("/blogs/:id", tc.CommentBlog)

}
