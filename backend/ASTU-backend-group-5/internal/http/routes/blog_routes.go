package routes

// import (
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"github.com/gin-gonic/gin"
// )


// func RegisterBlogRoutes(collection *mongo.Collection, router *gin.Engine) {
// 	// blogRepo := &mongodb.BlogRepositoryMongo{Collection: collection}
// 	// blogUsecase := blog.NewBlogUsecase(blogRepo)
// 	blogHandler := InstantaiteBlogHandler(collection)
// 	blogRoutes := router.Group("/api/v1/blog")

// 	{
// 		blogRoutes.POST("/", blogHandler.CreateBlogHandler)
// 		blogRoutes.GET("/:id", blogHandler.GetBlogByIDHandler)
// 		blogRoutes.PUT("/:id", blogHandler.UpdateBlogHandler)
// 		blogRoutes.DELETE("/:id", blogHandler.DeleteBlogHandler)
// 	}
// }