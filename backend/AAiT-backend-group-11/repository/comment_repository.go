package repository

import (
	"context"
	"time"

	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	collection *mongo.Collection
	ctx context.Context
}

func NewCommentRepository(collection *mongo.Collection, ctx context.Context ) interfaces.CommentRepository {
	return &commentRepository{
		collection: collection,
		ctx: ctx,
	}
}

func (cr *commentRepository) AddComment( comment *entities.Comment) (*entities.Comment, error) {
	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()

	_, err := cr.collection.InsertOne(cr.ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *commentRepository) DeleteComment( commentId string) error {
	// TODO: return error if comment does not exist

	objID, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return err
	}

	_, err = cr.collection.DeleteOne(cr.ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) GetCommentsByBlogPostId( blogPostId string) ([]entities.Comment, error) {
	objID, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return nil, err
	}

	cursor, err := cr.collection.Find(cr.ctx, bson.M{"blogPostId": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(cr.ctx)

	var comments []entities.Comment
	for cursor.Next(cr.ctx) {
		var comment entities.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (cr *commentRepository) UpdateComment( comment *entities.Comment) (*entities.Comment, error) {
	

	filter := bson.M{"_id": comment.ID}
	update := bson.M{
		"$set": comment,
	}

	_, err := cr.collection.UpdateOne(cr.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
