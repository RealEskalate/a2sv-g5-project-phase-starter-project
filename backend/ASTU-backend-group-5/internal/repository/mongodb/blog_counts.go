package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoBlogRepository) IncrementBlogViewCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"views_count": 1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) IncrementBlogLikeCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"likes_count": 1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) IncrementBlogCommentCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"comments_count": 1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) DecrementBlogViewCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"views_count": -1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) DecrementBlogLikeCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"likes_count": -1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoBlogRepository) DecrementBlogCommentCount(ctx context.Context, blogID string) error {
	// Create an update filter
	blogIdObj, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": blogIdObj}
	// Create an update operation
	update := bson.M{"$inc": bson.M{"comments_count": -1}}

	// Execute the update
	_, err := r.blogsCollection.UpdateOne(ctx, filter, update)

	return err
}
