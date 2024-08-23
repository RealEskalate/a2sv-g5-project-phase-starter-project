package commentrepository

import (
	domain "AAiT-backend-group-2/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	collection *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) domain.CommentRepository {
	return &commentRepository{
		collection: db.Collection("comments"),
	}
}

func (r *commentRepository) FindAll(ctx context.Context) ([]domain.Comment, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []domain.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) FindByID(ctx context.Context, id string) (*domain.Comment, error) {
	var comment domain.Comment
	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) FindByBlogID(ctx context.Context, blogID string) ([]domain.Comment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"blog_id": blogID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []domain.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) Save(ctx context.Context, comment *domain.Comment) error {
	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

func (r *commentRepository) Update(ctx context.Context, comment *domain.Comment) error {
	filter := bson.M{"_id": comment.ID}
	update := bson.M{
		"$set": bson.M{
			"content":    comment.Content,
			"author":     comment.Author,
			"updated_at": comment.CreatedAt,
		},
	}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *commentRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
