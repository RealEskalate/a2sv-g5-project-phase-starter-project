package routers

import (
	"group3-blogApi/config/db"
	"group3-blogApi/delivery/controllers"
	"group3-blogApi/infrastracture"
	"group3-blogApi/repository"
	"group3-blogApi/usecase"

	"github.com/gin-gonic/gin"
)

func SetUpUser(router *gin.Engine) {
	//user routes
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authController := controllers.NewUserController(userUsecase)

	blogRepo := repository.NewBlogRepositoryImpl(db.BlogCollection)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(blogUsecase)

	user := router.Group("/user")
	user.Use(infrastracture.AuthMiddleware())

	{
		// user.GET("/profile", authMiddleware, userController.Profile)
		// user.PUT("/update", authMiddleware, userController.Update)
		// user.POST("/upload-image", authMiddleware, userController.UploadImage)
		// user.POST("/logout", authMiddleware, authController.Logout)
		// user.POST("/reset-password", authMiddleware, authController.ResetPassword)
		user.POST("/blogs", blogController.CreateBlog)
		user.POST("/refresh-token", authController.RefreshToken)

	}
}
