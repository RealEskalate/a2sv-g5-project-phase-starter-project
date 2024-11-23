package main

// import (
// 	"blog-api/delivery/controller/blog_controller"
// 	"blog-api/delivery/controller/user_controller"
// 	"blog-api/infrastructure/bootstrap"
// 	"blog-api/repository/blog_repository"
// 	"blog-api/repository/comment_repository"
// 	"blog-api/repository/like_repository"
// 	"blog-api/repository/user_repository"
// 	"blog-api/usecase/blog_usecase"
// 	"blog-api/usecase/user_usecase"
// 	"time"
// )

// func Initialize() {
// 	app := bootstrap.App()
// 	defer app.CloseDBConnection()

// 	blogCollection := app.Mongo.Database("blog-api").Collection("blog")
// 	userCollection := app.Mongo.Database("blog-api").Collection("user")
// 	commentCollection := app.Mongo.Database("blog-api").Collection("comment")
// 	likeCollection := app.Mongo.Database("blog-api").Collection("like")

// 	blogRepository := blog_repository.NewBlogRepository(blogCollection)
// 	userRepository := user_repository.NewUserRepository(userCollection)
// 	commentRepository := comment_repository.NewCommentRepository(commentCollection)
// 	likeRepository := like_repository.NewLikeRepository(likeCollection)

// 	ctxTimeout := time.Duration(app.Env.ContextTimeout)

// 	userUsecase := user_usecase.NewUserUsecase(userRepository, ctxTimeout)
// 	blogUsecase := blog_usecase.NewBlogUsecase(blogRepository, commentRepository, likeRepository, ctxTimeout)

// 	userController := user_controller.NewUserController(userUsecase, app.Env)
// 	blogController := blog_controller.NewBlogController(blogUsecase)
// }
