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
	// Initialize repository
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)

	// Initialize token generator and password service
	tokenGen := infrastracture.NewTokenGenerator() 
	passwordSvc := infrastracture.NewPasswordService()

	// Initialize usecase with dependencies
	userUsecase := usecase.NewUserUsecase(userRepo, tokenGen, passwordSvc)

	// Initialize controller with usecase
	adminController := controllers.NewUserController(userUsecase)

	// Set up admin routes
	admin := router.Group("/admin")
	// admin.Use(infrastracture.AuthMiddleware()) // Ensure AuthMiddleware is properly defined
	{
		// Admin Profile Routes
		admin.GET("/me", infrastracture.RoleMiddleware("admin"), adminController.GetMyProfile)
		admin.PUT("/update", infrastracture.RoleMiddleware("admin"), adminController.UpdateMyProfile)
		admin.POST("/upload-image", infrastracture.RoleMiddleware("admin"), adminController.UploadImage)
		admin.DELETE("/me", infrastracture.RoleMiddleware("admin"), adminController.DeleteMyAccount)

		// User Management Routes
		admin.GET("/users", adminController.GetUsers)
		admin.GET("/users/:id", infrastracture.RoleMiddleware("admin"), adminController.GetUser)
		admin.DELETE("/users/:id", infrastracture.RoleMiddleware("admin"), adminController.DeleteUser)
		admin.PUT("/users/:id", infrastracture.RoleMiddleware("admin"), adminController.UpdateUserRole)

		// Blog Routes (Commented out, uncomment and update if needed)
		// admin.GET("/blogs", infrastracture.RoleMiddleware("admin"), adminController.GetBlogs)
		// admin.GET("/blogs/:id", infrastracture.RoleMiddleware("admin"), adminController.GetBlog)
		// admin.POST("/blogs", infrastracture.RoleMiddleware("admin"), adminController.CreateBlog)
		// admin.PUT("/blogs/:id", infrastracture.RoleMiddleware("admin"), adminController.UpdateBlog)
		// admin.DELETE("/blogs/:id", infrastracture.RoleMiddleware("admin"), adminController.DeleteBlog)
		// admin.PUT("/blogs/:id/visibility", infrastracture.RoleMiddleware("admin"), adminController.UpdateBlogVisibility)
	}
}
