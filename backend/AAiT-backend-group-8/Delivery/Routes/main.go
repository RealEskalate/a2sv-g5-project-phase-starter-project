package main

import (
	controllers "AAiT-backend-group-8/Delivery/Controllers"
	routes "AAiT-backend-group-8/Delivery/Routes"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	repository "AAiT-backend-group-8/Repository"
	usecase "AAiT-backend-group-8/Usecase"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// routes.SetUpRouter(gin.Default())
	myInfrastructure := infrastructure.NewInfrastructure()
	userCollection, err := DataBaseConnector("users")
	if err != nil {
		log.Fatal(err.Error())
	}
	// log("Connected to db")
	userRepository := repository.NewUserRepository(
		*userCollection,
		context.TODO(),
	)

	userUseCase := usecase.NewUserUseCase(
		*userRepository,
		*myInfrastructure,
	)

	userController := controllers.NewController(
		*userUseCase,
	)

	route := routes.NewRouter(
		*myInfrastructure,
		*userController,
	)
	route.SetUpRouter(gin.Default())

}

func DataBaseConnector(collectionName string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	collection := client.Database("Blog_Database").Collection(collectionName)
	fmt.Println("Connected to MongoDB")
	return collection, nil
}
