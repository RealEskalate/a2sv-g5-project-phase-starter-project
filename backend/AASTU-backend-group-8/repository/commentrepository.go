package repository

import (
	"context"
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	db *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) *CommentRepository {
	return &CommentRepository{
		db: db.Collection("comments"),
	}
}

type CommentRepositoryInterface interface {
	AddComment(comment *domain.Comment) error
	GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error)
	UpdateComment(commentID primitive.ObjectID, content string) error
	DeleteComment(commentID primitive.ObjectID) error
}

// AddComment adds a new comment to the collection
func (r *CommentRepository) AddComment(comment *domain.Comment) error {
	_, err := r.db.InsertOne(context.TODO(), comment)
	return err
}

// GetCommentsByBlogID retrieves all comments for a specific blog post
func (r *CommentRepository) GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error) {
	var comments []domain.Comment
	filter := bson.M{"blog_id": blogID}
	cursor, err := r.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var comment domain.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

// UpdateComment updates the content of an existing comment
func (r *CommentRepository) UpdateComment(commentID primitive.ObjectID, content string) error {
	filter := bson.M{"_id": commentID}
	update := bson.M{"$set": bson.M{"content": content}}
	_, err := r.db.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeleteComment deletes a comment by its ID
func (r *CommentRepository) DeleteComment(commentID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(context.TODO(), bson.M{"_id": commentID})
	return err
}
