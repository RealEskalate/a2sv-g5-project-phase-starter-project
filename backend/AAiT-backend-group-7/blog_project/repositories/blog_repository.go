package repositories

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"blog_project/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	collection *mongo.Collection
	cache      domain.Cache
}

func NewBlogRepository(collection *mongo.Collection, cache domain.Cache) domain.IBlogRepository {
	return &blogRepository{
		collection: collection,
		cache:      cache,
	}
}

func (blogRepo *blogRepository) GetAllBlogs(ctx context.Context) ([]domain.Blog, error) {
	cacheKey := "blogs:all"
	var blogs []domain.Blog

	// Check cache
	err := blogRepo.cache.Get(ctx, cacheKey, &blogs)
	if err == nil {
		return blogs, nil
	}

	// Fetch from DB if cache is empty
	cursor, err := blogRepo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Cache the result if blogs are found
	if len(blogs) > 0 {
		blogRepo.cache.Set(ctx, cacheKey, blogs, 1*time.Hour)
	}

	return blogs, nil
}

func (blogRepo *blogRepository) GetBlogsByPage(ctx context.Context, offset, limit int) ([]domain.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:page:%d:%d", offset, limit)
	var blogs []domain.Blog

	// Check cache
	err := blogRepo.cache.Get(ctx, cacheKey, &blogs)
	if err == nil {
		return blogs, nil
	}

	// Fetch from DB if cache is empty
	findOptions := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	cursor, err := blogRepo.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Cache the result if blogs are found
	if len(blogs) > 0 {
		blogRepo.cache.Set(ctx, cacheKey, blogs, 1*time.Hour)
	}

	return blogs, nil
}

func (blogRepo *blogRepository) GetBlogByID(ctx context.Context, id int) (domain.Blog, error) {
	cacheKey := fmt.Sprintf("blog:%d", id)
	var blog domain.Blog

	// Check cache
	err := blogRepo.cache.Get(ctx, cacheKey, &blog)
	if err == nil {
		return blog, nil
	}

	// Fetch from DB if cache is empty
	err = blogRepo.collection.FindOne(ctx, bson.M{"id": id}).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}
		return domain.Blog{}, err
	}

	// Cache the result
	blogRepo.cache.Set(ctx, cacheKey, blog, 1*time.Hour)
	return blog, nil
}

func (blogRepo *blogRepository) CreateBlog(ctx context.Context, blog domain.Blog) (domain.Blog, error) {
	blog.Date = time.Now() // Set the current time for the blog's creation date
	_, err := blogRepo.collection.InsertOne(ctx, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	cacheKey := fmt.Sprintf("blog:%d", blog.ID)

	// Cache the new blog
	blogRepo.cache.Set(ctx, cacheKey, blog, 1*time.Hour)

	// Invalidate related caches
	blogRepo.cache.Del(ctx, "blogs:all")
	blogRepo.invalidatePaginationAndSearchCaches(ctx)

	return blog, nil
}

func (blogRepo *blogRepository) UpdateBlog(ctx context.Context, id int, blog domain.Blog) (domain.Blog, error) {
	var updatedBlog domain.Blog

	// Update in DB
	result := blogRepo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": blog},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	err := result.Decode(&updatedBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Blog{}, errors.New("blog not found")
		}
		return domain.Blog{}, err
	}

	cacheKey := fmt.Sprintf("blog:%d", id)

	// Update the cache for the updated blog
	blogRepo.cache.Set(ctx, cacheKey, updatedBlog, 1*time.Hour)

	// Invalidate related caches
	blogRepo.cache.Del(ctx, "blogs:all")
	blogRepo.invalidatePaginationAndSearchCaches(ctx)

	return updatedBlog, nil
}

func (blogRepo *blogRepository) DeleteBlog(ctx context.Context, id int) error {
	result, err := blogRepo.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("blog not found")
	}

	cacheKey := fmt.Sprintf("blog:%d", id)

	// Invalidate the cache for the deleted blog
	blogRepo.cache.Del(ctx, cacheKey)

	// Invalidate related caches
	blogRepo.cache.Del(ctx, "blogs:all")
	blogRepo.invalidatePaginationAndSearchCaches(ctx)

	return nil
}

func (blogRepo *blogRepository) SearchByTitle(ctx context.Context, title string) ([]domain.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:search:title:%s", title)
	var blogs []domain.Blog

	// Check cache
	err := blogRepo.cache.Get(ctx, cacheKey, &blogs)
	if err == nil {
		return blogs, nil
	}

	// Fetch from DB if cache is empty
	filter := bson.M{"title": bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(title) + "$", Options: "i"}}}
	cursor, err := blogRepo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Cache the result if blogs are found
	if len(blogs) > 0 {
		blogRepo.cache.Set(ctx, cacheKey, blogs, 1*time.Hour)
	}

	return blogs, nil
}

func (blogRepo *blogRepository) SearchByTags(ctx context.Context, tags []string) ([]domain.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:search:tags:%v", tags)
	var blogs []domain.Blog

	// Check cache
	err := blogRepo.cache.Get(ctx, cacheKey, &blogs)
	if err == nil {
		return blogs, nil
	}

	// Fetch from DB if cache is empty
	filter := bson.M{"tags": bson.M{"$in": tags}}
	cursor, err := blogRepo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Cache the result if blogs are found
	if len(blogs) > 0 {
		blogRepo.cache.Set(ctx, cacheKey, blogs, 1*time.Hour)
	}

	return blogs, nil
}

func (blogRepo *blogRepository) SearchByAuthor(ctx context.Context, author string) ([]domain.Blog, error) {
	cacheKey := fmt.Sprintf("blogs:search:author:%s", author)
	var blogs []domain.Blog

	// Check cache
	err := blogRepo.cache.Get(ctx, cacheKey, &blogs)
	if err == nil {
		return blogs, nil
	}

	// Fetch from DB if cache is empty
	cursor, err := blogRepo.collection.Find(ctx, bson.M{"author": author})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Cache the result if blogs are found
	if len(blogs) > 0 {
		blogRepo.cache.Set(ctx, cacheKey, blogs, 1*time.Hour)
	}

	return blogs, nil
}

// Helper function to invalidate pagination and search caches
func (blogRepo *blogRepository) invalidatePaginationAndSearchCaches(ctx context.Context) {
	// Invalidate all pagination caches
	blogRepo.cache.DelByPattern(ctx, "blogs:page:*")

	// Invalidate search caches
	blogRepo.cache.DelByPattern(ctx, "blogs:search:*")
}

func (blogRepo *blogRepository) UpdateAuthorName(ctx context.Context, oldName, newName string) error {
	println("Updating author name from", oldName, "to", newName)
	filter := bson.M{"author": oldName}
	update := bson.M{"$set": bson.M{"author": newName}}
	_, err := blogRepo.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
