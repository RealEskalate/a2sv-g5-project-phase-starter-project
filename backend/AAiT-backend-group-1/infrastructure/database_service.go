package infrastructure

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService interface {
	GetCollection(collectionName string) *mongo.Collection
}

type databaseService struct {
	client *mongo.Client
	dbName string
}

func NewDatabaseService(dbUri string, dbName string) DatabaseService {
	clientOptions := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &databaseService{client: client, dbName: dbName}
}

func (d *databaseService) GetCollection(collectionName string) *mongo.Collection {
	return d.client.Database(d.dbName).Collection(collectionName)
}
