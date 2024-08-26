package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlogByIdPipeline(blog_id primitive.ObjectID) mongo.Pipeline {
	pipeline := mongo.Pipeline{
		bson.D{{Key : "$match", Value: bson.D{{Key:"_id", Value: blog_id}}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "comment"},
			{Key: "localField", Value: "comment_ids"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "comments"},
		}}},
	}
	return pipeline
}


