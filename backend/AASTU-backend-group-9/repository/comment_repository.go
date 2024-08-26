package repository

import (
	"blog/database"
	"blog/domain"
	"context"
	"errors"
	"fmt"

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

func (r *CommentRepository) AddComment(ctx context.Context, post_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) *domain.Error {

	comment.BlogID = post_id
	comment.AuthorID = userID

	collection := r.commentdb.Collection(r.collection)
	_, err := collection.InsertOne(ctx, comment)

	if err != nil {
		return &domain.Error{Message: "Failed to add comment", Err: err}
	}

	return nil

}

func (r *CommentRepository) AddReply(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, reply *domain.Comment) *domain.Error {
	reply.BlogID = post_id
	reply.AuthorID = userID
	reply.ParentID = comment_id

	collection := r.commentdb.Collection(r.collection)
	_, err := collection.InsertOne(ctx, reply)

	if err != nil {
		return &domain.Error{

			StatusCode: 500,
			Message:    "Failed to add reply",
			Err:        err}
	}

	return nil

}

func (r *CommentRepository) IncrementCommentPopularity(ctx context.Context, post_id primitive.ObjectID, commentID primitive.ObjectID, metric string) *domain.Error {
	collection := r.commentdb.Collection(r.collection)
	fmt.Println("test for", metric)
	if metric != "likes" && metric != "dislikes" && metric != "comments" {
		return &domain.Error{
			Message:    "Invalid metric",
			StatusCode: 400,
			Err:        errors.New("invalid metric"),

	}
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": commentID}, bson.M{"$inc": bson.M{metric: 1}})
	if err != nil {
		return &domain.Error{
			Message:    "Failed to increment comment popularity",
			Err:        err,
			StatusCode: 500,
		}

	}

	return nil

}

func (r *CommentRepository) GetComments(ctx context.Context, postID primitive.ObjectID) ([]domain.Comment, *domain.Error) {
	var comments []domain.Comment
	collection := r.commentdb.Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.M{"blog_id": postID})
	if err != nil {
		return nil, &domain.Error{
			Message:    "Failed to get comments",
			Err:        err,
			StatusCode: 500,
		}
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after iteration

	for cursor.Next(ctx) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, &domain.Error{
				Message:    "Failed to decode comment",
				Err:        err,
				StatusCode: 500,
			}
		}
		comments = append(comments, comment)
	}

	if cursor!=nil  { // Check for any errors during iteration
		return nil, &domain.Error{
			Message:    "Error occurred during cursor iteration",
			Err:        err,
			StatusCode: 500,
		}
	}

	if len(comments) == 0 { // Check if no comments were found
		return nil, &domain.Error{
			Message:    "No comments found",
			StatusCode: 404,
			Err:        errors.New("no comments found"),
		}
	}

	return comments, nil
}

func (r *CommentRepository) DeleteComments(ctx context.Context, post_id primitive.ObjectID) *domain.Error {
	collection := r.commentdb.Collection(r.collection)
	_, err := collection.DeleteMany(ctx, bson.M{"blog_id": post_id})
	if err != nil {
		return &domain.Error{
			StatusCode: 500,
			Message:    "Failed to delete comments",
			Err:        err,
		}

	}
	return nil
}

func (r *CommentRepository) DeleteComment(ctx context.Context, postID, commentID, userID primitive.ObjectID) *domain.Error {
	var comment domain.Comment
	collection := r.commentdb.Collection(r.collection)

	// Find the comment by ID and blog ID
	err := collection.FindOne(ctx, bson.M{"_id": commentID, "blog_id": postID}).Decode(&comment)
	if err != nil {
		return &domain.Error{
			StatusCode: 404,
			Message:    "Comment not found",
			Err:        err,
		}
	}
	// Check if the user is authorized to delete the comment
	if comment.AuthorID != userID {
		return &domain.Error{
			StatusCode: 403,
			Message:    "You are not authorized to delete this comment",
			Err:        errors.New("you are not authorized to delete this comment"),
		}
	}
	// Delete the comment
	_, err = collection.DeleteOne(ctx, bson.M{"_id": commentID})
	if err != nil {
		return &domain.Error{
			StatusCode: 500,
			Message:    "Failed to delete comment",
			Err:        err,
		}
	}
	return nil
}

func (r *CommentRepository) UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) *domain.Error {
	var oldComment domain.Comment
	collection := r.commentdb.Collection(r.collection)
	err := collection.FindOne(ctx, bson.M{"_id": comment_id, "blog_id": post_id}).Decode(&oldComment)

	if err != nil {
		return &domain.Error{
			StatusCode: 404,
			Message:    "Comment not found",
			Err:        err,
		}
	}
	if oldComment.AuthorID != userID {
		return &domain.Error{
			StatusCode: 403,
			Message:    "You are not authorized to update this comment",
			Err:        errors.New("you are not authorized to update this comment"),
		}

	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": comment_id}, bson.M{"$set": bson.M{"content": comment.Content}})
	if err != nil {
		return &domain.Error{
			StatusCode: 500,
			Message:    "Failed to update comment",
			Err:        err,
		}
	}
	return nil
}
