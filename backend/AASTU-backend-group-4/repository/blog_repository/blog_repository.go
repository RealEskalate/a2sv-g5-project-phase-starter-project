package blog_repository

import (
	domain "blog-api/domain/blog"
	"blog-api/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository struct {
	collection mongo.Collection
}

func NewBlogRepository(collection mongo.Collection) domain.BlogRepository {
	return &BlogRepository{
		collection: collection,
	}
}

func (br *BlogRepository) GetBlog(ctx context.Context, blogID primitive.ObjectID) (*domain.Blog, error) {
	var b domain.Blog
	filter := bson.M{"_id": blogID}

	err := br.collection.FindOne(ctx, filter).Decode(&b)
	return &b, err
}
