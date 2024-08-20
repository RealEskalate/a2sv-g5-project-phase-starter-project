package router

import (
	"os"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) {
	router := gin.Default()

	jwtService := infrastructures.JwtService{JwtSecret: os.Getenv("JWT_SECRET")}

	userRepo := repositories.NewUserRepository(db, os.Getenv("USER_COLLECTION"))
	var UserRepo = repositories.NewUserRepositoryMongo(db)
	var Inf = infrastructures.NewInfrastructure()
	var UserUsecases = usecases.NewUserUsecase(UserRepo, Inf)
	type RouteHandler struct {
		UserRouter *controllers.UserController
	}

	var Route = RouteHandler{
		UserRouter: controllers.NewUserController(UserUsecases),
	}
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

	router.POST("/blogs", infrastructures.AuthMiddleware(&jwtService), blogController.CreateBlog)
	router.GET("/blogs", infrastructures.AuthMiddleware(&jwtService), blogController.GetAllBlogs)
	router.GET("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.GetBlogByID)
	router.PUT("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.UpdateBlog)
	router.DELETE("/blogs/:id", infrastructures.AuthMiddleware(&jwtService), blogController.DeleteBlog)
	router.PATCH("/blogs/:id/view", infrastructures.AuthMiddleware(&jwtService), blogController.AddView)
	router.GET("/blogs/search", infrastructures.AuthMiddleware(&jwtService), blogController.SearchBlogs)

	router.PATCH("/users/promote", infrastructures.AuthMiddleware(&jwtService), infrastructures.AdminMiddleWare(), userController.PromoteUser)

	router.GET("/comment/:blog_id", infrastructures.AuthMiddleware(&jwtService), commentController.GetComments)
	router.GET("/comment_count/:blog_id", infrastructures.AuthMiddleware(&jwtService), commentController.GetCommentsCount)
	router.POST("/comment", infrastructures.AuthMiddleware(&jwtService), commentController.AddComment)
	router.PUT("/comment/:id", infrastructures.AuthMiddleware(&jwtService), commentController.UpdateComment)
	router.DELETE("/comment/:id", infrastructures.AuthMiddleware(&jwtService), commentController.DelelteComment)

	router.PUT("/like", infrastructures.AuthMiddleware(&jwtService), likeController.LikeBlog)
	router.DELETE("/like", infrastructures.AuthMiddleware(&jwtService), likeController.DeleteLike)
	router.GET("/like/:blog_id", infrastructures.AuthMiddleware(&jwtService), likeController.BlogLikeCount)

	router.POST("/register", Route.UserRouter.Register)
	router.POST("/login", Route.UserRouter.Login)
	router.POST("/refresh", Route.UserRouter.RefreshToken)
	port := os.Getenv("PORT")
	router.Run(":" + port)

}
