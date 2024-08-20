package repository

import (
	// "aastu-backend-group-3/domain"
    "fmt"
    "log"
	"context"
	"errors"
	"group3-blogApi/domain"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoBlogRepository implements the BlogRepository interface using MongoDB
type MongoBlogRepository struct {
    collection *mongo.Collection
}

// NewBlogRepositoryImpl creates a new instance of MongoBlogRepository
func NewBlogRepositoryImpl(coll *mongo.Collection) domain.BlogRepository {
   return &MongoBlogRepository{collection:coll}
}

// CreateBlog creates a new blog post
func (r *MongoBlogRepository) CreateBlog(ctx context.Context, blog domain.Blog) (string, error) {
    blog.CreatedAt = time.Now()
    blog.UpdatedAt = time.Now()

    result, err := r.collection.InsertOne(ctx, blog)
    if err != nil {
        return "", err
    }

    // Convert the inserted ID to a string
    id := result.InsertedID.(primitive.ObjectID).Hex()

    return id, nil
}

// GetBlogByID retrieves a blog post by its ID
func (r *MongoBlogRepository) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
    
    objectID,err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err // Return an error if the ID is not a valid ObjectID
    }
    filter := bson.M{"_id": objectID}
    var blog domain.Blog
    
    er := r.collection.FindOne(ctx, filter).Decode(&blog)
    if er != nil {
        if er == mongo.ErrNoDocuments {
            return nil, nil // domain.Blog not found
        }
        return nil, er
    }

    return &blog, nil
}

// UpdateBlog updates an existing blog post
func (r *MongoBlogRepository) UpdateBlog(ctx context.Context, blog domain.Blog) error {
    filter := bson.M{"id": blog.ID}
    update := bson.M{
        "$set": bson.M{
            "title":         blog.Title,
            "content":       blog.Content,
            "tags":          blog.Tags,
            "updated_at":    time.Now(),
            "likes_count":   blog.LikesCount,
            "dislikes_count": blog.DislikesCount,
            "view_count":    blog.ViewCount,
            "comments_count": blog.CommentsCount,
        },
    }

    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return errors.New("blog not found")
    }

    return nil
}

// DeleteBlog deletes a blog post by its ID
func (r *MongoBlogRepository) DeleteBlog(ctx context.Context, id string) error {
    filter := bson.M{"id": id}
    result, err := r.collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return errors.New("blog not found")
    }

    return nil
}

// GetBlogs retrieves a list of blogs with pagination and sorting
func (r *MongoBlogRepository) GetBlogs(ctx context.Context, offset int64, limit int64, sortBy string) ([]domain.Blog, error) {
    var blogs []domain.Blog

    // Define sorting options
    sortOptions := bson.D{{Key: sortBy, Value: -1}}

    // Attempt to find blogs with the provided parameters
    cursor, err := r.collection.Find(ctx, bson.M{}, &options.FindOptions{
        Skip:  &offset,
        Limit: &limit,
        Sort:  sortOptions,
    })
    if err != nil {
        // Log the error with context and return it
        log.Printf("Error finding blogs: %v. Parameters - SortBy: %s, Offset: %d, Limit: %d", err, sortBy, offset, limit)
        return nil, fmt.Errorf("failed to find blogs. Please check the parameters and try again")
    }
    defer cursor.Close(ctx)

    // Attempt to decode the cursor results into the blogs slice
    if err = cursor.All(ctx, &blogs); err != nil {
        // Log the error with context and return it
        log.Printf("Error decoding blogs from cursor: %v", err)
        return nil, fmt.Errorf("failed to process blogs data. Please try again later")
    }

    return blogs, nil
}

// SearchBlogs searches for blogs based on a query and additional filters
func (r *MongoBlogRepository) SearchBlogs(ctx context.Context, query string, filters map[string]interface{}) ([]domain.Blog, error) {
    var blogs []domain.Blog
    filter := bson.M{"$text": bson.M{"$search": query}}

    if len(filters) > 0 {
        for key, value := range filters {
            filter[key] = value
        }
    }

    cursor, err := r.collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &blogs); err != nil {
        return nil, err
    }

    return blogs, nil
}

// FilterBlogs filters blogs by tags, date, or popularity
func (r *MongoBlogRepository) FilterBlogs(ctx context.Context, filters map[string]interface{}, sortBy string) ([]domain.Blog, error) {
    var blogs []domain.Blog
    filter := bson.M{}

    if len(filters) > 0 {
        for key, value := range filters {
            filter[key] = value
        }
    }

    sortOptions := bson.D{{Key: sortBy, Value: -1}}
    cursor, err := r.collection.Find(ctx, filter, &options.FindOptions{
        Sort: sortOptions,
    })
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &blogs); err != nil {
        return nil, err
    }

    return blogs, nil
}

// TrackPopularity updates the popularity metrics for a blog post
func (r *MongoBlogRepository) TrackPopularity(ctx context.Context, blogID string, action string) error {
    filter := bson.M{"id": blogID}
    var update bson.M

    switch action {
    case "view":
        update = bson.M{"$inc": bson.M{"view_count": 1}}
    case "like":
        update = bson.M{"$inc": bson.M{"likes_count": 1}}
    case "dislike":
        update = bson.M{"$inc": bson.M{"dislikes_count": 1}}
    case "comment":
        update = bson.M{"$inc": bson.M{"comments_count": 1}}
    default:
        return errors.New("invalid action")
    }

    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return errors.New("blog not found")
    }

    return nil
}
