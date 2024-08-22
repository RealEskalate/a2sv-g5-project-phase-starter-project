package repository

import (
	"context"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogCommentRepository struct {
	BlogCommentCollection interfaces.Collection
}

func NewBlogCommentRepository(db interfaces.Database) interfaces.BlogCommentRepository {
	return &BlogCommentRepository{
		BlogCommentCollection: db.Collection("comments"),
	}
}

func (r *BlogCommentRepository) AddComment(ctx context.Context, comment models.Comment) *models.ErrorResponse {
	comment.ID = primitive.NewObjectID().Hex()
	comment.CreatedAt = primitive.NewObjectID().Timestamp()

	_, err := r.BlogCommentCollection.InsertOne(ctx, comment)
	if err != nil {
		return models.InternalServerError("Failed to create comment")
	}

	return models.Nil()
}

func (r *BlogCommentRepository) GetComments(ctx context.Context, blogID string) ([]models.Comment, *models.ErrorResponse) {

	var comments []models.Comment
	cursor, err := r.BlogCommentCollection.Find(ctx, map[string]string{"blog_id": blogID})

	if err != nil {
		return nil, models.InternalServerError("Failed to retrieve comments")
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &comments); err != nil {
		return nil, models.InternalServerError("Failed to retrieve comments")
	}

	return comments, models.Nil()
}

func (r *BlogCommentRepository) UpdateComment(ctx context.Context, commentID string, comment dtos.CommentUpdateRequest) *models.ErrorResponse {
	filter := bson.M{"_id": commentID}
	update := bson.M{"$set": comment}

	_, err := r.BlogCommentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.InternalServerError("Failed to update comment")
	}

	return models.Nil()
}

func (r *BlogCommentRepository) DeleteComment(ctx context.Context, commentID string) *models.ErrorResponse {

	filter := bson.M{"_id": commentID}
	_, err := r.BlogCommentCollection.DeleteOne(ctx, filter)
	if err != nil {
		return models.InternalServerError("Failed to delete comment")
	}

	return nil
}

func (r *BlogCommentRepository) GetComment(ctx context.Context, commentID string) (*models.Comment, *models.ErrorResponse) {
	var comment models.Comment
	filter := bson.M{"_id": commentID}
	err := r.BlogCommentCollection.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.NotFound("Comment not found")
		}
		return nil, models.InternalServerError("Failed to retrieve comment")
	}

	return &comment, models.Nil()
}

func (r *BlogCommentRepository) DeleteComments(ctx context.Context, blogID string) *models.ErrorResponse {
	filter := bson.M{"blog_id": blogID}
	_, err := r.BlogCommentCollection.DeleteMany(ctx, filter)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return models.Nil()
		}

		return models.InternalServerError("Failed to delete comments")
	}

	return nil
}
