package blog_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog-api/domain"
)

func (r *BlogRepository) GetBlogByID(ctx context.Context, blogID primitive.ObjectID) (*domain.Blog, error) {
	var blog *domain.Blog
	filter := bson.M{"_id": blogID}
	err := r.collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return nil, err
	}
	blog.ViewCount += 1
	update := bson.M{
		"$set": bson.M{
			"view_count": blog.ViewCount,
		},
	}
	result, _ := r.collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		err = errors.New("couldn't add view count")
	}

	return blog, err
}
