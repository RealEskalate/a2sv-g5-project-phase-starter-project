package repository

import (
	"Blog_Starter/domain"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository struct {
	db             *mongo.Database
	blogCollection string
}

// DeleteRating implements domain.BlogRepository.
// InsertRating implements domain.BlogRepository

func NewBlogRepository(db *mongo.Database, blogCollection string, c *context.Context) domain.BlogRepository {
	return &BlogRepository{
		db:             db,
		blogCollection: blogCollection,
	}
}

func (r *BlogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	_, err := collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) GetBlogByID(ctx context.Context, blogID string) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	// bson filtretion
	filter := bson.M{"_id": blogID}
	var blog domain.Blog
	err := collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("blog not found")
		} else {
			return nil, err
		}
	}
	return &blog, nil
}

func (r *BlogRepository) GetAllBlog(ctx context.Context) ([]*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	var blogs []*domain.Blog
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepository) UpdateBlog(ctx context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	filter := bson.M{"_id": blogID}
	update := bson.M{"$set": blog}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &domain.Blog{}, nil
}

func (r *BlogRepository) DeleteBlog(ctx context.Context, blogID string) error {
	// implementation
	collection := r.db.Collection(r.blogCollection)
	filter := bson.M{"_id": blogID}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// FilterBlogs implements domain.BlogRepository.
func (bf *BlogRepository) FilterBlogs(ctx context.Context, blogRequest *domain.BlogFilterRequest) ([]*domain.Blog, error) {
	collection := bf.db.Collection(bf.blogCollection)
	filter := bson.M{}
	if blogRequest.LikeLowerRange > 0 {
		filter["like_count"] = bson.M{
			"$gt": blogRequest.LikeLowerRange,
		}
	}

	if blogRequest.ViewLowerRange > 0 {
		filter["view_count"] = bson.M{
			"$gt": blogRequest.ViewLowerRange,
		}
	}

	if blogRequest.Date != nil {
		filter["createtimestamp"] = blogRequest.Date
	}

	if blogRequest.Tags != nil {
		filter["tags"] = bson.M{
			"$in": blogRequest.Tags,
		}
	}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var blogResponse []*domain.Blog
	for cur.Next(context.TODO()) {
		var singleResponse domain.Blog
		err := cur.Decode(&singleResponse)
		if err != nil {
			return nil, err
		}

		blogResponse = append(blogResponse, &singleResponse)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return blogResponse, nil
}

// IncrementViewCount implements domain.BlogRepository.
func (r *BlogRepository) IncrementViewCount(ctx context.Context, blogID string) error {
	collection := r.db.Collection(r.blogCollection)
	objectID, _ := primitive.ObjectIDFromHex(blogID)
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "view_count", Value: 1},
		}},
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// SearchBlogs implements domain.BlogRepository.
func (sr *BlogRepository) SearchBlogs(ctx context.Context, searchRequest *domain.BlogSearchRequest) ([]*domain.Blog, error) {
	collection := sr.db.Collection(sr.blogCollection)
	filter := bson.M{}
	if searchRequest.Title != "" && searchRequest.Author != "" {
		filter = bson.M{"author": searchRequest.Author, "title": searchRequest.Title}
	}
	if searchRequest.Title == "" && searchRequest.Author != "" {
		filter = bson.M{"author": searchRequest.Author}
	}
	if searchRequest.Title != "" && searchRequest.Author == "" {
		filter = bson.M{"title": searchRequest.Title}
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var searchedBlogs []*domain.Blog
	for cur.Next(context.TODO()) {
		var singleBlog domain.Blog
		err := cur.Decode(&singleBlog)
		if err != nil {
			return nil, err
		}

		searchedBlogs = append(searchedBlogs, &singleBlog)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return searchedBlogs, nil
}



func (sr *BlogRepository) InsertRating(ctx context.Context, insertedRating *domain.BlogRating) error {
	objectID, err := primitive.ObjectIDFromHex(insertedRating.BlogID)
	if err != nil {
		return err
	}
	collection := sr.db.Collection(sr.blogCollection)
	filter := bson.M{"_id": objectID}

	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "rating_count", Value: 1},
			{Key: "total_rating", Value: insertedRating.Rating},
		}},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	var result struct {
		TotalRating int `bson:"total_rating"`
		RatingCount int `bson:"rating_count"`
	}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return err
	}

	updatedAverageRating := float64(result.TotalRating) / float64(result.RatingCount)
	_, err = collection.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "average_rating", Value: updatedAverageRating},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

// UpdateRating implements domain.BlogRepository.
func (sr *BlogRepository) UpdateRating(ctx context.Context, updatedRating *domain.BlogRating, prevRating int) error {

	objectID, err := primitive.ObjectIDFromHex(updatedRating.BlogID)
	if err != nil {
		return err
	}
	collection := sr.db.Collection(sr.blogCollection)
	filter := bson.M{"_id": objectID}

	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "total_rating", Value: updatedRating.Rating - prevRating},
		}},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	var result struct {
		TotalRating int `bson:"total_rating"`
		RatingCount int `bson:"rating_count"`
	}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return err
	}

	updatedAverageRating := float64(result.TotalRating) / float64(result.RatingCount)
	_, err = collection.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "average_rating", Value: updatedAverageRating},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

func (sr *BlogRepository) DeleteRating(ctx context.Context, deletedRating *domain.BlogRating) error {

	objectID, err := primitive.ObjectIDFromHex(deletedRating.BlogID)
	if err != nil {
		return err
	}
	collection := sr.db.Collection(sr.blogCollection)
	filter := bson.M{"_id" : objectID}
	update := bson.D{{Key: "$inc", Value: bson.D{
		{Key: "total_rating", Value: -deletedRating.Rating},
		{Key: "rating_count", Value: -1},
	}}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	var result struct {
		TotalRating int `bson:"total_rating"`
		RatingCount int `bson:"rating_count"`
	}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return err
	}

	updatedAverageRating := float64(result.TotalRating) / float64(result.RatingCount)
	_, err = collection.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "average_rating", Value: updatedAverageRating},
		}},
	})

	if err != nil {
		return err
	}

	return nil
}

// UpdateCommentCount implements domain.BlogRepository.
func (sr *BlogRepository) UpdateCommentCount(ctx context.Context, blogID string, increment bool) error {
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	collection := sr.db.Collection(sr.blogCollection)
	filter := bson.M{"_id" : objectID}
	update := bson.D{{}}
	if increment {
		update = bson.D{{
			Key : "$inc", Value: bson.D{
				{Key : "comment_count", Value : 1},
			},
		}}
	} else {
		update = bson.D{{
			Key : "$inc", Value: bson.D{
				{Key : "comment_count", Value : -1},
			},
		}}
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}

// UpdateLikeCount implements domain.BlogRepository.
func (r *BlogRepository) UpdateLikeCount(c context.Context, blogID string, increment bool) error {
	panic("unimplemented")
}
