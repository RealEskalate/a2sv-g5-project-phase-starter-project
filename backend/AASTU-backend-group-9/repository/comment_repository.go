package repository

import (
	"blog/database"
	"blog/domain"
	"context"
	"errors"
	// "fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "context"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository struct {
	commentdb  database.Database
	collection string
}

// GetComments implements domain.CommentRepository.

func NewCommentRepository(db database.Database, collection string) domain.CommentRepository {
	return &CommentRepository{
		commentdb:  db,
		collection: collection,
	}
}

func (r *CommentRepository) AddComment(ctx context.Context, post_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {

	comment.BlogID = post_id
	comment.AuthorID = userID

	collection := r.commentdb.Collection(r.collection)
	_, err := collection.InsertOne(ctx, comment)
	return err

}
func (r *CommentRepository) GetComments(ctx context.Context, post_id primitive.ObjectID) ([]domain.Comment, error) {
	var comments []domain.Comment
	collection := r.commentdb.Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.M{"blog_id": post_id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after iteration

	for cursor.Next(ctx) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if cursor == nil {
		return nil, err
	}

	return comments, nil
}


func (r *CommentRepository) DeleteComment(ctx context.Context, postID, commentID, userID primitive.ObjectID) error {
	var comment domain.Comment
	collection := r.commentdb.Collection(r.collection)

	// Find the comment by ID and blog ID
	err := collection.FindOne(ctx, bson.M{"_id": commentID, "blog_id": postID}).Decode(&comment)
	if err != nil {
		return err
	}

	// Check if the user is authorized to delete the comment
	if comment.AuthorID != userID {
		return errors.New("you are not authorized to delete this comment")
	}

	// Delete the comment
	_, err = collection.DeleteOne(ctx, bson.M{"_id": commentID})
	if err != nil {
		return err
	}

	return nil
}


func (r *CommentRepository) UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	var oldComment domain.Comment
	collection := r.commentdb.Collection(r.collection)
	err := collection.FindOne(ctx, bson.M{"_id": comment_id, "blog_id": post_id}).Decode(&oldComment)

	if err != nil {
		return err
	}
	if oldComment.AuthorID != userID {
		return errors.New("you are not authorized to update this comment")
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": comment_id}, bson.M{"$set": bson.M{"content": comment.Content}})
	return err
}
