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

	//initialize the user collections
	usercol := DataBase.Collection("Users")
	blogcol := DataBase.Collection("Blogs")
	refreshtokencol := DataBase.Collection("RefreshTokens")

	// Initialize the custonm user collections
	customUserCol := custommongo.NewMongoCollection(usercol)
	customBlogCol := custommongo.NewMongoCollection(blogcol)
	customRefreshTokenCol := custommongo.NewMongoCollection(refreshtokencol)

	// Initialize the blog collections
	posts := DataBase.Collection("Posts")
	comments := DataBase.Collection("Comments")
	likes := DataBase.Collection("Likes")
	tags := DataBase.Collection("Tags")
	// Initialize the custom blog collections
	
	customPostCol := custommongo.NewMongoCollection(posts)
	customCommentCol := custommongo.NewMongoCollection(comments)
	customLikeCol := custommongo.NewMongoCollection(likes)
	customTagCol := custommongo.NewMongoCollection(tags)


	BlogCollections = Domain.BlogCollections{

		Users:         customUserCol,
		Blogs:         customBlogCol,
		RefreshTokens: customRefreshTokenCol,
		Posts:         customPostCol,
		Comments:      customCommentCol,
		Likes:         customLikeCol,
		Tags:          customTagCol,
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
