package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MongoNoFilter() primitive.D {
	return primitive.D{{}}
}
