package blog_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog-api/domain"
)

func (r *BlogRepository) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {
	var blog domain.Blog
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}
