package main

import (
	controller "AAiT-backend-group-8/Delivery/Controller"
	Router "AAiT-backend-group-8/Delivery/Routes"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	"AAiT-backend-group-8/Infrastructure/mongodb"
	usecase "AAiT-backend-group-8/Usecase"

	"context"
)

var SecretKey = "123456abcd"

func main() {
	mongoClient := mongodb.InitMongoDB()

	dbName := "started-project"
	userCollection := mongodb.CreateCollection(mongoClient, dbName, "users")
	tokenCollection := mongodb.CreateCollection(mongoClient, dbName, "token")
	blogCollection := mongodb.CreateCollection(mongoClient, dbName, "blogs")
	commentCollection := mongodb.CreateCollection(mongoClient, dbName, "comments")
	likeCollection := mongodb.CreateCollection(mongoClient, dbName, "likes")

	userRepo := mongodb.NewUserRepository(userCollection, context.TODO())
	ts := infrastructure.NewTokenService(SecretKey)
	ps := infrastructure.NewPasswordService()
	tr := mongodb.NewTokenRepository(tokenCollection, context.TODO())
	ms := infrastructure.NewMailService()
	//	ts := infrastructure.NewTokenService(SECRET_KEY)
	infra := infrastructure.NewInfrastructure()

	blogRepo := mongodb.NewBlogRepository(blogCollection)
	blogUseCase := usecase.NewBlogUseCase(blogRepo)

	commentRepo := mongodb.NewCommentRepository(commentCollection, context.TODO())

	commentUseCase := usecase.NewCommentUseCase(commentRepo, *infra, ts)
	userUseCase := usecase.NewUserUseCase(userRepo, ts, ps, tr, ms)
	likeRepo := mongodb.NewLikeRepository(likeCollection, context.TODO())
	likeUseCase := usecase.NewLikeUseCase(*likeRepo, *infra)

	ctrl := controller.NewController(commentUseCase, userUseCase, likeUseCase, blogUseCase)

	r := Router.InitRouter(ctrl)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
