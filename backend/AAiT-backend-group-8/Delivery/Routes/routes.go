package Router

import (
	controller "AAiT-backend-group-8/Delivery/Controller"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	"github.com/gin-gonic/gin"
)

var SECRET_KEY = "123456abcd"

func InitRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()

	router.POST("/register", controller.RegisterUser)
	router.GET("/verify", controller.VerifyEmail)
	router.POST("/login", controller.Login)
	router.POST("/refresh", controller.RefreshToken)
	router.POST("/logout", infrastructure.UserMiddleware(SECRET_KEY), controller.Logout)

	router.POST("/users/delete/:email", infrastructure.UserMiddleware(SECRET_KEY), infrastructure.AdminMiddleware(), controller.DeleteUser)
	router.POST("/users/promote/:email", infrastructure.UserMiddleware(SECRET_KEY), infrastructure.AdminMiddleware(), controller.PromoteUser)
	router.POST("/users/demote/:email", infrastructure.UserMiddleware(SECRET_KEY), infrastructure.SuperAdminMiddleware(), controller.DemoteUser)

	router.POST("/comment/:blogID", infrastructure.UserMiddleware(SECRET_KEY), controller.CreateComment)
	router.GET("/comment/:blogID", infrastructure.UserMiddleware(SECRET_KEY), controller.GetComments)
	router.PATCH("/comment/:commentID", infrastructure.UserMiddleware(SECRET_KEY), controller.UpdateComment)
	router.DELETE("/comment/:commentID", infrastructure.UserMiddleware(SECRET_KEY), controller.DeleteComment)

	router.POST("/forgot-password", controller.ForgotPassword)
	router.GET("/store-token", controller.StoreToken)
	router.POST("/reset-password", controller.ResetPassword)

	router.POST("/like/:blogID", infrastructure.UserMiddleware(SECRET_KEY), controller.LikeBlog)
	router.GET("/like/:blogID", infrastructure.UserMiddleware(SECRET_KEY), controller.GetLikes)

	router.POST("/blog", infrastructure.UserMiddleware(SECRET_KEY), controller.CreateBlog)
	router.GET("/blogs", infrastructure.UserMiddleware(SECRET_KEY), controller.GetBlogs)
	router.GET("blogs/:id", infrastructure.UserMiddleware(SECRET_KEY), controller.GetBlogByID)
	router.DELETE("/blogs/:id", infrastructure.UserMiddleware(SECRET_KEY), controller.DeleteBlog)
	router.PUT("/blogs", infrastructure.UserMiddleware(SECRET_KEY), controller.UpdateBlog)
	router.GET("/blogs/search", infrastructure.UserMiddleware(SECRET_KEY), controller.SearchBlog)
	router.GET("/blogs/filter", infrastructure.UserMiddleware(SECRET_KEY), controller.SearchBlog)

	// Ai routes
	router.POST("/generate-blog", infrastructure.UserMiddleware(SECRET_KEY), controller.GenerateBlog)
	router.POST("/blog-suggest-improvement", infrastructure.UserMiddleware(SECRET_KEY), controller.SuggestImprovements)
	return router
}
