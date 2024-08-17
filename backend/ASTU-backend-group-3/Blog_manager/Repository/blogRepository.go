package Repository

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogRepository interface {
	Save(blog *Domain.Blog) (*Domain.Blog, error)
	DeleteBlogByID(id string) error
	SearchBlogs(title string, author string, tags []string) ([]*Domain.Blog, error)
}

type MongoBlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(collection *mongo.Collection) *MongoBlogRepository {
	return &MongoBlogRepository{collection: collection}
}

func (r *MongoBlogRepository) Save(blog *Domain.Blog) (*Domain.Blog, error) {
	blog.Id = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.TODO(), blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *MongoBlogRepository) DeleteBlogByID(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (r *MongoBlogRepository) SearchBlogs(title string, author string, tags []string) ([]Domain.Blog, error) {
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
