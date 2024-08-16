package routers

// import (
// 	"group3-blogApi/delivery/controllers/adminController"
// 	"group3-blogApi/delivery/controllers/authController"

// 	"github.com/gin-gonic/gin"
// )

// func SetUpAdmin(router *gin.Engine) {
// 	admin := router.Group("/admin")
// 	{
// 		// Admin Routes
// 		admin.GET("/profile", authMiddleware, adminController.Profile)
// 		admin.PUT("/update", authMiddleware, adminController.Update)
// 		admin.POST("/upload-image", authMiddleware, adminController.UploadImage)
// 		admin.POST("/logout", authMiddleware, authController.Logout)
// 		admin.POST("/reset-password", authMiddleware, authController.ResetPassword)
// 		admin.POST("/refresh-token", authController.RefreshToken)

// 		// User Routes
// 		admin.GET("/users", authMiddleware, adminController.GetUsers)
// 		admin.GET("/users/:id", authMiddleware, adminController.GetUser)
// 		admin.DELETE("/users/:id", authMiddleware, adminController.DeleteUser)
// 		admin.PUT("/users/:id/role", authMiddleware, adminController.UpdateUserRole)

// 		// Blog Routes
// 		admin.GET("/blogs", authMiddleware, adminController.GetBlogs)
// 		admin.GET("/blogs/:id", authMiddleware, adminController.GetBlog)
// 		admin.POST("/blogs", authMiddleware, adminController.CreateBlog)
// 		admin.PUT("/blogs/:id", authMiddleware, adminController.UpdateBlog)
// 		admin.DELETE("/blogs/:id", authMiddleware, adminController.DeleteBlog)
// 		admin.PUT("/blogs/:id/visibility", authMiddleware, adminController.UpdateBlogVisibility)

// 	}
// }