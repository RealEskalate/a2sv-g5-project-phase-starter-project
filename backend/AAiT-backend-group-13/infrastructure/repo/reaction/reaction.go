package reactionrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/reaction"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository defines the MongoDB repository for reactions.
type Repository struct {
	collection *mongo.Collection
}

// New creates a new Repository for managing reactions with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repository {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repository{
		collection: collection,
	}
}

func (r Repository) Save(reaction reaction.Reaction) error {
	filter := bson.M{"_id": reaction.ID()}
	update := bson.M{

		"$set": bson.M{
			"isLike": reaction.IsLike(),
			"blogId": reaction.BlogId(),
			"userId": reaction.UserId(),
		},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error saving comment: %w", err)
	}
	return nil
}

// Delete removes a reaction by ID.
func (r Repository) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

// Find Reaction By UserId and BlogId
func (r Repository) FindReactionById(id uuid.UUID) (*reaction.Reaction, error) {
	filter := bson.M{"_id": id}

	var reaction reaction.Reaction

	err := r.collection.FindOne(context.Background(), filter).Decode(&reaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	return &reaction, nil
}

func (r Repository) FindReactionByBlogId(blogId uuid.UUID) (*[]reaction.Reaction, error) {
	filter := bson.M{"_id": blogId}
	var reactions []reaction.Reaction

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	for cur.Next(context.Background()) {
		var r reaction.Reaction

		err := cur.Decode(&r)

		if err != nil {
			return nil, err
		}
		reactions = append(reactions, r)
	}

	return &reactions, nil
}

func (r *Repository) FindReactionByUserIdAndBlogId(userId, blogId string) (*reaction.Reaction, error) {
	var reaction reaction.Reaction

	filter := bson.M{
		"userId": userId,
		"blogId": blogId,
	}

	err := r.collection.FindOne(context.Background(), filter).Decode(&reaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No reaction found
		}
		return nil, err // Other error occurred
	}

	return &reaction, nil
}
