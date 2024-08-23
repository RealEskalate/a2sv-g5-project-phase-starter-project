package repository

import (
	"Blog_Starter/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogCommentRepository struct {
	DataBase          *mongo.Database
	commentCollection string
}

func NewCommentRepository(dataBase *mongo.Database, commentCollection string, ctx *context.Context) domain.CommentRepository {
	return &BlogCommentRepository{
		DataBase:          dataBase,
		commentCollection: commentCollection,
	}
}

// Create implements domain.CommentRepository.
func (bcr *BlogCommentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	comment.CommentID = primitive.NewObjectID()
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	_, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}

	var foundComment domain.Comment
	filter := bson.M{"_id": comment.CommentID}
	err = collection.FindOne(ctx, filter).Decode(&foundComment)
	return &foundComment, err
}

// Delete implements domain.CommentRepository.
func (bcr *BlogCommentRepository) Delete(ctx context.Context, commentID string) (*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(commentID)
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var deletedComment domain.Comment
	err = collection.FindOne(ctx, filter).Decode(&deletedComment)
	if err != nil {
		return nil, err
	}

	_, err = collection.DeleteOne(ctx, filter)
	return &deletedComment, err
}

// Update implements domain.CommentRepository.
func (bcr *BlogCommentRepository) Update(ctx context.Context, content string, commentID string) (*domain.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(commentID)
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "content", Value: content},
		},
	}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var foundComment domain.Comment
	err = collection.FindOne(ctx, filter).Decode(&foundComment)
	return &foundComment, err
}

// GetCommentByID implements domain.CommentRepository.
func (bcr *BlogCommentRepository) GetCommentByID(ctx context.Context, commentID string) (*domain.Comment, error) {

	objectID, err := primitive.ObjectIDFromHex(commentID)
	collection := bcr.DataBase.Collection(bcr.commentCollection)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var foundComment domain.Comment
	err = collection.FindOne(ctx, filter).Decode(&foundComment)
	return &foundComment, err
}

// GetComments implements domain.CommentRepository.
func (bcr *BlogCommentRepository) GetComments(ctx context.Context, userID string, blogID string) ([]*domain.Comment, error) {
	collection := bcr.DataBase.Collection(bcr.commentCollection)

	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "blog_id", Value: blogID},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx) {
		var comment domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
