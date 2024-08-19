package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/infrastructure"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository interface {
	Save(blog *Domain.Blog) (*Domain.Blog, error)
	DeleteBlogByID(id string) error
	SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error)
	RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error)
	FindByID(id string) (*Domain.Blog, error)
	IncrementViewCount(blogID string) error
	AddComment(blogID string, comment Domain.Comment) error
	ToggleDislike(blogID, username string) error
	ToggleLike(blogID, username string) error
	FilterBlogs(tags []string, startDate, endDate time.Time, sortBy string) ([]Domain.Blog, error)
}

type blogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(collection *mongo.Collection) *blogRepository {
	return &blogRepository{collection: collection}
}

func (r *blogRepository) Save(blog *Domain.Blog) (*Domain.Blog, error) {
	if blog.Id == "" {
		// Handle the case where the ObjectID is not set
		blog.Id = primitive.NewObjectID().Hex()
	}
	_, err := r.collection.InsertOne(context.TODO(), blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *blogRepository) RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error) {
	var blogs []Domain.Blog
	skip := (page - 1) * pageSize
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: sortBy, Value: -1}}) // Sort by descending order

	cursor, err := r.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var blog Domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, 0, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	// Get the total count of blog posts
	totalPosts, err := r.collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return nil, 0, err
	}

	return blogs, totalPosts, nil
}

func (r *blogRepository) DeleteBlogByID(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (r *blogRepository) SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error) {
	filter := bson.M{}

	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"} // Case-insensitive search
	}
	if author != "" {
		filter["author"] = author
	}
	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	cur, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var blogs []Domain.Blog
	for cur.Next(context.TODO()) {
		var blog Domain.Blog
		if err := cur.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (r *blogRepository) FindByID(id string) (*Domain.Blog, error) {
	var blog Domain.Blog
	filter := bson.M{"id": id}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if no document is found
		}
		return nil, err
	}
	return &blog, nil
}

func (r *blogRepository) IncrementViewCount(blogID string) error {
	filter := bson.M{"id": blogID}
	update := bson.M{"$inc": bson.M{"view_count": 1}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *blogRepository) AddComment(blogID string, comment Domain.Comment) error {
	filter := bson.M{"id": blogID}
	update := bson.M{"$push": bson.M{"comments": comment}}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *blogRepository) ToggleLike(blogID, username string) error {
	blog, err := r.FindByID(blogID)
	if err != nil {
		return err
	}

	var update bson.M
	if infrastructure.Contains(blog.Likes, username) {

		// Remove the like
		update = bson.M{"$pull": bson.M{"likes": username}}
	} else {
		if infrastructure.Contains(blog.Dislikes, username) {
			update = bson.M{"$pull": bson.M{"dislikes": username}}
		}
		// Add the like
		update = bson.M{"$push": bson.M{"likes": username}}
		// Remove dislike if exists
		update["$pull"] = bson.M{"dislikes": username}
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"id": blogID}, update)
	return err
}

func (r *blogRepository) ToggleDislike(blogID, username string) error {
	blog, err := r.FindByID(blogID)
	if err != nil {
		return err
	}

	var update bson.M
	if infrastructure.Contains(blog.Dislikes, username) {
		// Remove the dislike
		update = bson.M{"$pull": bson.M{"dislikes": username}}
	} else {
		if infrastructure.Contains(blog.Likes, username) {
			update = bson.M{"$pull": bson.M{"likes": username}}
		}
		// Add the dislike
		update = bson.M{"$push": bson.M{"dislikes": username}}
		// Remove like if exists
		update["$pull"] = bson.M{"likes": username}
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"id": blogID}, update)
	return err
}

func (r *userRepository) ExpireToken(token string) error {
	// Define the filter to find the token
	filter := bson.M{"access_token": token}

	// Define the update to set the ExpiresAt field to the current time
	update := bson.M{
		"$set": bson.M{
			"expires_at": time.Now().Unix(), // Assuming ExpiresAt is stored as a Unix timestamp
		},
	}

	// Perform the update operation
	_, err := r.tokenCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *blogRepository) FilterBlogs(tags []string, startDate, endDate time.Time, sortBy string) ([]Domain.Blog, error) {
	filter := bson.M{}

	// Filter by tags
	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	// Filter by date range
	if !startDate.IsZero() && !endDate.IsZero() {
		filter["created_at"] = bson.M{
			"$gte": startDate,
			"$lte": endDate,
		}
	} else if !startDate.IsZero() {
		filter["created_at"] = bson.M{"$gte": startDate}
	} else if !endDate.IsZero() {
		filter["created_at"] = bson.M{"$lte": endDate}
	}

	// Create a find options struct to handle sorting
	findOptions := options.Find()

	// Sort by popularity (e.g., view count)
	switch sortBy {
	case "popularity":
		findOptions.SetSort(bson.D{{Key: "view_count", Value: -1}})
	case "latest":
		findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})
	case "oldest":
		findOptions.SetSort(bson.D{{Key: "created_at", Value: 1}})
	default:
		findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})
	}

	cursor, err := r.collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var blogs []Domain.Blog
	for cursor.Next(context.TODO()) {
		var blog Domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}
