package blog_repository

import (
	"blog-api/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *BlogRepository) UpdateBlog(ctx context.Context, blogID primitive.ObjectID, updateData domain.UpdateBlogRequest, userID primitive.ObjectID) error {
	blog, err := r.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	if blog.AuthorID != userID {
		return errors.New("unauthorized to update this blog")
	}

	update := bson.M{}
	if updateData.Title != "" {
		update["title"] = updateData.Title
	}
	if updateData.Author != "" {
		update["author"] = updateData.Author
	}
	if updateData.Content != "" {
		update["content"] = updateData.Content
	}
	if len(updateData.Tags) > 0 {
		update["tags"] = updateData.Tags
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	update["updated_at"] = time.Now()

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no blog found with the given ID")
	}

	return nil
}
