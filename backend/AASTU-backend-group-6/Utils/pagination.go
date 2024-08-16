package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PaginationByPage(pageNo int, pageSize int) *options.FindOptions{
	skip := (int64(pageNo) - 1) * int64(pageSize)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key:"createdAt", Value : -1}})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(int64(pageSize))
	return findOptions
}