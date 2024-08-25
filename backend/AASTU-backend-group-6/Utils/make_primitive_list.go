package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func MakePrimitiveList(size int) []primitive.ObjectID {
    return make([]primitive.ObjectID, size)
}