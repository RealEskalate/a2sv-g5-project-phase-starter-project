package repositories

import (
	"context"
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	Collection *mongo.Collection
}

func NewCommentRepository(db mongo.Database, collectionName string) interfaces.CommentRepositoryInterface {
	collection := db.Collection(collectionName)
	return &CommentRepository{
		Collection: collection,
	}
}

// AddComment implements interfaces.CommentRepositoryInterface.
func (c *CommentRepository) AddComment(comment domain.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	comment.ID = uuid.New()
	_, err := c.Collection.InsertOne(ctx, comment)
	return err
}

// DelelteComment implements interfaces.CommentRepositoryInterface.
func (c *CommentRepository) DelelteComment(blogID uuid.UUID, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "blog_id", Value: blogID},
		{Key: "user_id", Value: userID},
	}
	_, err := c.Collection.DeleteOne(ctx, filter)
	return err
}

// GetComments implements interfaces.CommentRepositoryInterface.
func (c *CommentRepository) GetComments(blogID uuid.UUID) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	cursor, err := c.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// UpdateComment implements interfaces.CommentRepositoryInterface.
func (c *CommentRepository) UpdateComment(updatedComment domain.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "blog_id", Value: updatedComment.BlogID},
		{Key: "user_id", Value: updatedComment.UserID},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "comment", Value: updatedComment.Comment},
		}},
	}
	_, err := c.Collection.UpdateOne(ctx, filter, update)
	return err
}
