package main

import (
	controller "AAiT-backend-group-8/Delivery/Controller"
	Router "AAiT-backend-group-8/Delivery/Routes"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	"AAiT-backend-group-8/Infrastructure/mongodb"
	usecase "AAiT-backend-group-8/Usecase"

	"context"
)

var SecretKey = "123456abed"

func main() {
	mongoClient := mongodb.InitMongoDB("mongodb://localhost:27017")

	userCollection := mongoClient.Database("starter-project").Collection("users")
	tokenCollection := mongoClient.Database("starter-project").Collection("token")
	blogCollection := mongoClient.Database("starter-project").Collection("blogs")
	commentCollection := mongoClient.Database("starter-project").Collection("comments")
	likeCollection := mongoClient.Database("starter-project").Collection("likes")

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
