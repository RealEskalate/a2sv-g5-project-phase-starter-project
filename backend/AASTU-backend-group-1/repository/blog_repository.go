package repository

import (
	"blogs/domain"
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCachedUser(cache domain.Cache, key string) (*domain.User, error) {
	cachedUser, err := cache.GetCache(key)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = bson.UnmarshalExtJSON([]byte(cachedUser), true, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type BlogRepository struct {
	blogCollection    *mongo.Collection
	viewCollection    *mongo.Collection
	likeCollection    *mongo.Collection
	commentCollection *mongo.Collection
	cache             domain.Cache
}

func NewBlogRepository(database *mongo.Database, cache domain.Cache) *BlogRepository {
    return &BlogRepository{
        blogCollection:    database.Collection("blogs"),
        viewCollection:    database.Collection("views"),
        likeCollection:    database.Collection("likes"),
        commentCollection: database.Collection("comments"),
        cache:             cache,
    }
}


// InsertBlog implements domain.BlogRepository.
func (b *BlogRepository) InsertBlog(blog *domain.Blog) (*domain.Blog, error) {

	newblog, err := b.blogCollection.InsertOne(context.Background(), blog)
	if err != nil {
		return nil, err
	}
	blog.ID = newblog.InsertedID.(primitive.ObjectID)
	return blog, nil
}

// GetBlogByID implements domain.BlogRepository.
func (b *BlogRepository) GetBlogByID(id string) (*domain.Blog, error) {
	cacheKey := fmt.Sprintf("blog:%s", id)
	getCachedUser(b.cache, cacheKey)

	blogid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog domain.Blog
	err = b.blogCollection.FindOne(context.Background(), bson.M{"_id": blogid}).Decode(&blog)
	if err != nil {
		return nil, err
	}

	blogJson, err := bson.MarshalExtJSON(blog, true, true)
	if err == nil {
		_ = b.cache.SetCache(cacheKey, string(blogJson))
	}

	return &blog, nil
}

// UpdateBlogByID implements domain.BlogRepository.
func (b *BlogRepository) UpdateBlogByID(id string, blog *domain.Blog) error {
	blogid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$set": blog})

	return err
}

// DeleteBlogByID implements domain.BlogRepository.
func (b *BlogRepository) DeleteBlogByID(id string) error {
	blogid, err := primitive.ObjectIDFromHex(id)
	// Delete blog from cache

	err = b.cache.DeleteCache(fmt.Sprintf("blog:%s", id))
	if err != nil {
		return err
	}

	

	
	_,err = b.commentCollection.DeleteMany(context.Background(), bson.M{"blogid": blogid})


	if err != nil {
		return err
	}

	_,err = b.likeCollection.DeleteMany(context.Background(), bson.M{"blogid": blogid})


	if err != nil {
		return err
	}

	_,err = b.viewCollection.DeleteMany(context.Background(), bson.M{"blogid": blogid})


	if err != nil {
		return err
	}


	_, err = b.blogCollection.DeleteOne(context.Background(), bson.M{"_id": blogid})
	return err
}

func (b *BlogRepository) SearchBlog(title, author string, tags []string) ([]*domain.Blog, error) {
	//implement caching for search
	cachekey := fmt.Sprintf("search:%s:%s:%s", title, author, strings.Join(tags, ","))
	cachedPost, err := b.cache.GetCache(cachekey)

	if err == nil && cachedPost != "" {
		var blogs []*domain.Blog
		err = bson.UnmarshalExtJSON([]byte(cachedPost), true, &blogs)
		if err != nil {
			return nil, err
		}
		return blogs, nil
	}

	blogs := []*domain.Blog{}
	tagFilters := make([]bson.M, len(tags))
	for i, tag := range tags {
		tagFilters[i] = bson.M{"tags": bson.M{"$regex": tag, "$options": "i"}}
	}

	otherFilters := []bson.M{}
	if title != "" {
		otherFilters = append(otherFilters, bson.M{"title": bson.M{"$regex": title, "$options": "i"}})
	}

	if author != "" {
		otherFilters = append(otherFilters, bson.M{"author": bson.M{"$regex": author, "$options": "i"}})
	}

	filter := bson.M{
		"$or": append(tagFilters, otherFilters...),
	}

	cursor, err := b.blogCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var blog domain.Blog
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	blogsJSON, err := bson.MarshalExtJSON(blogs, true, true)
	if err == nil {
		_ = b.cache.SetCache(cachekey, string(blogsJSON))
	}

	return blogs, nil
}

// FilterBlog implements domain.BlogRepository.
func (b *BlogRepository) FilterBlog(tags []string, dateFrom time.Time, dateTo time.Time) ([]*domain.Blog, error) {
	//implement caching for filter
	cacheKey := fmt.Sprintf("filter:%s:%s:%s", strings.Join(tags, ","), dateFrom, dateTo)
	cachedPost, err := b.cache.GetCache(cacheKey)

	if err == nil && cachedPost != "" {
		var blogs []*domain.Blog
		err = bson.UnmarshalExtJSON([]byte(cachedPost), true, &blogs)
		if err != nil {
			return nil, err
		}
		return blogs, nil
	}

	blogs := []*domain.Blog{}
	filter := bson.M{
		"tags": bson.M{"$in": tags},
		"date": bson.M{"$gte": dateFrom, "$lte": dateTo},
	}
	cursor, err := b.blogCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err

	}
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	blogJSON, err := bson.MarshalExtJSON(blogs, true, true)
	if err == nil {
		// Store in cache if marshalling is successful
		_ = b.cache.SetCache(cacheKey, string(blogJSON))
	}

	return blogs, nil
}

func (b *BlogRepository) GetBlogsByPopularity(page, limit int, reverse bool) ([]*domain.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:popularity:page=%d:limit=%d:reverse=%v", page, limit, reverse)

	// Check cache first
	cachedBlogs, err := b.cache.GetCache(cacheKey)
	if err == nil && cachedBlogs != "" {
		var blogs []*domain.Blog
		err = bson.UnmarshalExtJSON([]byte(cachedBlogs), true, &blogs)
		if err == nil {
			return blogs, nil
		}
	}

	var blogs []*domain.Blog

	// Calculate how many documents to skip
	skip := (page - 1) * limit
	sortOrder := -1 // Default sort order is descending
	if reverse {
		sortOrder = 1 // If reverse is true, sort in ascending order
	}

	// MongoDB aggregation pipeline with pagination
	pipeline := mongo.Pipeline{
		// Add popularityScore field based on weights for views, likes, comments
		bson.D{
			{Key: "$addFields", Value: bson.D{
				{Key: "popularityScore", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$multiply", Value: bson.A{"$views_count", 0.5}}},
						bson.D{{Key: "$multiply", Value: bson.A{"$likes_count", 1}}},
						bson.D{{Key: "$multiply", Value: bson.A{"$comments_count", 2}}},
					}},
				}},
			}},
		},
		// Sort by popularity score
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "popularityScore", Value: sortOrder},
			}},
		},
		// Add pagination stage
		bson.D{
			{Key: "$skip", Value: skip}, // Skip N documents for pagination
		},
		bson.D{
			{Key: "$limit", Value: limit}, // Limit the number of documents per page
		},
	}

	// Execute the aggregation pipeline
	cursor, err := b.blogCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO()) // Ensure cursor is closed

	// Decode the result into blogs
	if err := cursor.All(context.TODO(), &blogs); err != nil {
		return nil, err
	}

	// Cache the result
	blogJSON, err := bson.MarshalExtJSON(blogs, true, true)
	if err == nil {
		// Store in cache if marshalling is successful
		_ = b.cache.SetCache(cacheKey, string(blogJSON))
	}

	return blogs, nil
}

func (b *BlogRepository) GetBlogsByRecent(page, limit int, reverse bool) ([]*domain.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:recent:page=%d:limit=%d:reverse=%v", page, limit, reverse)
	cachedPost, err := b.cache.GetCache(cacheKey)

	if err == nil && cachedPost != "" {
		var blogs []*domain.Blog
		err = bson.UnmarshalExtJSON([]byte(cachedPost), true, &blogs)
		if err == nil {
			return blogs, nil
		}
	}

	var blogs []*domain.Blog

	// Calculate the number of documents to skip for pagination
	skip := (page - 1) * limit
	sortOrder := -1 // Default sort order is descending
	if reverse {
		sortOrder = 1 // If reverse is true, sort in ascending order
	}

	// MongoDB query to find all blogs, sort them by CreatedAt in descending order, and apply pagination
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: sortOrder}}). // Sort by CreatedAt in descending order
		SetSkip(int64(skip)).                                   // Skip N documents for pagination
		SetLimit(int64(limit))                                  // Limit the number of documents per page

	cursor, err := b.blogCollection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO()) // Ensure the cursor is closed after usage

	// Decode all the blogs into the blogs slice
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	// Check if the cursor encountered any errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Cache the result
	blogJSON, err := bson.MarshalExtJSON(blogs, true, true)
	if err == nil {
		// Store in cache if marshalling is successful
		_ = b.cache.SetCache(cacheKey, string(blogJSON))
	}

	return blogs, nil
}

// AddView implements domain.BlogRepository.
func (b *BlogRepository) AddView(view []*domain.View) error {
	var views []interface{}
	for _, v := range view {
		views = append(views, v)
	}
	_, err := b.viewCollection.InsertMany(context.Background(), views)
	return err

}

// GetLikebyAuthorAndBlogID implements domain.BlogRepository.
func (b *BlogRepository) GetLikebyAuthorAndBlogID(blogID string, author string) (*domain.Like, error) {
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	var like domain.Like
	err = b.likeCollection.FindOne(context.Background(), bson.M{"blogid": blogid, "user": author}).Decode(&like)
	if err != nil {
		return nil, err
	}

	return &like, nil
}

func (b *BlogRepository) UpdateLike(like *domain.Like) error {
	_, err := b.likeCollection.UpdateOne(context.Background(), bson.M{"blogid": like.BlogID, "user": like.User}, bson.M{"$set": like})
	return err
}

// AddLike implements domain.BlogRepository.
func (b *BlogRepository) AddLike(like *domain.Like) error {
	_, err := b.likeCollection.InsertOne(context.Background(), like)
	return err

}

// RemoveLike removes a like by blogID and author from the database.
func (b *BlogRepository) RemoveLike(blogID string, author string) error {
	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	filter := bson.M{"blogid": id, "user": author}

	_, err = b.likeCollection.DeleteOne(context.Background(), filter)
	return err
}

// AddComment implements domain.BlogRepository.
func (b *BlogRepository) AddComment(comment *domain.Comment) error {
	_, err := b.commentCollection.InsertOne(context.Background(), comment)
	return err
}

//DeleteComment implements domain.BlogRepository.
func (b *BlogRepository) DeleteComment(commentID string) error {
	commentid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	_, err = b.commentCollection.DeleteOne(context.Background(), bson.M{"_id": commentid})
	return err
}
//Decrement Blog Views
func (b *BlogRepository) DecrementBlogViews(blogID string) error {
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$inc": bson.M{"views_count": -1}})
	return err

}

//Get Comment By ID
func (b *BlogRepository) GetCommentByID(commentID string) (*domain.Comment, error) {
	commentid, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, err
	}

	var comment domain.Comment
	err = b.commentCollection.FindOne(context.Background(), bson.M{"_id": commentid}).Decode(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

//Decrement Blog Comments
func (b *BlogRepository) DecrementBlogComments(blogID string) error {
	blogid, err := primitive.ObjectIDFromHex(blogID)

	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$inc": bson.M{"comments_count": -1}})
	return err
}

//GET Single Blog's Comments

func (b *BlogRepository) GetBlogComments(blogID string) ([]*domain.Comment, error) {
	cacheKey := fmt.Sprintf("comments:%s", blogID)
	cachedPost, err := b.cache.GetCache(cacheKey)
	if err == nil && cachedPost != "" {
		var comments []*domain.Comment
		err = bson.UnmarshalExtJSON([]byte(cachedPost), true, &comments)
		if err == nil {
			return comments, nil
		}
	}

	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	var comments []*domain.Comment
	cursor, err := b.commentCollection.Find(context.Background(), bson.M{"blogid": blogid})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var comment domain.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	commentsJSON, err := bson.MarshalExtJSON(comments, true, true)
	if err == nil {
		_ = b.cache.SetCache(cacheKey, string(commentsJSON))
	}

	return comments, nil
}

// Get likes for specific blog

func (b *BlogRepository) GetBlogLikes(blogID string) ([]*domain.Like, error) {
	cacheKey := fmt.Sprintf("likes:%s", blogID)
	cachedLike, err := b.cache.GetCache(cacheKey)
	if err == nil && cachedLike != "" {
		var likes []*domain.Like
		err = bson.UnmarshalExtJSON([]byte(cachedLike), true, &likes)
		if err == nil {
			return likes, nil
		}
	}
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	var likes []*domain.Like
	cursor, err := b.likeCollection.Find(context.Background(), bson.M{"blogid": blogid})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var like domain.Like
		err := cursor.Decode(&like)
		if err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}

	likesJSON, err := bson.MarshalExtJSON(likes, true, true)
	if err == nil {
		_ = b.cache.SetCache(cacheKey, string(likesJSON))
	}
	return likes, nil
}

func (b *BlogRepository) IncrmentBlogViews(blogID string) error {
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$inc": bson.M{"views_count": 1}})
	return err
}

func (b *BlogRepository) IncrmentBlogLikes(blogID string) error {
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$inc": bson.M{"likes_count": 1}})
	return err
}

func (b *BlogRepository) IncrmentBlogComments(blogID string) error {
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$inc": bson.M{"comments_count": 1}})
	return err
}


func (b *BlogRepository) DecrementBlogLikes(blogID string) error {
	blogid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	_, err = b.blogCollection.UpdateOne(context.Background(), bson.M{"_id": blogid}, bson.M{"$inc": bson.M{"likes_count": -1}})
	return err
}

func (b *BlogRepository) GetTotalBlogs() (int, error) {
	total, err := b.blogCollection.CountDocuments(context.Background(), bson.M{})
	return int(total), err
}
