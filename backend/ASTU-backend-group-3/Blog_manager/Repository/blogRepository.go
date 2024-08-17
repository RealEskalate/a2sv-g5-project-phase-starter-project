package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository interface {
	Save(blog *Domain.Blog) (*Domain.Blog, error)
	RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error)
}

type blogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(collection *mongo.Collection) *blogRepository {
	return &blogRepository{collection: collection}
}

func (r *blogRepository) Save(blog *Domain.Blog) (*Domain.Blog, error) {
	blog.Id = primitive.NewObjectID().Hex()
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
