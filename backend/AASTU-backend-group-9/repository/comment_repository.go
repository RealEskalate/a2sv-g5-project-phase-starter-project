package repository

import (
	"blog/database"
	"blog/domain"
	"context"
	"errors"

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

func (r *CommentRepository) GetComments(ctx context.Context, post_id primitive.ObjectID) (*domain.Comment, error) {
	var blog domain.Comment
	collection := r.commentdb.Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.M{"blog_id": post_id})
	if err != nil {
		return nil, err
	}
	err = cursor.Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *CommentRepository) DeleteComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID) error {
	var comment domain.Comment
	collection := r.commentdb.Collection(r.collection)
	err := collection.FindOne(ctx, bson.M{"_id": comment_id, "blog_id": post_id}).Decode(&comment)

	if err != nil {
		return err
	}
	if comment.AuthorID != userID {
		return errors.New("you are not authorized to delete this comment")
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": comment_id})
	return err
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
