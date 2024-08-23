package routers

import (
	custommongo "blogapp/CustomMongo"
	"blogapp/Domain"
	email_service "blogapp/Infrastructure/email_service"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Router *gin.Engine
var BlogCollections Domain.BlogCollections

func Setuprouter(client *mongo.Client) *gin.Engine {
	// Initialize the database
	DataBase := client.Database("Blog")

	//initialize the user collections
	usercol := DataBase.Collection("Users")
	// blogcol := DataBase.Collection("Blogs")
	refreshtokencol := DataBase.Collection("RefreshTokens")
	resettokencol := DataBase.Collection("ResetTokens")

	// Initialize the custonm user collections
	customUserCol := custommongo.NewMongoCollection(usercol)
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}}, // index in ascending order
			Options: options.Index().SetUnique(true),  // make index unique
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}}, // index in ascending order
			Options: options.Index().SetUnique(true),     // make index unique
		},
	}
	_, err := customUserCol.CreateIndexes(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}

	// customBlogCol := custommongo.NewMongoCollection(blogcol)
	customRefreshTokenCol := custommongo.NewMongoCollection(refreshtokencol)
	customResetTokenCol := custommongo.NewMongoCollection(resettokencol)

	// Initialize the blog collections
	posts := DataBase.Collection("Posts")
	comments := DataBase.Collection("Comments")
	likesDislikes := DataBase.Collection("likesDislikes")
	tags := DataBase.Collection("Tags")
	// Initialize the custom blog collections

	customPostCol := custommongo.NewMongoCollection(posts)
	indexModels = []mongo.IndexModel{

		{
			Keys: bson.D{{Key: "slug", Value: 1}}, // index in ascending order
		},
		{
			Keys: bson.D{{Key: "authorname", Value: 1}}, // index in ascending order
		},
		{
			Keys: bson.D{{Key: "authorid", Value: 1}}, // index in ascending order
		},
		{
			Keys: bson.D{{Key: "tags", Value: 1}}, // index in ascending order
		},
	}
	_, err = customPostCol.CreateIndexes(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}
	// generate custom comment collection
	customCommentCol := custommongo.NewMongoCollection(comments)
	indexModels = []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "postid", Value: 1}}, // index in ascending order
		},
		{
			Keys: bson.D{{Key: "authorid", Value: 1}}, // index in ascending order
		},
	}
	_, err = customCommentCol.CreateIndexes(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}

	customlikesDislikesCol := custommongo.NewMongoCollection(likesDislikes)
	indexModels = []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "postid", Value: 1}}, // index in ascending order
		},
		{
			Keys: bson.D{{Key: "userid", Value: 1}}, // index in ascending order
		},
	}
	_, err = customlikesDislikesCol.CreateIndexes(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}

	// generate custom tag collection
	customTagCol := custommongo.NewMongoCollection(tags)
	indexModels = []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "slug", Value: 1}}, // index in ascending order
			Options: options.Index().SetUnique(true), // make index unique
		},
	}
	_, err = customTagCol.CreateIndexes(context.Background(), indexModels)
	if err != nil {
		panic(err)
	}

	BlogCollections = Domain.BlogCollections{

		Users: customUserCol,
		// Blogs:         customBlogCol,
		RefreshTokens: customRefreshTokenCol,
		ResetTokens:   customResetTokenCol,
		Posts:         customPostCol,
		Comments:      customCommentCol,
		LikesDislikes: customlikesDislikesCol,
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

	// user router
	UserRouter()

	// profile router
	ProfileRouter()
	// chat router
	ChatRouter()

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
