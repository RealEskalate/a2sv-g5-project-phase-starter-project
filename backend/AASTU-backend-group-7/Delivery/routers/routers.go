package routers

import (
	custommongo "blogapp/CustomMongo"
	"blogapp/Domain"
	email_service "blogapp/Infrastructure/email_service"

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
	refreshtokencol := DataBase.Collection("RefreshTokens")

	// Initialize the collections
	customUserCol := custommongo.NewMongoCollection(usercol)
	customBlogCol := custommongo.NewMongoCollection(blogcol)

  customRefreshTokenCol := custommongo.NewMongoCollection(refreshtokencol)
	BlogCollections = Domain.BlogCollections{

		Users: customUserCol,
		Blogs: customBlogCol,
		RefreshTokens: customRefreshTokenCol,
	}
	// Initialize the router
	Router = gin.Default()

	// go to auth router
	AuthRouter()

	// go to blog router
	BlogRouter()

  // go to refresh token router
	RefreshTokenRouter()
	Router.POST("/sendemail", sendemail)
	return Router
}

func sendemail(c *gin.Context) {
	em := email_service.NewMailTrapService()
	err := em.SendEmail("abel.bekele@a2sv.org", "Test", "This is a test email", "catagory1")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Email sent successfully",
		})
	}
}
