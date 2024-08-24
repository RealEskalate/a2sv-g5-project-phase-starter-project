package mongodb

import (
	domain "AAiT-backend-group-8/Domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCommentRepository(collection *mongo.Collection, ctx context.Context) *CommentRepository {
	return &CommentRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (repo *CommentRepository) CreateComment(comment *domain.Comment) error {
	inserted, err := repo.collection.InsertOne(repo.ctx, comment)
	if err != nil {
		return err
	}
	comment.Id = inserted.InsertedID.(primitive.ObjectID)
	return nil
}

func (repo *CommentRepository) GetComments(blogID primitive.ObjectID) ([]domain.Comment, error) {
	var comments []domain.Comment
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	cursor, err := repo.collection.Find(repo.ctx, filter)
	if err != nil {
		return nil, errors.New("error reading comments")
	}
	defer cursor.Close(repo.ctx)
	for cursor.Next(repo.ctx) {
		var comment domain.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	if len(comments) == 0 {
		return nil, errors.New("no comments found")
	}
	return comments, nil
}

func (repo *CommentRepository) DeleteComment(commentID primitive.ObjectID) (string, error) {
	filter := bson.D{{Key: "_id", Value: commentID}}
	findResult := repo.collection.FindOne(repo.ctx, filter)
	var comment domain.Comment

	if findResult.Decode(&comment) != nil {
		//fmt.Println("error deleting comment")
		return "", errors.New("error deleting comment")
	}

	result, err := repo.collection.DeleteOne(repo.ctx, filter)

	if err != nil {
		return "", errors.New("error deleting comment")
	}
	if result.DeletedCount == 0 {
		return "", errors.New("comment not found")
	}

	return comment.BlogID.Hex(), nil
}

func (repo *CommentRepository) DeleteCommentsOfBlog(blogID primitive.ObjectID) error {
	filter := bson.D{{Key: "blog_id", Value: blogID}}
	result, err := repo.collection.DeleteMany(repo.ctx, filter)
	if err != nil {
		return errors.New("error deleting comments")
	}
	if result.DeletedCount == 0 {
		return errors.New("no comments found")
	}
	return nil
}

func (repo *CommentRepository) UpdateComment(comment *domain.Comment) (string, error) {
	filter := bson.D{{Key: "_id", Value: comment.Id}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "body", Value: comment.Body},
	}}}

	var commnt domain.Comment

	findOneResult := repo.collection.FindOne(repo.ctx, filter)

	if findOneResult.Decode(&commnt) != nil {
		return "", errors.New("error updating comment")
	}

	updated, err := repo.collection.UpdateOne(repo.ctx, filter, update)

	if err != nil {
		return "", errors.New("error updating comment")
	}

	if updated.ModifiedCount == 0 {
		return "", errors.New("comment not found")
	}
	return commnt.BlogID.Hex(), nil
}

func (repo *CommentRepository) GetCommentByID(commentID primitive.ObjectID) (*domain.Comment, error) {

	filter := bson.D{{Key: "_id", Value: commentID}}

	var comment domain.Comment

	err := repo.collection.FindOne(repo.ctx, filter).Decode(&comment)

	if err != nil {
		return nil, errors.New("error getting comment" + err.Error())
	}

	return &comment, nil
}

func (repo *CommentRepository) DropDataBase() error {
	filter := bson.M{}
	_, err := repo.collection.DeleteMany(repo.ctx, filter)
	if err != nil {
		return err
	}
	return nil

}
