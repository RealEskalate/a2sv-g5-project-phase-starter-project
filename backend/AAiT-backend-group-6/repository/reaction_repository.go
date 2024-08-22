package repository

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reactionRepository struct {
	db         mongo.Database
	collection string
}

func NewReactionRepository(db mongo.Database, collection string) domain.ReactionRepository {
	return &reactionRepository{
		db:         db,
		collection: collection,
	}
}

func (r *reactionRepository) LikeBlog(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID) error {
	collection := r.db.Collection(r.collection)

	// Check if the like already exists
	existingLike, err := r.GetLike(ctx, userID, blogID)
	if err != nil && err != mongo.ErrNoDocuments {
		println("Error while checking for existing like:", err.Error())
		return err
	}

	if existingLike != nil {
		// Update existing like to set IsLiked to true
		update := bson.M{"$set": bson.M{"is_liked": true}}
		_, err := collection.UpdateOne(ctx, bson.M{"user_id": userID, "blog_id": blogID}, update)
		return err
	}

	// Create a new like
	var like domain.Reaction
	like.BlogID = blogID
	like.UserID = userID
	like.IsLiked = true

	// Insert the new like document into the collection
	_, err = collection.InsertOne(ctx, &like) // Pass a pointer to the like struct
	return err
}

func (r *reactionRepository) UnLikeBlog(ctx context.Context, userID, blogID primitive.ObjectID) error {
	collection := r.db.Collection(r.collection)

	existingLike, err := r.GetLike(ctx, userID, blogID)
	if err != nil {
		return err
	}

	if existingLike != nil {
		// Update existing like to false
		update := bson.M{"$set": bson.M{"is_liked": false}}
		_, err := collection.UpdateOne(ctx, bson.M{"user_id": userID, "blog_id": blogID}, update)
		return err
	}
	var unLike domain.Reaction

	unLike.BlogID = blogID
	unLike.UserID = userID
	unLike.IsLiked = false
	_, err = collection.InsertOne(ctx, &unLike)

	return err
}

func (r *reactionRepository) DeleteLike(ctx context.Context, userID, blogID primitive.ObjectID) error {
	collection := r.db.Collection(r.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"user_id": userID, "blog_id": blogID})
	return err
}

func (r *reactionRepository) GetLike(ctx context.Context, userID primitive.ObjectID, blogID primitive.ObjectID) (*domain.Reaction, error) {
	collection := r.db.Collection(r.collection)

	var like domain.Reaction
	err := collection.FindOne(ctx, bson.M{"user_id": userID, "blog_id": blogID}).Decode(&like)
	if err != nil {
		return nil, err
	}
	return &like, nil
}
