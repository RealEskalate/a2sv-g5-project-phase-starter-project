package routers

import (
	custommongo "blogapp/CustomMongo"
	"blogapp/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine
var BlogCollections Domain.BlogCollections

func Setuprouter(client *mongo.Client) *gin.Engine {
	// Initialize the database
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
