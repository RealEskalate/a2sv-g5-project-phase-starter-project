package repository

import (
	"context"
	"errors"

	"github.com/RealEskalate/blogpost/database"
	"github.com/RealEskalate/blogpost/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository struct {
	collection database.CollectionInterface
}

func NewCommentRepository(collection database.CollectionInterface) *CommentRepository {
	return &CommentRepository{
		collection: collection,
	}
}

func (CR *CommentRepository) GetComments(post_id string) ([]domain.Comment, error) {
	var comments []domain.Comment
	query := bson.M{"post_id": post_id}

	cursor, err := CR.collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (CR *CommentRepository) CreateComment(post_id string, user_id string) error {
	userObjectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err // Return an error if the conversion fails
	}

	postObjectID, err := primitive.ObjectIDFromHex(post_id)
	if err != nil {
		return err // Return an error if the conversion fails
	}
	comment := domain.Comment{
		ID:     primitive.NewObjectID(),
		UserID: userObjectID,
		PostID: postObjectID,
	}

	_, err = CR.collection.InsertOne(context.TODO(), comment)
	return err
}

func (CR *CommentRepository) DeleteComment(comment_id string) error {
	obID, _ := primitive.ObjectIDFromHex(comment_id)
	query := bson.M{"_id": obID}

	res, err := CR.collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}

	if res.DeletedCount() == 0 {
		return errors.New("no comment with this ID found")
	}

	return nil
}

func (CR *CommentRepository) UpdateComment(comment_id string) error {
	obID, _ := primitive.ObjectIDFromHex(comment_id)
	update := bson.M{"$set": bson.M{"content": "Updated content"}}

	_, err := CR.collection.UpdateOne(context.TODO(), bson.M{"_id": obID}, update)
	return err
}
