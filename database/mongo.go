package database

import (
	"context"
	"loan-management/config"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(config config.Config) (mongoifc.Database, error) {
	uri := config.Database.Url
	opts := options.Client().ApplyURI(uri)
	clnt, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return mongoifc.WrapDatabase(&mongo.Database{}), err
	}
	db := clnt.Database(config.Database.Name)
	return mongoifc.WrapDatabase(db), nil
}
