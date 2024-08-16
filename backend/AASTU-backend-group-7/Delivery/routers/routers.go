package routers

import (
	config "blogapp/Config"
	custommongo "blogapp/CustomMongo"
	"blogapp/Domain"
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
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	//initialize the collections
	usercol := DataBase.Collection("Users")
	blogcol := DataBase.Collection("Blogs")

	// Initialize the collections
	customUserCol := custommongo.NewMongoCollection(usercol)
	customBlogCol := custommongo.NewMongoCollection(blogcol)
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
