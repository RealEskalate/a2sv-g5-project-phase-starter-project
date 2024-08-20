package commentrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/comment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository defines the MongoDB repository for comments.
type Repository struct {
	collection *mongo.Collection
}

// New creates a new Repository for managing comments with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repository {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repository{
		collection: collection,
	}
}

// Save adds a new comment if it does not exist, else updates the existing one.
func (r *Repository) Save(comment comment.Comment) error {

	filter := bson.M{"_id": comment.Id()}
	update := bson.M{
		"$set": bson.M{
			"content": comment.Content(),
			"blogId":  comment.BlogId(),
			"userId":  comment.UserId(),
		},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error saving comment: %w", err)
	}
	return nil
}

// Delete removes a comment by ID.
func (r *Repository) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

// GetCommentById retrieves a comment by ID.
func (r *Repository) GetCommentById(id uuid.UUID) (*comment.Comment, error) {
	filter := bson.M{"_id": id}

	var c comment.Comment

	err := r.collection.FindOne(context.Background(), filter).Decode(&c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	return &c, nil

}

// GetCommentsByBlogId retrieves all comment by blogId.
func (r *Repository) GetCommentsByBlogId(id uuid.UUID) (*[]comment.Comment, error) {
	// id is blog id
	filter := bson.M{"_id": id}

	var comments []comment.Comment

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	for cur.Next(context.Background()) {
		var c comment.Comment

		err := cur.Decode(&c)

		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	err = cur.Close(context.Background())

	if err != nil {
		return nil, err
	}

	return &comments, nil

}
