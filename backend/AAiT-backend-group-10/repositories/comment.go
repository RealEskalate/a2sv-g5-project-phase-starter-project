package repositories

import (
	"context"
	"time"

	"aait.backend.g10/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	Collection *mongo.Collection
}

func NewCommentRepository(db mongo.Database, collectionName string) *CommentRepository {
	collection := db.Collection(collectionName)
	return &CommentRepository{
		Collection: collection,
	}
}

// AddComment implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) AddComment(comment domain.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	comment.ID = uuid.New()
	_, err := cr.Collection.InsertOne(ctx, comment)
	return err
}

// DelelteComment implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) DelelteComment(commentID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "id", Value: commentID},
	}
	_, err := cr.Collection.DeleteOne(ctx, filter)
	return err
}

// GetComments implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) GetComments(blogID uuid.UUID) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	cursor, err := cr.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var comments []domain.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *CommentRepository) GetCommentsCount(blogID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	count, err := cr.Collection.CountDocuments(ctx, filter)
	return int(count), err
}

// UpdateComment implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) UpdateComment(commentID uuid.UUID, updatedComment domain.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "id", Value: commentID},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "comment", Value: updatedComment.Comment},
		}},
	}
	_, err := cr.Collection.UpdateOne(ctx, filter, update)
	return err
}
