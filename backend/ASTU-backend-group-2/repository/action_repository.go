package repository

import "go.mongodb.org/mongo-driver/mongo"

type ActionRepository struct {
	// This is a struct that will hold the mongo client and the collections
	// This will be used to interact with the database

	// This is the mongo collection that will be used to interact with the database
	Collection *mongo.Collection
}

func NewActionRepository(client *mongo.Client) *ActionRepository {
	// This is a function that will return a new instance of the ActionRepository struct
	// This will be used to interact with the database

	// This will return a new instance of the ActionRepository struct
	return &ActionRepository{
		// This will set the Collection field to the action collection from the database
		Collection: client.Database("blogApp").Collection("actions"),
	}
}
