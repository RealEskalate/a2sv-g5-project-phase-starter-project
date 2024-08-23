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

func NewCommentRepository(db *mongo.Database, collectionName string) *CommentRepository {
	collection := db.Collection(collectionName)
	return &CommentRepository{
		Collection: collection,
	}
}
func (cr *CommentRepository) GetCommentByID(commentID uuid.UUID) (domain.Comment, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: commentID}}
	var comment domain.Comment
	err := cr.Collection.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return comment, domain.ErrCommentNotFound
		}
		return comment, domain.ErrCommentFetchFailed
	}
	return comment, nil
}

// AddComment implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) AddComment(comment domain.Comment) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := cr.Collection.InsertOne(ctx, comment)
	if err != nil {
		return domain.ErrCommentCreationFailed
	}
	return nil
}

// DelelteComment implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) DeleteComment(commentID uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "_id", Value: commentID},
	}
	result, err := cr.Collection.DeleteOne(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		return domain.ErrCommentDeletionFailed
	}
	return nil
}

// GetComments implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) GetComments(blogID uuid.UUID) ([]domain.Comment, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	cursor, err := cr.Collection.Find(ctx, filter)
	if err != nil {
		return nil, domain.ErrCommentFetchFailed
	}
	var comments []domain.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, domain.ErrCommentFetchFailed
	}
	return comments, nil
}

func (cr *CommentRepository) GetCommentsCount(blogID uuid.UUID) (int, *domain.CustomError) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	count, err := cr.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, domain.ErrCommentFetchFailed
	}
	return int(count), nil
}

// UpdateComment implements interfaces.CommentRepositoryInterface.
func (cr *CommentRepository) UpdateComment(updatedComment domain.Comment) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "_id", Value: updatedComment.ID},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "comment", Value: updatedComment.Comment},
		}},
	}
	_, err := cr.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.ErrCommentUpdateFailed
	}
	return nil
}

func (cr *CommentRepository) DeleteCommentsByBlog(blogID uuid.UUID) *domain.CustomError {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{
		{Key: "blog_id", Value: blogID},
	}
	_, err := cr.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return domain.ErrCommentDeletionFailed
	}
	return nil
}