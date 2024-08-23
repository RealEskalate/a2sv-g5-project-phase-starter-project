package main

import (
	controller "AAiT-backend-group-8/Delivery/Controller"
	Router "AAiT-backend-group-8/Delivery/Routes"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	"AAiT-backend-group-8/Infrastructure/mongodb"
	usecase "AAiT-backend-group-8/Usecase"
	"context"
	"log"
)

var SecretKey = "123456abcd"

func main() {
	mongoClient := mongodb.InitMongoDB("mongodb://localhost:27017")
	rdb := infrastructure.InitRedis()

	userCollection := mongoClient.Database("starter-project").Collection("users")
	tokenCollection := mongoClient.Database("starter-project").Collection("token")
	blogCollection := mongoClient.Database("starter-project").Collection("blogs")
	commentCollection := mongoClient.Database("starter-project").Collection("comments")
	likeCollection := mongoClient.Database("starter-project").Collection("likes")
	cacheCollection := mongoClient.Database("starter-project").Collection("cache")
	mongoClient := mongodb.InitMongoDB()

	dbName := "starter-project"
	userCollection := mongodb.CreateCollection(mongoClient, dbName, "users")
	tokenCollection := mongodb.CreateCollection(mongoClient, dbName, "token")
	blogCollection := mongodb.CreateCollection(mongoClient, dbName, "blogs")
	commentCollection := mongodb.CreateCollection(mongoClient, dbName, "comments")
	likeCollection := mongodb.CreateCollection(mongoClient, dbName, "likes")
	cacheCollection := mongoClient.Database("starter-project").Collection("cache")

	userRepo := mongodb.NewUserRepository(userCollection, context.TODO())
	ts := infrastructure.NewTokenService(SecretKey)
	ps := infrastructure.NewPasswordService()
	tr := mongodb.NewTokenRepository(tokenCollection, context.TODO())
	ms := infrastructure.NewMailService()
	//	ts := infrastructure.NewTokenService(SECRET_KEY)
	infra := infrastructure.NewInfrastructure()

	blogRepo := mongodb.NewBlogRepository(blogCollection)
	blogUseCase := usecase.NewBlogUseCase(blogRepo)

	cacheRepo := mongodb.NewCacheRepository(cacheCollection)
	cacheUseCase := usecase.NewCacheUseCase(cacheRepo)

	commentRepo := mongodb.NewCommentRepository(commentCollection, context.TODO())

	commentUseCase := usecase.NewCommentUseCase(commentRepo, *infra, ts)
	userUseCase := usecase.NewUserUseCase(userRepo, ts, ps, tr, ms)
	likeRepo := mongodb.NewLikeRepository(likeCollection, context.TODO())
	likeUseCase := usecase.NewLikeUseCase(*likeRepo, *infra)

	aiService, ai_err := infrastructure.NewGenAIService("AIzaSyCcpZ8utOr8xCRTc-QufZWKSDPIbYz2v7Q")
	if ai_err != nil {
		log.Fatal(ai_err)
	}

	aiblogUsecase := usecase.NewAiBlogUsecase(aiService)

	ctrl := controller.NewController(commentUseCase, userUseCase, likeUseCase, blogUseCase, rdb, cacheUseCase, aiblogUsecase)
	r := Router.InitRouter(ctrl)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
