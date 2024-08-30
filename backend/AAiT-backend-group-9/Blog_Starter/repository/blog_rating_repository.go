package repository

import (
	"Blog_Starter/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRatingRepository struct {
	DataBase         *mongo.Database
	ratingCollection string
}

func NewBlogRatingRepository(dataBase *mongo.Database, ratingCollection string) domain.BlogRatingRepository {
	return &BlogRatingRepository{
		DataBase:         dataBase,
		ratingCollection: ratingCollection,
	}
}

// GetRatingByBlogID implements domain.BlogRatingRepository.
func (br *BlogRatingRepository) GetRatingByBlogID(ctx context.Context, blogID string) ([]*domain.BlogRating, error) {
	collection := br.DataBase.Collection(br.ratingCollection)
	filter := bson.M{"blog_id": blogID}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var blogRatings []*domain.BlogRating
	for cur.Next(ctx) {
		var singleBlog domain.BlogRating
		err := cur.Decode(&singleBlog)
		if err != nil {
			return nil, err
		}

		blogRatings = append(blogRatings, &singleBlog)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return blogRatings, nil
}

// GetRatingByID implements domain.BlogRatingRepository.
func (br *BlogRatingRepository) GetRatingByID(ctx context.Context, ratingID string) (*domain.BlogRating, error) {
	objectID, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		return nil, err
	}
	collection := br.DataBase.Collection(br.ratingCollection)
	var foundRating domain.BlogRating
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&foundRating)
	if err != nil {
		return nil, err
	}

	return &foundRating, nil
}

// InsertRating implements domain.BlogRatingRepository.
func (br *BlogRatingRepository) InsertRating(ctx context.Context, rating *domain.BlogRating) (*domain.BlogRating, error) {
	rating.RatingID = primitive.NewObjectID()
	collection := br.DataBase.Collection(br.ratingCollection)
	_, err := collection.InsertOne(ctx, rating)
	if err != nil {
		return nil, err
	}

	var insertedRating domain.BlogRating
	filter := bson.M{"_id": rating.RatingID}
	err = collection.FindOne(ctx, filter).Decode(&insertedRating)
	if err != nil {
		return nil, err
	}

	return &insertedRating, nil
}

// UpdateRating implements domain.BlogRatingRepository.
func (br *BlogRatingRepository) UpdateRating(ctx context.Context, rating int, ratingID string) (*domain.BlogRating, int, error) {
	objectID, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		return nil, 0, err
	}
	collection := br.DataBase.Collection(br.ratingCollection)
	filter := bson.M{"_id": objectID}
	var prevRating domain.BlogRating
	err = collection.FindOne(ctx, filter).Decode(&prevRating)
	if err != nil {
		return nil, 0, err
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "rating", Value: rating},
		{Key: "updatetimestamp", Value: time.Now()},
	}}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, 0, err
	}

	var curRating domain.BlogRating
	err = collection.FindOne(ctx, filter).Decode(&curRating)
	if err != nil {
		return nil, 0, err
	}

	return &curRating, prevRating.Rating, nil
}

func (br *BlogRatingRepository) DeleteRating(ctx context.Context, ratingID string) (*domain.BlogRating, error) {
	objectID, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		return nil, err
	}
	collection := br.DataBase.Collection(br.ratingCollection)
	filter := bson.M{"_id": objectID}
	var deletedRating domain.BlogRating
	err = collection.FindOne(ctx, filter).Decode(&deletedRating)
	if err != nil {
		return nil, err
	}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &deletedRating, nil
}

func (br *BlogRatingRepository) DeleteRatingByBlogID(ctx context.Context, blogID string) error {

	collection := br.DataBase.Collection(br.ratingCollection)
	_, err := collection.DeleteMany(ctx, bson.M{"blog_id": blogID})
	return err
}
