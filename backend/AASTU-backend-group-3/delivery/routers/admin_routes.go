package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpAdmin(router *gin.Engine) {
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    adminController := controllers.NewUserController(userUsecase)
	admin := router.Group("/admin")
	admin.Use(infrastracture.AuthMiddleware())
	{
		// Admin Routes
		admin.GET("/me",infrastracture.RoleMiddleware("admin"), adminController.GetMyProfile)
		admin.PUT("/update",infrastracture.RoleMiddleware("admin"), adminController.UpdateMyProfile)
		admin.POST("/upload-image", infrastracture.RoleMiddleware("admin"),adminController.UploadImage)
		admin.DELETE("/me", infrastracture.RoleMiddleware("admin"), adminController.DeleteMyAccount)


		// User Routes
		admin.GET("/users",infrastracture.RoleMiddleware("admin") ,adminController.GetUsers)
		admin.GET("/users/:id",  infrastracture.RoleMiddleware("admin"),adminController.GetUser)
		admin.DELETE("/users/:id",  infrastracture.RoleMiddleware("admin"),adminController.DeleteUser)
		admin.PUT("/users/:id",  infrastracture.RoleMiddleware("admin"),adminController.UpdateUserRole)

// 		// Blog Routes
// 		admin.GET("/blogs", authMiddleware, adminController.GetBlogs)
// 		admin.GET("/blogs/:id", authMiddleware, adminController.GetBlog)
// 		admin.POST("/blogs", authMiddleware, adminController.CreateBlog)
// 		admin.PUT("/blogs/:id", authMiddleware, adminController.UpdateBlog)
// 		admin.DELETE("/blogs/:id", authMiddleware, adminController.DeleteBlog)
// 		admin.PUT("/blogs/:id/visibility", authMiddleware, adminController.UpdateBlogVisibility)

	}
}