package routers

import (
	config "blogapp/Config"
	custommongo "blogapp/CustomMongo"
	"blogapp/Domain"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var BlogCollections Domain.BlogCollections

func Setuprouter() *gin.Engine {
	// Initialize the database
	client := config.ConnectDB()
	log.Println(client)
	DataBase := client.Database("Blog")

	// Defer the closing of the database
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize the collections
	customUserCol := custommongo.NewMongoCollection(DataBase.Collection("Users"))
	customBlogCol := custommongo.NewMongoCollection(DataBase.Collection("Blogs"))
	BlogCollections = Domain.BlogCollections{

		Users: customUserCol,
		Blogs: customBlogCol,
	}
	// Initialize the router
	Router = gin.Default()

	// go to auth router
	AuthRouter()

	return Router
}
