package repository

import (
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog/database"
	"errors"

	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	database   database.Database
	collection string
}

// NewBlogRepository returns a new instance of blogRepository implementing the domain.BlogRepository interface.
func NewBlogRepository(db database.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

// CreateBlog inserts a new blog into the MongoDB collection.
func (r *blogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(ctx, blog)

	return err
}

// GetBlogByID fetches a blog by its ID from the MongoDB collection.
func (r *blogRepository) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {
	var blog domain.Blog
	collection := r.database.Collection(r.collection)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

// GetAllBlogs fetches all blogs with pagination and sorting.
func (r *blogRepository) GetAllBlogs(ctx context.Context, page, limit int, sortBy string) ([]*domain.Blog, error) {
	var blogs []*domain.Blog

	// Validate pagination inputs
	if page < 1 || limit < 1 {
		return nil, fmt.Errorf("invalid pagination parameters: page and limit must be greater than 0")
	}

	// Determine the sort field based on the input
	sortField := sortBy
	if sortBy == "likes" {
		sortField = "likes" // Ensure this matches the exact field name in your MongoDB collection
	}

	// Calculate the number of documents to skip based on the page number and limit
	skip := (page - 1) * limit

	// Configure find options with sorting, skipping, and limiting
	findOptions := options.Find().
		SetSort(bson.D{{Key: sortField, Value: -1}}). // Sort in descending order (most likes first)
		SetSkip(int64(skip)).
		SetLimit(int64(limit))

	// Access the collection
	collection := r.database.Collection(r.collection)

	// Execute the query to retrieve the blogs
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch blogs: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into a Blog object
	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, fmt.Errorf("failed to decode blog: %w", err)
		}
		blogs = append(blogs, &blog)
	}

	// Check for any errors encountered during the iteration
	if cursor == nil {
		return nil, fmt.Errorf("cursor iteration error: %w", err)
	}

	return blogs, nil
}

// UpdateBlog updates a blog in the MongoDB collection.
func (r *blogRepository) UpdateBlog(ctx context.Context, blog *domain.Blog) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": blog.ID}, bson.M{"$set": blog})
	return err
}

// DeleteBlog deletes a blog by its ID from the MongoDB collection.
func (r *blogRepository) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Repositories/blog_repository.go
// SearchBlogs searches for blogs based on query and filters.
func (r *blogRepository) SearchBlogs(ctx context.Context, title string, author string) (*[]domain.Blog, error) {
	filter := bson.M{}

	// Add search filters based on the provided title and author
	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"}
	}
	if author != "" {
		filter["author"] = bson.M{"$regex": author, "$options": "i"}
	}

	var blogs []domain.Blog
	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the cursor results into the blogs slice
	if err = cursor.All(context.Background(), &blogs); err != nil {
		return nil, err
	}

	return &blogs, nil

}

func (r *blogRepository) FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*domain.Blog, error) {
	var blogs []*domain.Blog

	// Define a filter without conditions initially
	filter := bson.M{}

	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	// Parse the startDate and endDate
	if startDate != "" {
		startTime, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			return nil, fmt.Errorf("invalid startDate format")
		}
		filter["created_at"] = bson.M{"$gte": startTime}
	}

	if endDate != "" {
		endTime, err := time.Parse(time.RFC3339, endDate)
		if err != nil {
			return nil, fmt.Errorf("invalid endDate format")
		}
		if existingFilter, ok := filter["created_at"].(bson.M); ok {
			existingFilter["$lte"] = endTime
		} else {
			filter["created_at"] = bson.M{"$lte": endTime}
		}
	}

	sortOptions := bson.D{}

	if popularity != "" {
		switch popularity {
		case "most_viewed":
			sortOptions = bson.D{{Key: "views", Value: -1}}
		case "most_liked":
			sortOptions = bson.D{{Key: "likes", Value: -1}}
		case "most_commented":
			sortOptions = bson.D{{Key: "comments", Value: -1}}
		case "most_disliked":
			sortOptions = bson.D{{Key: "dislikes", Value: -1}}
		case "most_popular":
			sortOptions = bson.D{{Key: "popularity", Value: -1}}
		}

	}

	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(sortOptions))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (r *blogRepository) AddComment(ctx context.Context, id primitive.ObjectID, comment *domain.Comment) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"comments": comment}}
	_, err := r.database.Collection(r.collection).UpdateOne(ctx, filter, update)
	return err
}

func (r *blogRepository) HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID string) (bool, error) {
	filter := bson.M{"_id": id, "dislikes": userID}
	count, err := r.database.Collection(r.collection).CountDocuments(ctx, filter)
	return count > 0, err
}

func (r *blogRepository) IncrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{metric: 1}}
	_, err := r.database.Collection(r.collection).UpdateOne(ctx, filter, update)
	return err
}

func (r *blogRepository) DecrementPopularity(ctx context.Context, postID primitive.ObjectID, metric string) error {
	collection := r.database.Collection(r.collection)

	// Find the current value of the metric
	var result bson.M
	err := collection.FindOne(ctx, bson.M{"_id": postID}).Decode(&result)
	if err != nil {
		return err
	}

	// Handle the value according to its type
	currentValue, ok := result[metric].(int32) // Attempt to assert as int32
	if !ok {
		// If not int32, check if it's an int
		if intValue, ok := result[metric].(int); ok {
			currentValue = int32(intValue)
		} else {
			return errors.New("unsupported type for popularity metric")
		}
	}

	// Prevent the value from decreasing below 0
	if currentValue > 0 {
		newValue := currentValue - 1

		// Update the metric in the database
		_, err = collection.UpdateOne(ctx, bson.M{"_id": postID}, bson.M{"$set": bson.M{metric: newValue}})
		if err != nil {
			return err
		}
	}

	return nil
}
