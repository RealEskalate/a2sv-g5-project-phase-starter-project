package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func FilterByTaskAndUserID(user_object_id primitive.ObjectID, blog_object_id primitive.ObjectID) primitive.D {
	filter := primitive.D{
		primitive.E{Key: "_id", Value: blog_object_id},
		primitive.E{Key: "creater_id", Value: user_object_id},
	}
	return filter
}
