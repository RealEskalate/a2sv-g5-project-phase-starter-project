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
	DeleteBlogByID(id string) error
	SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error)
	RetrieveBlogs(page, pageSize int, sortBy string) ([]Domain.Blog, int64, error)
	FindByID(id string) (*Domain.Blog, error)
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
