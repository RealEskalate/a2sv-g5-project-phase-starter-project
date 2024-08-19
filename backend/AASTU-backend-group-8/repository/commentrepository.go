package repository

import (
	"context"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository struct {
	collection domain.Collection
}
type CommentRepositoryInterface interface {
	AddComment(comment *domain.Comment) error
	GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error)
	UpdateComment(commentID primitive.ObjectID, content string) error
	DeleteComment(commentID primitive.ObjectID) error
}

func NewCommentRepository(col domain.Collection) *CommentRepository {
	return &CommentRepository{collection: col}
}

func (r *CommentRepository) AddComment(comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

func (r *CommentRepository) GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var comments []domain.Comment
	filter := bson.M{"blog_id": blogID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment domain.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) UpdateComment(commentID primitive.ObjectID, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": commentID}
	update := bson.M{"$set": bson.M{"content": content}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *CommentRepository) DeleteComment(commentID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": commentID})
	return err
}
