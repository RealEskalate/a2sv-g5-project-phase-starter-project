package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PaginationByPage(pageNo int64, pageSize int64, popularity string) *options.FindOptions {
	skip := (int64(pageNo) - 1) * pageSize
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	if len(popularity) > 0 {
		findOptions.SetSort(bson.D{{Key: "popularity", Value: -1}})
	}
	findOptions.SetSkip(skip)
	findOptions.SetLimit(pageSize)
	return findOptions
}