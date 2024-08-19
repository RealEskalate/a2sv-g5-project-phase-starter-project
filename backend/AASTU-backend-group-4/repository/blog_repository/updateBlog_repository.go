package blog_repository

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (br *BlogRepository) UpdateBlog(ctx context.Context, blogID primitive.ObjectID, updatedBlog *domain.Blog) error {

	filter := bson.M{"_id": blogID}

	update := bson.M{
		"$set": bson.M{
			"title":      updatedBlog.Title,
			"content":    updatedBlog.Content,
			"tags":       updatedBlog.Tags,
			"updated_at": updatedBlog.UpdatedAt,
		},
	}

	_, err := br.collection.UpdateOne(ctx, filter, update)
	return err
}
