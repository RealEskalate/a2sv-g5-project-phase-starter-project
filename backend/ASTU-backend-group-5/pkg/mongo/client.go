package mongo


import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewMongoStorage(mongoURL string)(*mongo.Client, error){

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	
	if err != nil {
		return nil, err
    }
    
	return client, nil
}