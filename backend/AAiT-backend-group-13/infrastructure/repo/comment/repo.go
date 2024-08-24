package commentrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repo defines the MongoDB repository for comments.
type Repo struct {
	collection *mongo.Collection
}

// New creates a new Repository for managing comments with the given MongoDB client, database name, and collection name.
func New(client *mongo.Client, dbName, collectionName string) *Repo {
	collection := client.Database(dbName).Collection(collectionName)
	return &Repo{
		collection: collection,
	}
}

// Save adds a new comment if it does not exist, else updates the existing one.
func (r Repo) Save(comment models.Comment) error {

	filter := bson.M{"_id": comment.ID()}
	update := bson.M{
		"$set": bson.M{
			"content": comment.Content(),
			"blogId":  comment.BlogID(),
			"userId":  comment.UserID(),
		},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("error saving comment: %w", err)
	}
	return nil
}

// Delete removes a comment by ID.
func (r Repo) Delete(id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

// GetCommentById retrieves a comment by ID.
func (r Repo) GetCommentById(id uuid.UUID) (*models.Comment, error) {
	filter := bson.M{"_id": id}

	var c CommentDto

	err := r.collection.FindOne(context.Background(), filter).Decode(&c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}
	comment := toCommentModel(&c)

	return comment, nil

}

// GetCommentsByBlogId retrieves all comment by blogId.
func (r Repo) GetCommentsByBlogId(id uuid.UUID) (*[]models.Comment, error) {

	filter := bson.M{"blogId": id}
	var comments []models.Comment

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving blog: %w", err)
	}

	for cur.Next(context.Background()) {
		var c CommentDto

		err := cur.Decode(&c)

		if err != nil {
			return nil, err
		}
		comment := toCommentModel(&c)
		comments = append(comments, *comment)
	}

	err = cur.Close(context.Background())

	if err != nil {
		return nil, err
	}

	return &comments, nil

}

// toBlogModel converts a BlogDTO to a blogmodel.Blog.
func toCommentModel(dto *CommentDto) *models.Comment {
	comment := models.MapComment(models.MapCommentConfig{
		Id:      dto.ID,
		BlogID:  dto.BlogId,
		UserID:  dto.UserId,
		Content: dto.Content,
	})
	return comment
}
