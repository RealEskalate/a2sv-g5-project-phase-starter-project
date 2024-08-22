package repository

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/domain/dtos"
	"AAiT-backend-group-6/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentRepository struct {
	database   mongo.Database
	collection string
}

func NewCommentRepository(db mongo.Database, collection string) domain.CommentRepository {
	return &commentRepository{
		database:   db,
		collection: collection,
	}
}

func (r *commentRepository) CreateComment(c context.Context, comment *domain.Comment) error {
	collection := r.database.Collection(r.collection)
	comment.ID = primitive.NewObjectID() // Generate a new ObjectID
	_, err := collection.InsertOne(c, comment)
	return err
}

func (r *commentRepository) GetComment(c context.Context, id string) (*domain.Comment, error) {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var comment domain.Comment
	err = collection.FindOne(c, bson.M{"_id": objectID}).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Comment not found
		}
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) UpdateComment(c context.Context, comment *dtos.UpdateDto, commentID primitive.ObjectID) error {
	collection := r.database.Collection(r.collection)

	var oldComment domain.Comment

	err := collection.FindOne(c, bson.M{"_id": commentID}).Decode(&oldComment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err // Comment not found
		}
		return nil
	}

	_, err = collection.UpdateOne(
		c,
		bson.M{"_id": commentID},
		bson.M{"$set": comment},
	)
	return err
}

func (r *commentRepository) DeleteComment(c context.Context, id string) error {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectID})
	return err
}
