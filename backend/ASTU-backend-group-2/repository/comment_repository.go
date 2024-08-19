package repository

import "go.mongodb.org/mongo-driver/mongo"

type CommentRepository struct {
	// This is a struct that will hold the mongo client and the collections
	// This will be used to interact with the database

	// This is the mongo collection that will be used to interact with the database
	Collection *mongo.Collection
}

func NewCommentRepository(client *mongo.Client) *CommentRepository {
	// This is a function that will return a new instance of the CommentRepository struct
	// This will be used to interact with the database

	// This will return a new instance of the CommentRepository struct
	return &CommentRepository{
		// This will set the Collection field to the comment collection from the database
		Collection: client.Database("blogApp").Collection("comments"),
	}
}
