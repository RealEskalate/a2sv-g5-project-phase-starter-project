package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func FilterTaskByUserID(user_id primitive.ObjectID) primitive.D {
	return primitive.D{primitive.E{Key: "creater_id", Value: user_id}}
}
