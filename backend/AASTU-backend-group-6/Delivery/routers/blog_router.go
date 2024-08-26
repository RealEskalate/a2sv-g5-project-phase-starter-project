package routers

import (
	"blogs/Delivery/controllers"
	domain "blogs/Domain" // Add this line to import the "domain" package

	"github.com/gin-gonic/gin"
)

func NewBlogrouter(blogRouter *gin.RouterGroup, controller controllers.BlogController, authHandler gin.HandlerFunc, redisClient *domain.CacheService) {
	// unprotected
	//blogRouter.Use(infrastructure.NewAuthMiddleware(*config).AuthenticationMiddleware())
	blogRouter.GET("/", controller.GetBlogs)
	blogRouter.GET("/:id", controller.GetBlogByID)

	blogRouter.GET("/search", controller.SearchBlogByTitleAndAuthor)
	blogRouter.GET("/filter", controller.FilterBlogsByTag)

	// protected
	blogRouter.GET("/my", authHandler, controller.GetMyBlogs)
	blogRouter.GET("/my/:id", authHandler, controller.GetMyBlogByID)

	blogRouter.POST("/create", authHandler, controller.CreateBlog)
	blogRouter.PUT("/update/:id", authHandler, controller.UpdateBlogByID)
	blogRouter.DELETE("/delete/:id", authHandler, controller.DeleteBlogByID)
	blogRouter.POST("/comment/create", authHandler, controller.CommentOnBlog)
	blogRouter.POST("/react/:id", authHandler, controller.ReactOnBlog)
}
