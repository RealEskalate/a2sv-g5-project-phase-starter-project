package infrastructure

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoDBConfig struct {
	URI string
	Database string
}

func NewMongoDBConfig(uri,database string) *MongoDBConfig{
	return &MongoDBConfig{
		URI: uri,
		Database: database,
	}
}

func (c *MongoDBConfig) Connect() (*mongo.Client,error){
	clientOptions := options.Client().ApplyURI(c.URI)
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()

	client,err := mongo.Connect(ctx,clientOptions)
	if err != nil {
		return nil,err
	}
	if err := client.Ping(ctx,nil);err != nil {
		return nil,err
	}

	log.Println("Connected to MongoDB")
	return client,nil
}