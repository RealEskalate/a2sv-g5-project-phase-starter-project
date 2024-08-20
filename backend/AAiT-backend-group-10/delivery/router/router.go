package router

import (
	"os"

	"aait.backend.g10/delivery/controllers"
	inf "aait.backend.g10/infrastructures"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) {
	router := gin.Default()

	userRepo := repositories.NewUserRepository(db, os.Getenv("USER_COLLECTION"))

	blogRepo := repositories.NewBlogRepository(db, os.Getenv("BLOG_COLLECTION"))
	blogUseCase := usecases.NewBlogUseCase(blogRepo, userRepo)
	blogController := controllers.NewBlogController(blogUseCase)

	commentRepo := repositories.NewCommentRepository(db, os.Getenv("COMMENT_COLLECTION_NAME"))
	commentController := controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo),
	}

	likeRepo := repositories.NewLikeRepository(db, os.Getenv("LIKE_COLLECTION_NAME"))
	likeController := controllers.LikeController{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo),
	}

	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	jwtService := inf.JwtService{JwtSecret: os.Getenv("JWT_SECRET")}

	router.POST("/blogs", blogController.CreateBlog)
	router.GET("/blogs", blogController.GetAllBlogs)
	router.GET("/blogs/:id", blogController.GetBlogByID)
	router.PUT("/blogs/:id", blogController.UpdateBlog)
	router.DELETE("/blogs/:id", blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", blogController.AddView)
	router.GET("/blogs/search", blogController.SearchBlogs)

	router.PATCH("/users/promote", userController.PromoteUser)

	router.GET("/comment/:blog_id", inf.AuthMiddleware(&jwtService), commentController.GetComments)
	router.GET("/comment_count/:blog_id", inf.AuthMiddleware(&jwtService), commentController.GetCommentsCount)
	router.POST("/comment", inf.AuthMiddleware(&jwtService), commentController.AddComment)
	router.PUT("/comment/:id", inf.AuthMiddleware(&jwtService), commentController.UpdateComment)
	router.DELETE("/comment/:id", inf.AuthMiddleware(&jwtService), commentController.DelelteComment)

	router.PUT("/like", inf.AuthMiddleware(&jwtService), likeController.LikeBlog)
	router.DELETE("/like", inf.AuthMiddleware(&jwtService), likeController.DeleteLike)
	router.GET("/like/:blog_id", inf.AuthMiddleware(&jwtService), likeController.BlogLikeCount)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
